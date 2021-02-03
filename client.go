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
	abi       abi.ABI
	config    *Config
	ctx       context.Context
	ethClient *ethclient.Client
	broker    *Broker
	session   *BrokerSession
	conn      *rpc.Client
	eventC    chan *pb.IBTP
	pierID    string
}

var (
	eventSigs = map[string]string{
		"0x436160f7c24c5f31561ec9422a629accdbbd4e9e8ce21e86e634f497997769a8": "logInterchainData",
		"0x23de11857b4338b8e6ccaec81162b447b44040ff3cfdd1174d548975eb5c1c3e": "logInterchainStatus",
		"0xad89cfa05a757be8d2179bb6609bf9034971b2427bd49d48e79552d3e8493e99": "interchainEvent",
	}
	_      plugins.Client = (*Client)(nil)
	logger                = hclog.New(&hclog.LoggerOptions{
		Name:   "client",
		Output: os.Stderr,
		Level:  hclog.Trace,
	})
	EtherType = "ethereum"
)

func (c *Client) Initialize(configPath string, pierID string, extra []byte) error {
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

	// deploy a contract first
	auth := bind.NewKeyedTransactor(unlockedKey.PrivateKey)
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

	b, err := ioutil.ReadFile(filepath.Join(configPath, cfg.Ether.AbiPath))
	ab, err := abi.JSON(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	conn, err := rpc.Dial(cfg.Ether.Addr)
	if err != nil {
		return fmt.Errorf("rpc dial: %s", err.Error())
	}

	c.config = cfg
	c.eventC = make(chan *pb.IBTP, 1024)
	c.ethClient = etherCli
	c.broker = broker
	c.session = session
	c.abi = ab
	c.conn = conn
	c.pierID = pierID
	c.ctx = context.Background()
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
	ret := &pb.SubmitIBTPResponse{}
	pd := &pb.Payload{}
	if err := pd.Unmarshal(ibtp.Payload); err != nil {
		return nil, fmt.Errorf("ibtp payload unmarshal: %w", err)
	}

	content := &pb.Content{}
	if err := content.Unmarshal(pd.Content); err != nil {
		return ret, fmt.Errorf("ibtp content unmarshal: %w", err)
	}

	args := make([][]byte, 0)
	args = append(args,
		[]byte(ibtp.From),
		[]byte(strconv.FormatUint(ibtp.Index, 10)),
		[]byte(content.DstContractId),
	)
	args = append(args, content.Args...)
	if ibtp.Type == pb.IBTP_ROLLBACK {
		// use an unexist contract address, so only inCounter will be increased
		args[2] = []byte("0x0000000000000000000000000000000000ffffff")
	}
	if pb.IBTP_ASSET_EXCHANGE_REDEEM == ibtp.Type || pb.IBTP_ASSET_EXCHANGE_REFUND == ibtp.Type {
		args = append(args, ibtp.Extra)
	}
	resultArgs, err := solidity.ABIUnmarshal(c.abi, args, content.Func)
	if err != nil {
		return ret, fmt.Errorf("unmarshal ibtp function abi args %w", err)
	}

	var (
		tx *types.Transaction
	)
	if err := retry.Retry(func(attempt uint) error {
		tx, err = c.broker.BrokerTransactor.contract.Transact(&c.session.TransactOpts, content.Func, resultArgs...)
		if err != nil {
			logger.Info("Invoke contract failed",
				"func", content.Func,
				"args", string(bytes.Join(content.Args, []byte(","))),
				"error", err,
			)
			return err
		}

		return nil
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't invoke contract", "error", err)
	}

	// get transaction result from subscription
	var (
		result [][]byte
		status bool
	)

	receipt := c.waitForMined(tx.Hash())

	for _, lg := range receipt.Logs {
		switch eventSigs[lg.Topics[0].Hex()] {
		case "logInterchainData":
			out, err := c.broker.BrokerFilterer.ParseLogInterchainData(*lg)
			if err != nil {
				return ret, fmt.Errorf("unpack log: %w", err)
			}
			status = out.Status
			result = append(result, []byte(strconv.FormatBool(out.Status)), []byte(out.Data))
		case "logInterchainStatus":
			out, err := c.broker.BrokerFilterer.ParseLogInterchainStatus(*lg)
			if err != nil {
				return ret, fmt.Errorf("unpack log: %w", err)
			}

			result = append(result, content.Args[2:len(content.Args)-1]...)
			status = out.Status
			//default:
			//	return ret, fmt.Errorf("unsupported method")
		}
	}

	ret.Status = status
	// If no callback function To invoke, then simply return
	if content.Callback == "" {
		return ret, nil
	}

	responseStatus := true
	// execution invoke no err
	var newArgs = make([][]byte, 0)
	switch content.Func {
	case "interchainGet":
		newArgs = append(newArgs, content.Args[0])
		newArgs = append(newArgs, result...)
	case "interchainCharge":
		newArgs = append(newArgs, []byte(strconv.FormatBool(status)), content.Args[0])
		newArgs = append(newArgs, content.Args[2:len(content.Args)-1]...)
		responseStatus = status
	case "interchainAssetExchangeRedeem":
		newArgs = append(newArgs, args[3:]...)
	case "interchainAssetExchangeRefund":
		newArgs = append(newArgs, args[3:]...)
	default:
		newArgs = append(newArgs, result...)
	}

	ret.Result, err = c.generateCallback(ibtp, newArgs, responseStatus)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

// GetOutMessage gets crosschain tx by `to` address and index
func (c *Client) GetOutMessage(to string, idx uint64) (*pb.IBTP, error) {
	blockNum, err := c.session.GetOutMessage(common.HexToAddress(to), idx)
	if err != nil {
		return nil, err
	}
	fmt.Println(blockNum)

	return c.toIBTP(blockNum, idx)
}

func (c *Client) toIBTP(blockNum *big.Int, idx uint64) (*pb.IBTP, error) {
	var (
		block *types.Block
		err   error
		ibtp  *pb.IBTP
	)

	if err := retry.Retry(func(attempt uint) error {
		block, err = c.ethClient.BlockByNumber(c.ctx, blockNum)
		if err != nil {
			return err
		}

		return nil
	}, strategy.Wait(1*time.Second)); err != nil {
		logger.Error("Query block by number failed", "error", err)
	}

	txs := block.Transactions()
	for _, tx := range txs {
		receipt, err := c.ethClient.TransactionReceipt(c.ctx, tx.Hash())
		if err != nil {
			logger.Info("get receipt error")
			continue
		}

		for _, lg := range receipt.Logs {
			if eventSigs[lg.Topics[0].Hex()] != "interchainEvent" {
				continue
			}
			ev, err := c.broker.BrokerFilterer.ParseThrowEvent(*lg)
			if err != nil {
				return nil, fmt.Errorf("unpack log: %w", err)
			}
			if ev.Index != idx {
				continue
			}
			ibtp = Convert2IBTP(ev, c.pierID, pb.IBTP_INTERCHAIN)
			break
		}
	}

	if ibtp == nil {
		return nil, fmt.Errorf("can't find historical event")
	}
	return ibtp, nil
}

func (c *Client) getMeta(method func() ([]common.Address, []uint64, error)) (map[string]uint64, error) {
	var (
		addresses []common.Address
		indices   []uint64
		err       error
	)
	meta := make(map[string]uint64, 0)

	addresses, indices, err = method()
	if err != nil {
		return nil, err
	}

	for i, addr := range addresses {
		meta[strings.ToLower(addr.Hex())] = indices[i]
	}

	return meta, nil
}

// GetInMessage gets the execution results from contract by from-index key
func (c *Client) GetInMessage(from string, idx uint64) ([][]byte, error) {
	blockNum, err := c.session.GetInMessage(common.HexToAddress(from), idx)
	if err != nil {
		return nil, err
	}

	fmt.Println(blockNum)

	return nil, nil
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

func (c *Client) RollbackIBTP(ibtp *pb.IBTP, isSrcChain bool) (*pb.RollbackIBTPResponse, error) {
	ret := &pb.RollbackIBTPResponse{}
	pd := &pb.Payload{}
	if err := pd.Unmarshal(ibtp.Payload); err != nil {
		return nil, fmt.Errorf("ibtp payload unmarshal: %w", err)
	}
	content := &pb.Content{}
	if err := content.Unmarshal(pd.Content); err != nil {
		return ret, fmt.Errorf("ibtp content unmarshal: %w", err)
	}

	// only support rollback for interchainCharge
	if "interchainCharge" != content.Func {
		return nil, nil
	}

	var rollbackArgs [][]byte
	var rollbackFunc string
	var args [][]byte
	if isSrcChain {
		rollbackFunc = content.Callback
		rollbackArgs = append(rollbackArgs, []byte("false"), content.Args[0], content.Args[2])

		args = append(args, []byte(ibtp.To),
			[]byte(strconv.FormatUint(ibtp.Index, 10)),
			[]byte(strings.ToLower(content.SrcContractId)))
		args = append(args, rollbackArgs...)
	} else {
		rollbackFunc = content.Func
		args = append(args, []byte(ibtp.From),
			[]byte(strconv.FormatUint(ibtp.Index, 10)),
			[]byte(strings.ToLower(content.DstContractId)))
		args = append(args, content.Args...)
		args[len(args)-1] = []byte("true")
	}

	resultArgs, err := solidity.ABIUnmarshal(c.abi, args, rollbackFunc)
	if err != nil {
		return ret, fmt.Errorf("unmarshal ibtp function abi args %w", err)
	}

	var (
		tx *types.Transaction
	)
	if err := retry.Retry(func(attempt uint) error {
		tx, err = c.broker.BrokerTransactor.contract.Transact(&c.session.TransactOpts, rollbackFunc, resultArgs...)
		if err != nil {
			logger.Info("Invoke contract failed",
				"func", rollbackFunc,
				"args", string(bytes.Join(args, []byte(","))),
				"error", err,
			)
			return err
		}

		return nil
	}, strategy.Wait(2*time.Second)); err != nil {
		logger.Error("Can't invoke contract", "error", err)
	}

	receipt := c.waitForMined(tx.Hash())
	if len(receipt.Logs) == 0 {
		ret.Status = false
		ret.Message = "wrong contract doesn't emit log event"
	}
	for _, lg := range receipt.Logs {
		switch eventSigs[lg.Topics[0].Hex()] {
		case "logInterchainStatus":
			out, err := c.broker.BrokerFilterer.ParseLogInterchainStatus(*lg)
			if err != nil {
				return ret, fmt.Errorf("unpack log: %w", err)
			}
			ret.Status = out.Status
		}
	}

	return ret, nil
}

func (c *Client) GetSrcRollbackMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetSrcRollbackMeta)
}

func (c *Client) GetDstRollbackMeta() (map[string]uint64, error) {
	return c.getMeta(c.session.GetDstRollbackMeta)
}

func (c *Client) CommitCallback(ibtp *pb.IBTP) error {
	return nil
}

func (c *Client) waitForMined(hash common.Hash) *types.Receipt {
	var (
		receipt *types.Receipt
		err     error
	)
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
