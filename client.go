package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hashicorp/go-hclog"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier/pkg/plugins"
)

//go:generate abigen --sol ./example/broker.sol --pkg main --out broker.go
type Client struct {
	abi       abi.ABI
	config    *Config
	ctx       context.Context
	cancel    context.CancelFunc
	ethClient *ethclient.Client
	session   *BrokerSession
	eventC    chan *pb.IBTP
}

var (
	_      plugins.Client = (*Client)(nil)
	logger                = hclog.New(&hclog.LoggerOptions{
		Name:   "client",
		Output: os.Stderr,
		Level:  hclog.Trace,
	})
	EtherType = "ethereum"
)

const InvokeInterchain = "invokeInterchain"

func (c *Client) GetUpdateMeta() chan *pb.UpdateMeta {
	panic("implement me")
}

func (c *Client) Initialize(configPath string, extra []byte) error {
	cfg, err := UnmarshalConfig(configPath)
	if err != nil {
		return fmt.Errorf("unmarshal config for plugin :%w", err)
	}

	logger.Info("Basic appchain info",
		"broker address", cfg.Ether.ContractAddress,
		"ethereum node ip", cfg.Ether.Addr)

	etherCli, err := ethclient.Dial(cfg.Ether.Addr)
	if err != nil {
		return fmt.Errorf("dial ethereum node: %w", err)
	}

	keyPath := filepath.Join(configPath, cfg.Ether.KeyPath)
	keyByte, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}

	psdPath := filepath.Join(configPath, cfg.Ether.Password)
	password, err := ioutil.ReadFile(psdPath)
	if err != nil {
		return err
	}

	unlockedKey, err := keystore.DecryptKey(keyByte, strings.TrimSpace(string(password)))
	if err != nil {
		return err
	}

	chainID, err := etherCli.ChainID(context.TODO())
	if err != nil {
		return fmt.Errorf("cannot get ethereum chain ID: %sv", err)
	}

	// deploy a contract first
	auth, err := bind.NewKeyedTransactorWithChainID(unlockedKey.PrivateKey, chainID)
	if err != nil {
		return err
	}
	if auth.Context == nil {
		auth.Context = context.TODO()
	}
	auth.Value = nil
	broker, err := NewBroker(common.HexToAddress(cfg.Ether.ContractAddress), etherCli)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Broker contract: %w", err)
	}
	session := &BrokerSession{
		Contract: broker,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}

	ab, err := abi.JSON(bytes.NewReader([]byte(BrokerABI)))
	if err != nil {
		return fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	c.config = cfg
	c.eventC = make(chan *pb.IBTP, 1024)
	c.ethClient = etherCli
	c.session = session
	c.abi = ab
	c.ctx, c.cancel = context.WithCancel(context.Background())

	return nil
}

func (c *Client) Start() error {
	return c.StartConsumer()
}

func (c *Client) Stop() error {
	c.cancel()
	return nil
}

func (c *Client) GetIBTPCh() chan *pb.IBTP {
	return c.eventC
}

func (c *Client) Name() string {
	return c.config.Ether.Name
}

func (c *Client) Type() string {
	return EtherType
}

// SubmitIBTP submit interchain ibtp. It will unwrap the ibtp and execute
// the function inside the ibtp. If any execution results returned, pass
// them to other modules.
func (c *Client) SubmitIBTP(from string, index uint64, serviceID string, ibtpType pb.IBTP_Type, content *pb.Content, proof *pb.BxhProof, isEncrypted bool) (*pb.SubmitIBTPResponse, error) {
	ret := &pb.SubmitIBTPResponse{Status: true}
	receipt, err := c.invokeInterchain(from, index, serviceID, uint64(ibtpType), content.Func, content.Args, uint64(proof.TxStatus), proof.MultiSign, isEncrypted)
	if err != nil {
		ret.Status = false
		ret.Message = err.Error()
		return ret, nil
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		ret.Status = false
		ret.Message = fmt.Sprintf("SubmitIBTP tx execution failed")
		return ret, nil
	}

	return ret, nil
}

func (c *Client) SubmitReceipt(to string, index uint64, serviceID string, ibtpType pb.IBTP_Type, result *pb.Result, proof *pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	ret := &pb.SubmitIBTPResponse{Status: true}
	receipt, err := c.invokeReceipt(serviceID, to, index, uint64(ibtpType), result.Data, uint64(proof.TxStatus), proof.MultiSign)
	if err != nil {
		ret.Status = false
		ret.Message = err.Error()
		return ret, nil
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		ret.Status = false
		ret.Message = fmt.Sprintf("SubmitReceipt tx execution failed")
	}

	return ret, nil
}

func (c *Client) invokeInterchain(srcFullID string, index uint64, destAddr string, reqType uint64, callFunc string, args [][]byte, txStatus uint64, multiSign [][]byte, encrypt bool) (*types.Receipt, error) {
	var tx *types.Transaction
	var txErr error
	if err := retry.Retry(func(attempt uint) error {
		tx, txErr = c.session.InvokeInterchain(srcFullID, common.HexToAddress(destAddr), index, reqType, callFunc, args, txStatus, multiSign, encrypt)
		if txErr != nil {
			logger.Warn("Call InvokeInterchain failed",
				"srcFullID", srcFullID,
				"destAddr", destAddr,
				"index", fmt.Sprintf("%d", index),
				"reqType", strconv.Itoa(int(reqType)),
				"callFunc", callFunc,
				"args", string(bytes.Join(args, []byte(","))),
				"txStatus", strconv.Itoa(int(txStatus)),
				"multiSign size", strconv.Itoa(len(multiSign)),
				"encrypt", strconv.FormatBool(encrypt),
				"error", txErr.Error(),
			)

			for i, arg := range args {
				logger.Warn("args", strconv.Itoa(i), hexutil.Encode(arg))
			}

			for i, sign := range multiSign {
				logger.Warn("multiSign", strconv.Itoa(i), hexutil.Encode(sign))
			}

			if strings.Contains(txErr.Error(), "execution reverted") {
				return nil
			}
		}

		return txErr
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't invoke contract", "error", err)
	}

	if txErr != nil {
		return nil, txErr
	}

	return c.waitForConfirmed(tx.Hash()), nil
}

func (c *Client) invokeReceipt(srcAddr string, dstFullID string, index uint64, reqType uint64, result [][]byte, txStatus uint64, multiSign [][]byte) (*types.Receipt, error) {
	var tx *types.Transaction
	var txErr error
	if err := retry.Retry(func(attempt uint) error {
		tx, txErr = c.session.InvokeReceipt(common.HexToAddress(srcAddr), dstFullID, index, reqType, result, txStatus, multiSign)
		if txErr != nil {
			logger.Warn("Call InvokeReceipt failed",
				"srcAddr", srcAddr,
				"dstFullID", dstFullID,
				"index", fmt.Sprintf("%d", index),
				"reqType", strconv.Itoa(int(reqType)),
				"result", string(bytes.Join(result, []byte(","))),
				"txStatus", strconv.Itoa(int(txStatus)),
				"multiSign size", strconv.Itoa(len(multiSign)),
				"error", txErr.Error(),
			)

			for i, arg := range result {
				logger.Warn("result", strconv.Itoa(i), hexutil.Encode(arg))
			}

			for i, sign := range multiSign {
				logger.Warn("multiSign", strconv.Itoa(i), hexutil.Encode(sign))
			}

			if strings.Contains(txErr.Error(), "execution reverted") {
				return nil
			}
		}

		return txErr
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't invoke contract", "error", err)
	}

	if txErr != nil {
		return nil, txErr
	}

	return c.waitForConfirmed(tx.Hash()), nil
}

// GetOutMessage gets crosschain tx by `to` address and index
func (c *Client) GetOutMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	srcService, dstService, err := pb.ParseServicePair(servicePair)
	if err != nil {
		return nil, err
	}

	ev := &BrokerThrowInterchainEvent{
		Index:     idx,
		DstFullID: dstService,
		SrcFullID: srcService,
	}

	return c.Convert2IBTP(ev, int64(c.config.Ether.TimeoutHeight))
}

// GetInMessage gets the execution results from contract by from-index key
func (c *Client) GetReceiptMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	var (
		data    [][]byte
		status  bool
		encrypt bool
	)

	if err := retry.Retry(func(attempt uint) error {
		var err error
		data, status, encrypt, err = c.session.GetReceiptMessage(servicePair, idx)
		if err != nil {
			logger.Error("get receipt message", "servicePair", servicePair, "err", err.Error())
		}
		return err
	}); err != nil {
		logger.Error("retry error in GetInMessage", "err", err.Error())
		return nil, err
	}

	srcServiceID, dstServiceID, err := pb.ParseServicePair(servicePair)
	if err != nil {
		return nil, err
	}

	return generateReceipt(srcServiceID, dstServiceID, idx, data, status, encrypt)
}

// GetInMeta queries contract about how many interchain txs have been
// executed on this appchain for different source chains.
func (c *Client) GetInMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetInnerMeta)
}

// GetOutMeta queries contract about how many interchain txs have been
// sent out on this appchain to different destination chains.
func (c *Client) GetOutMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetOuterMeta)
}

// GetCallbackMeta queries contract about how many callback functions have been
// executed on this appchain from different destination chains.
func (c *Client) GetCallbackMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetCallbackMeta)
}

func (c *Client) getMeta(getMetaFunc func() ([]string, []uint64, error)) (map[string]uint64, error) {
	var (
		appchainIDs []string
		indices     []uint64
		err         error
	)
	meta := make(map[string]uint64, 0)

	appchainIDs, indices, err = getMetaFunc()
	if err != nil {
		return nil, err
	}

	for i, did := range appchainIDs {
		meta[did] = indices[i]
	}

	return meta, nil
}

func (c *Client) getBestBlock() uint64 {
	var blockNum uint64

	if err := retry.Retry(func(attempt uint) error {
		var err error
		blockNum, err = c.ethClient.BlockNumber(c.ctx)
		if err != nil {
			logger.Error("retry failed in getting best block", "err", err.Error())
		}
		return err
	}, strategy.Wait(time.Second*10)); err != nil {
		logger.Error("retry failed in get best block", "err", err.Error())
		panic(err)
	}

	return blockNum
}

func (c *Client) waitForConfirmed(hash common.Hash) *types.Receipt {
	var (
		receipt *types.Receipt
		err     error
	)

	start := c.getBestBlock()

	for c.getBestBlock()-start < c.config.Ether.MinConfirm {
		time.Sleep(time.Second * 5)
	}
	if err := retry.Retry(func(attempt uint) error {
		receipt, err = c.ethClient.TransactionReceipt(c.ctx, hash)
		if err != nil {
			return err
		}

		return nil
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't get receipt for tx", hash.Hex(), "error", err)
	}

	return receipt
}

func (c *Client) GetDstRollbackMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetDstRollbackMeta)
}

func (c *Client) GetChainID() (string, string, error) {
	return c.session.GetChainID()
}

func (c *Client) GetServices() ([]string, error) {
	return c.session.GetLocalServiceList()
}

func (c *Client) GetAppchainInfo(chainID string) (string, []byte, string, error) {
	broker, trustRoot, ruleAddr, err := c.session.GetAppchainInfo(chainID)
	if err != nil {
		return "", nil, "", err
	}

	return broker, trustRoot, ruleAddr.String(), nil
}
