package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
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
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier-client-ethereum/solidity"
	"github.com/meshplus/pier/pkg/plugins"
)

//go:generate abigen --sol ./example/broker.sol --pkg main --out broker.go
type Client struct {
	abi        abi.ABI
	config     *Config
	ctx        context.Context
	ethClient  *ethclient.Client
	broker     *Broker
	session    *BrokerSession
	conn       *rpc.Client
	eventC     chan *pb.IBTP
	appchainID string // the method of connected appchain like did:bitxhub:appchain:.
	bizABI     map[string]*abi.ABI
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

func (c *Client) Initialize(configPath string, appchainID string, extra []byte) error {
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

	conn, err := rpc.Dial(cfg.Ether.Addr)
	if err != nil {
		return fmt.Errorf("rpc dial: %s", err.Error())
	}

	bizAbi := make(map[string]*abi.ABI)
	for _, service := range cfg.Services {
		logger.Info("ContractABI", "addr", service.ID, "name", service.Name)
		content, err := ioutil.ReadFile(filepath.Join(configPath, service.Abi))
		if err != nil {
			return fmt.Errorf("read abi %s for addr %s: %w", service.Name, service.ID, err)
		}
		abi, err := abi.JSON(bytes.NewReader(content))
		if err != nil {
			return fmt.Errorf("unmarshal abi %s for addr %s: %s", service.Name, service.ID, err.Error())
		}
		bizAbi[service.ID] = &abi
	}

	c.config = cfg
	c.eventC = make(chan *pb.IBTP, 1024)
	c.ethClient = etherCli
	c.broker = broker
	c.session = session
	c.abi = ab
	c.conn = conn
	c.appchainID = appchainID
	c.ctx = context.Background()
	c.bizABI = bizAbi

	return nil
}

func (c *Client) Start() error {
	return c.StartConsumer()
}

func (c *Client) Stop() error {
	return nil
}

func (c *Client) Name() string {
	return c.config.Ether.Name
}

func (c *Client) Type() string {
	return EtherType
}

func (c *Client) GetIBTP() chan *pb.IBTP {
	return c.eventC
}

// SubmitIBTP submit interchain ibtp. It will unwrap the ibtp and execute
// the function inside the ibtp. If any execution results returned, pass
// them to other modules.
func (c *Client) SubmitIBTP(ibtp *pb.IBTP) (*pb.SubmitIBTPResponse, error) {
	ret := &pb.SubmitIBTPResponse{Status: true}
	pd := &pb.Payload{}
	if err := pd.Unmarshal(ibtp.Payload); err != nil {
		return nil, fmt.Errorf("ibtp payload unmarshal: %w", err)
	}

	content := &pb.Content{}
	if err := content.Unmarshal(pd.Content); err != nil {
		return ret, fmt.Errorf("ibtp content unmarshal: %w", err)
	}

	if ibtp.Category() == pb.IBTP_UNKNOWN {
		return nil, fmt.Errorf("invalid ibtp category")
	}

	var (
		bizData           []byte
		err               error
		serviceID         string
		srcChainServiceID string
		ok                bool
		bAbi              *abi.ABI
	)

	if ibtp.Category() == pb.IBTP_REQUEST {
		srcChainServiceID = ibtp.From
		_, _, serviceID, err = parseChainServiceID(ibtp.To)
	} else {
		srcChainServiceID = ibtp.To
		_, _, serviceID, err = parseChainServiceID(ibtp.From)
	}

	var result [][]byte
	logger.Info("submit ibtp", "from", srcChainServiceID, "to", serviceID, "index", ibtp.Index, "func", content.Func)
	for i, arg := range content.Args {
		logger.Info("arg", strconv.Itoa(i), string(arg))
	}

	if ibtp.Category() == pb.IBTP_RESPONSE && content.Func == "" || ibtp.Type == pb.IBTP_ROLLBACK {
		logger.Info("InvokeIndexUpdate", "ibtp", ibtp.ID())
		success, err := c.InvokeIndexUpdate(srcChainServiceID, ibtp.Index, serviceID, uint64(ibtp.Category()))
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("update index for ibtp %s failed", ibtp.ID())
		}
		ret.Status = true

		if ibtp.Type == pb.IBTP_ROLLBACK {
			ret.Result, err = c.generateCallback(ibtp, nil, ret.Status)
			if err != nil {
				return nil, err
			}
		}
		return ret, nil
	}

	if err != nil {
		ret.Message = fmt.Sprintf("pack for ibtp %s func %s and args: %s", ibtp.ID(), content.Func, err)
	} else {
		bAbi, ok = c.bizABI[serviceID]
		if !ok {
			ret.Message = fmt.Sprintf("no abi for contract %s", serviceID)
		} else {
			bizData, err = c.packFuncArgs(content.Func, content.Args, bAbi)
			if err != nil {
				ret.Message = fmt.Sprintf("pack for ibtp %s func %s and args: %s", ibtp.ID(), content.Func, err)
			}
		}
	}
	if !ok || err != nil {
		ret.Status = false
		success, err := c.InvokeIndexUpdateWithError(srcChainServiceID, ibtp.Index, serviceID, uint64(ibtp.Category()), ret.Message)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("update index for ibtp %s failed", ibtp.ID())
		}
	} else {
		receipt, err := c.InvokeInterchain(srcChainServiceID, ibtp.Index, serviceID, uint64(ibtp.Category()), bizData)
		if err != nil {
			return nil, fmt.Errorf("invoke interchain for ibtp %s to call %s: %w", ibtp.ID(), content.Func, err)
		}

		for i, log := range receipt.Logs {
			logger.Info("log", "index", strconv.Itoa(i), "data", hexutil.Encode(log.Data))
		}

		if receipt.Status == types.ReceiptStatusSuccessful {
			if len(receipt.Logs) == 0 {
				return nil, fmt.Errorf("no log found for ibtp %s", ibtp.ID())
			}
			ret.Status, result, err = solidity.Unpack(*bAbi, content.Func, receipt.Logs[len(receipt.Logs)-1].Data)
			if err != nil {
				return nil, fmt.Errorf("unpack for ibtp %s: %w", ibtp.ID(), err)
			}
		} else {
			ret.Status = false
			ret.Message = fmt.Sprintf("InvokeInterchain tx execution failed")
			success, err := c.InvokeIndexUpdateWithError(srcChainServiceID, ibtp.Index, serviceID, uint64(ibtp.Category()), ret.Message)
			if err != nil {
				return nil, err
			}
			if !success {
				return ret, fmt.Errorf("invalid index for ibtp %s", ibtp.ID())
			}
		}
	}

	// If is response IBTP, then simply return
	if ibtp.Category() == pb.IBTP_RESPONSE {
		return ret, nil
	}

	ret.Result, err = c.generateCallback(ibtp, result, ret.Status)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) InvokeInterchain(srcChainServiceID string, index uint64, destAddr string, reqType uint64, bizCallData []byte) (*types.Receipt, error) {
	var tx *types.Transaction
	var txErr error
	if err := retry.Retry(func(attempt uint) error {
		tx, txErr = c.session.InvokeInterchain(srcChainServiceID, index, common.HexToAddress(destAddr), reqType, bizCallData)
		if txErr != nil {
			logger.Info("Call InvokeInterchain failed",
				"srcChainServiceID", srcChainServiceID,
				"index", fmt.Sprintf("%d", index),
				"destAddr", destAddr,
				"error", txErr.Error(),
			)
		}

		return txErr
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't invoke contract", "error", err)
	}

	return c.waitForConfirmed(tx.Hash()), nil
}

func (c *Client) InvokeIndexUpdateWithError(srcChainServiceID string, index uint64, dstServiceID string, reqType uint64, errMsg string) (bool, error) {
	var tx *types.Transaction
	var txErr error
	if err := retry.Retry(func(attempt uint) error {
		tx, txErr = c.session.InvokeIndexUpdateWithError(srcChainServiceID, index, common.HexToAddress(dstServiceID), reqType, errMsg)
		if txErr != nil {
			logger.Info("Call InvokeIndexUpdateWithError failed",
				"srcChainServiceID", srcChainServiceID,
				"index", fmt.Sprintf("%d", index),
				"error", txErr.Error(),
			)
		}

		return txErr
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't invoke contract", "error", err)
	}

	receipt := c.waitForConfirmed(tx.Hash())

	return receipt.Status == types.ReceiptStatusSuccessful, nil
}

func (c *Client) InvokeIndexUpdate(srcChainID string, index uint64, dstServiceID string, reqType uint64) (bool, error) {
	return c.InvokeIndexUpdateWithError(srcChainID, index, dstServiceID, reqType, "")
}

// GetOutMessage gets crosschain tx by `to` address and index
func (c *Client) GetOutMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	var blockNum *big.Int
	if err := retry.Retry(func(attempt uint) error {
		var err error
		blockNum, err = c.session.GetOutMessage(servicePair, idx)
		if err != nil {
			logger.Error("get out message", "err", err.Error())
		}
		return err
	}, strategy.Wait(time.Second*10)); err != nil {
		logger.Info("retry error in get out message", "err", err.Error())
	}

	height := blockNum.Uint64()
	var throwEvents *BrokerThrowEventIterator
	if err := retry.Retry(func(attempt uint) error {
		var err error
		throwEvents, err = c.broker.FilterThrowEvent(&bind.FilterOpts{
			Start:   height,
			End:     &height,
			Context: c.ctx,
		})
		if err != nil {
			logger.Error("FilterThrowEvent", "err", err.Error())
		}
		return err
	}, strategy.Wait(time.Second*3)); err != nil {
		logger.Error("retry failed", "err", err.Error())
	}

	srcService, dstService, err := parseServicePair(servicePair)
	if err != nil {
		return nil, err
	}

	for throwEvents.Next() {
		ev := throwEvents.Event
		logger.Info("throw event", "SrcFullID", ev.SrcFullID, "DstFullID", ev.DstFullID, "index", ev.Index)
		logger.Info("exp event", "srcService", srcService, "dstService", dstService, "idx", idx)

		if ev.SrcFullID == srcService && ev.DstFullID == dstService && ev.Index == idx {
			return Convert2IBTP(ev, int64(c.config.Ether.TimeoutHeight), pb.IBTP_INTERCHAIN), nil
		}
	}

	return nil, fmt.Errorf("cannot find out ibtp for service pair %s and index %d", servicePair, idx)
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

// GetInMessage gets the execution results from contract by from-index key
func (c *Client) GetInMessage(from string, idx uint64) ([][]byte, error) {
	var blockNum *big.Int
	if err := retry.Retry(func(attempt uint) error {
		var err error
		blockNum, err = c.session.GetInMessage(from, idx)
		if err != nil {
			logger.Error("get in message", "err", err.Error())
		}
		return err
	}); err != nil {
		logger.Error("retry error in GetInMessage", "err", err.Error())
	}

	return [][]byte{blockNum.Bytes()}, nil
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

func (c *Client) CommitCallback(ibtp *pb.IBTP) error {
	return nil
}

// @ibtp is the original ibtp merged from this appchain
func (c *Client) RollbackIBTP(ibtp *pb.IBTP, isSrcChain bool) (*pb.RollbackIBTPResponse, error) {
	ret := &pb.RollbackIBTPResponse{Status: true}
	pd := &pb.Payload{}
	if err := pd.Unmarshal(ibtp.Payload); err != nil {
		return nil, fmt.Errorf("ibtp payload unmarshal: %w", err)
	}
	content := &pb.Content{}
	if err := content.Unmarshal(pd.Content); err != nil {
		return ret, fmt.Errorf("ibtp content unmarshal: %w", err)
	}

	if content.Rollback == "" {
		logger.Info("rollback function is empty, ignore it", "func", content.Func, "callback", content.Callback, "rollback", content.Rollback)
		return ret, nil
	}

	var (
		bizData           []byte
		err               error
		serviceID         string
		srcChainServiceID string
		ok                bool
		bAbi              *abi.ABI
		rollbackFunc      string
		rollbackArgs      [][]byte
		reqType           uint64
	)

	if isSrcChain {
		rollbackFunc = content.Rollback
		rollbackArgs = content.ArgsRb
		srcChainServiceID = ibtp.To
		_, _, serviceID, err = parseChainServiceID(ibtp.From)
		reqType = 1
	} else {
		rollbackFunc = content.Func
		rollbackArgs = content.Args
		rollbackArgs[len(rollbackArgs)-1] = []byte("true")
		srcChainServiceID = ibtp.From
		_, _, serviceID, err = parseChainServiceID(ibtp.To)
		reqType = 2
	}

	bAbi, ok = c.bizABI[serviceID]
	if !ok {
		ret.Message = fmt.Sprintf("no abi for contract %s", serviceID)
	} else {
		bizData, err = c.packFuncArgs(rollbackFunc, rollbackArgs, bAbi)
		if err != nil {
			ret.Message = fmt.Sprintf("pack for ibtp %s func %s and args: %s", ibtp.ID(), content.Func, err)
		}
	}

	logger.Info("rollback ibtp", "from", srcChainServiceID, "to", serviceID, "index", ibtp.Index, "func", rollbackFunc)
	for i, arg := range rollbackArgs {
		logger.Info("arg", strconv.Itoa(i), string(arg))
	}

	if !ok || err != nil {
		ret.Status = false
		success, err := c.InvokeIndexUpdateWithError(srcChainServiceID, ibtp.Index, serviceID, reqType, ret.Message)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("update index for ibtp %s failed", ibtp.ID())
		}
	} else {
		// false indicates it is for rollback
		receipt, err := c.InvokeInterchain(srcChainServiceID, ibtp.Index, serviceID, reqType, bizData)
		if err != nil {
			return nil, err
		}
		if len(receipt.Logs) == 0 {
			ret.Status = false
			ret.Message = "wrong contract doesn't emit log event"
		}
	}

	return ret, nil
}

func (c *Client) IncreaseInMeta(original *pb.IBTP) (*pb.IBTP, error) {
	ibtp, err := c.generateCallback(original, nil, false)
	if err != nil {
		return nil, err
	}
	errMsg := "ibtp failed in bitxhub"
	_, _, serviceID, err := parseChainServiceID(ibtp.To)
	if err != nil {
		return nil, err
	}
	success, err := c.InvokeIndexUpdateWithError(original.From, original.Index, serviceID, 0, errMsg)
	if err != nil {
		logger.Error(errMsg, "ibtp_id", ibtp.ID(), "error", err.Error())
		return nil, err
	}
	if !success {
		logger.Error("invalid index of ibtp", "ibtp_id", ibtp.ID())
	}
	return ibtp, nil
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

func (c *Client) getTxReceipt(txHash common.Hash) *types.Receipt {
	var receipt *types.Receipt

	if err := retry.Retry(func(attempt uint) error {
		var err error
		receipt, err = c.ethClient.TransactionReceipt(c.ctx, txHash)
		if err != nil {
			logger.Info("get receipt error", "err", err.Error())
		}
		return err
	}, strategy.Wait(time.Second*10)); err != nil {
		logger.Error("retry failed in get tx receipt", "err", err.Error())
		panic(err)
	}

	return receipt
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

func (c *Client) GetReceipt(ibtp *pb.IBTP) (*pb.IBTP, error) {
	blockBytes, err := c.GetInMessage(ibtp.ServicePair(), ibtp.Index)
	if err != nil {
		return nil, err
	}

	if len(blockBytes) != 1 {
		return nil, fmt.Errorf("GetInMessage can not get block number ")
	}

	blockNum := &big.Int{}
	blockNum.SetBytes(blockBytes[0])

	block, err := c.ethClient.BlockByNumber(c.ctx, blockNum)
	if err != nil {
		return nil, err
	}

	pl := &pb.Payload{}
	if err := pl.Unmarshal(ibtp.Payload); err != nil {
		return nil, err
	}

	ct := &pb.Content{}
	if err := ct.Unmarshal(pl.Content); err != nil {
		return nil, err
	}

	_, _, serviceID, err := parseChainServiceID(ibtp.To)
	if err != nil {
		return nil, err
	}

	bAbi, ok := c.bizABI[serviceID]
	if !ok {
		return nil, fmt.Errorf("can not find abi for contract %s", serviceID)
	}
	bizData, err := c.packFuncArgs(ct.Func, ct.Args, bAbi)
	packData, err := c.abi.Pack(InvokeInterchain, ibtp.From, ibtp.Index, common.HexToAddress(serviceID), uint64(ibtp.Category()), bizData)
	if err != nil {
		return nil, err
	}
	for _, tx := range block.Transactions() {
		if tx.To().String() == c.config.Ether.ContractAddress && bytes.Equal(packData, tx.Data()) {
			receipt := c.getTxReceipt(tx.Hash())
			if len(receipt.Logs) == 0 || receipt.Status == types.ReceiptStatusFailed {
				return c.generateCallback(ibtp, nil, false)
			}

			log := receipt.Logs[len(receipt.Logs)-1]
			status, result, err := solidity.Unpack(*bAbi, ct.Func, log.Data)
			if err != nil {
				return nil, fmt.Errorf("unpack log data %s for func %s", hexutil.Encode(log.Data), ct.Func)
			}
			return c.generateCallback(ibtp, result, status)
		}
	}

	return c.generateCallback(ibtp, nil, false)
}

func (c *Client) GetSrcRollbackMeta() (map[string]uint64, error) {
	panic("implement me")
}

func (c *Client) GetDstRollbackMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetDstRollbackMeta)
}

func (c *Client) GetChainID() (string, string) {
	var (
		bxhID      string
		appchainID string
		err        error
	)
	if err := retry.Retry(func(attempt uint) error {
		bxhID, appchainID, err = c.session.GetChainID()
		if err != nil {
			return err
		}

		return nil
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't get chain ID", "error", err)
	}

	return bxhID, appchainID
}

func (c *Client) GetServices() []string {
	var services []string

	for _, service := range c.config.Services {
		services = append(services, service.ID)
	}

	return services
}

func (c *Client) packFuncArgs(function string, args [][]byte, abi *abi.ABI) ([]byte, error) {
	var argx []interface{}
	var err error

	if len(args) != 0 {
		argx, err = solidity.ABIUnmarshal(*abi, args, function)
		if err != nil {
			return nil, err
		}
	}

	packed, err := abi.Pack(function, argx...)
	if err != nil {
		return nil, err
	}

	return packed, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugins.Handshake,
		Plugins: map[string]plugin.Plugin{
			plugins.PluginName: &plugins.AppchainGRPCPlugin{Impl: &Client{}},
		},
		Logger: logger,
		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})

	logger.Info("Plugin server down")
}

func parseChainServiceID(id string) (string, string, string, error) {
	splits := strings.Split(id, ":")
	if len(splits) != 3 {
		return "", "", "", fmt.Errorf("invalid chain service ID: %s", id)
	}

	return splits[0], splits[1], splits[2], nil
}

func parseServicePair(servicePair string) (string, string, error) {
	splits := strings.Split(servicePair, "-")
	if len(splits) != 2 {
		return "", "", fmt.Errorf("invalid service pair: %s", servicePair)
	}

	return splits[0], splits[1], nil
}
