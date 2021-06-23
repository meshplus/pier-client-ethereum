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
	"github.com/ethereum/go-ethereum/trie"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	contracts "github.com/meshplus/bitxhub-core/eth-contracts"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/bitxid"
	"github.com/meshplus/pier-client-ethereum/solidity"
	"github.com/meshplus/pier/pkg/plugins"
)

//go:generate abigen --sol ./example/broker.sol --pkg main --out broker.go
//go:generate abigen --sol ./example/interchain.sol --pkg main --out interchain.go
//go:generate abigen --sol ./example/escrows.sol --pkg main --out escrows.go
type Client struct {
	abi            abi.ABI
	config         *Config
	ctx            context.Context
	ethClient      *ethclient.Client
	broker         *Broker
	brokerSession  *BrokerSession
	escrowsSession *contracts.EscrowsSession
	conn           *rpc.Client
	eventC         chan *pb.IBTP
	metaC          chan *pb.UpdateMeta
	filterOptCh    chan *bind.FilterOpts
	logCh          chan *contracts.EscrowsLock
	lockCh         chan *pb.LockEvent
	preLockCh      chan *PreLockEvent
	appchainID     string // the method of connected appchain like did:bitxhub:appchain:.
	bizABI         map[string]*abi.ABI
	headerPool     *headerPool
}

const (
	EtherType        = "ethereum"
	InvokeInterchain = "invokeInterchain"
	Threshold        = 20
	MintEventName    = "Mint"
)

var (
	_ plugins.Client = (*Client)(nil)

	eventSig = map[string]string{
		"0x4feaa67f2bfe27f3f037662df125ebe5bbba60fe4cbab27b0fc61b13c44789f8": MintEventName,
	}
	logger = hclog.New(&hclog.LoggerOptions{
		Name:   "client",
		Output: os.Stderr,
		Level:  hclog.Trace,
	})
)

func (c *Client) Initialize(configPath string, appchainID string, extra []byte) error {
	cfg, err := UnmarshalConfig(configPath)
	if err != nil {
		return fmt.Errorf("unmarshal config for plugin :%w", err)
	}

	logger.Info("Basic appchain info",
		"broker address", cfg.Ether.ContractAddress,
		"ethereum node ip", cfg.Ether.Addr)

	currentHeight, err := strconv.ParseUint(string(extra), 10, 64)
	if err != nil {
		return fmt.Errorf("parse current height from extra: %w", err)
	}
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
	brokerSession := &BrokerSession{
		Contract: broker,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}
	escrowsContract, err := contracts.NewEscrows(common.HexToAddress(cfg.Ether.EscrowsAddress), etherCli)
	if err != nil {
		return fmt.Errorf("failed to instantiate a Broker contract: %w", err)
	}
	escrowsSession := &contracts.EscrowsSession{
		Contract: escrowsContract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}
	c.headerPool = newHeaderPool(currentHeight)

	ab, err := abi.JSON(bytes.NewReader([]byte(BrokerABI)))
	if err != nil {
		return fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	conn, err := rpc.Dial(cfg.Ether.Addr)
	if err != nil {
		return fmt.Errorf("rpc dial: %s", err.Error())
	}

	bizAbi := make(map[string]*abi.ABI)
	for addr, name := range cfg.ContractABI {
		logger.Info("ContractABI", "addr", addr, "name", name)
		content, err := ioutil.ReadFile(filepath.Join(configPath, name))
		if err != nil {
			return fmt.Errorf("read abi %s for addr %s: %w", name, addr, err)
		}
		abi, err := abi.JSON(bytes.NewReader(content))
		if err != nil {
			return fmt.Errorf("unmarshal abi %s for addr %s: %s", name, addr, err.Error())
		}
		bizAbi[addr] = &abi
	}

	c.config = cfg
	c.eventC = make(chan *pb.IBTP, 1024)
	c.metaC = make(chan *pb.UpdateMeta, 1024)
	c.filterOptCh = make(chan *bind.FilterOpts, 1024)
	c.logCh = make(chan *contracts.EscrowsLock, 1024)
	c.lockCh = make(chan *pb.LockEvent, 1024)
	c.preLockCh = make(chan *PreLockEvent, 1024)
	c.ethClient = etherCli
	c.broker = broker
	c.brokerSession = brokerSession
	c.escrowsSession = escrowsSession
	c.abi = ab
	c.conn = conn
	c.appchainID = appchainID
	c.ctx = context.Background()
	c.bizABI = bizAbi

	return nil
}

func (c *Client) Start() error {
	go c.listenHeader()
	go c.postHeaders()
	go c.listenLock()
	go c.postLock()
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

func (c *Client) GetUpdateMeta() <-chan *pb.UpdateMeta {
	return c.metaC
}

func (c *Client) GetLockEvent() <-chan *pb.LockEvent {
	return c.lockCh
}

func (c *Client) Unescrow(unlock *pb.UnLock) error {
	_, err := c.escrowsSession.Unlock(
		common.HexToAddress(unlock.Token),
		common.HexToAddress(unlock.From),
		common.HexToAddress(unlock.Receipt),
		new(big.Int).SetBytes(unlock.Amount),
		unlock.TxId,
		new(big.Int).SetUint64(unlock.RelayIndex),
		unlock.GetMultiSigns())
	if err != nil {
		logger.Error("unescrow", "err", err.Error())
		return err
	}
	//logger.Info("unescrow", "tx-hash", transaction.Hash().Hex())
	//status := c.getTxReceipt(transaction.Hash()).Status
	//logger.Info("unescrow", "txReceipt", status)
	return nil
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

	if ibtp.Category() == pb.IBTP_UNKNOWN {
		return nil, fmt.Errorf("invalid ibtp category")
	}

	var result [][]byte
	logger.Info("submit ibtp", "id", ibtp.ID(), "contract", content.DstContractId, "func", content.Func)
	for i, arg := range content.Args {
		logger.Info("arg", strconv.Itoa(i), string(arg))
	}

	if ibtp.Category() == pb.IBTP_RESPONSE && content.Func == "" {
		logger.Info("InvokeIndexUpdate", "ibtp", ibtp.ID())
		success, err := c.InvokeIndexUpdate(ibtp.From, ibtp.Index, ibtp.Category())
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("update index for ibtp %s failed", ibtp.ID())
		}
		ret.Status = true
		return ret, nil
	}

	var bizData []byte
	var err error
	bAbi, ok := c.bizABI[strings.ToLower(content.DstContractId)]
	if !ok {
		ret.Message = fmt.Sprintf("no abi for contract %s", content.DstContractId)
	} else {
		bizData, err = c.packFuncArgs(content.Func, content.Args, bAbi)
		if err != nil {
			ret.Message = fmt.Sprintf("pack for ibtp %s func %s and args: %s", ibtp.ID(), content.Func, err)
		}
	}
	if !ok || err != nil {
		ret.Status = false
		success, err := c.InvokeIndexUpdateWithError(ibtp.From, ibtp.Index, ibtp.Category(), ret.Message)
		if err != nil {
			return nil, err
		}
		if !success {
			return nil, fmt.Errorf("update index for ibtp %s failed", ibtp.ID())
		}
	} else {
		receipt, err := c.InvokeInterchain(ibtp.From, ibtp.Index, content.DstContractId, ibtp.Category(), bizData)
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
			success, err := c.InvokeIndexUpdateWithError(ibtp.From, ibtp.Index, ibtp.Category(), ret.Message)
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

func (c *Client) InvokeInterchain(srcChainMethod string, index uint64, destAddr string, category pb.IBTP_Category, bizCallData []byte) (*types.Receipt, error) {
	var tx *types.Transaction
	var txErr error
	if err := retry.Retry(func(attempt uint) error {
		tx, txErr = c.brokerSession.InvokeInterchain(srcChainMethod, index, common.HexToAddress(destAddr), category == pb.IBTP_REQUEST, bizCallData)
		if txErr != nil {
			logger.Info("Call InvokeInterchain failed",
				"srcChainMethod", srcChainMethod,
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

func (c *Client) InvokeIndexUpdateWithError(srcChainMethod string, index uint64, category pb.IBTP_Category, errMsg string) (bool, error) {
	var tx *types.Transaction
	var txErr error
	if err := retry.Retry(func(attempt uint) error {
		tx, txErr = c.brokerSession.InvokeIndexUpdateWithError(srcChainMethod, index, category == pb.IBTP_REQUEST, errMsg)
		if txErr != nil {
			logger.Info("Call InvokeIndexUpdateWithError failed",
				"srcChainMethod", srcChainMethod,
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

func (c *Client) InvokeIndexUpdate(srcChainID string, index uint64, category pb.IBTP_Category) (bool, error) {
	return c.InvokeIndexUpdateWithError(srcChainID, index, category, "")
}

// GetOutMessage gets crosschain tx by `to` address and index
func (c *Client) GetOutMessage(to string, idx uint64) (*pb.IBTP, error) {
	var blockNum *big.Int
	if err := retry.Retry(func(attempt uint) error {
		var err error
		blockNum, err = c.brokerSession.GetOutMessage(to, idx)
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

	for throwEvents.Next() {
		ev := throwEvents.Event
		if string(bitxid.DID(ev.DestDID).GetChainDID()) == to && ev.Index == idx {
			return Convert2IBTP(ev, c.appchainID, pb.IBTP_INTERCHAIN), nil
		}
	}

	return nil, fmt.Errorf("cannot find out ibtp for to %s and index %d", to, idx)
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
		blockNum, err = c.brokerSession.GetInMessage(from, idx)
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
	return make(map[string]uint64, 0), nil
}

// GetOutMeta queries contract about how many interchain txs have been
// sent out on this appchain to different destination chains.
func (c *Client) GetOutMeta() (map[string]uint64, error) {
	return make(map[string]uint64, 0), nil
}

// GetCallbackMeta queries contract about how many callback functions have been
// executed on this appchain from different destination chains.
func (c *Client) GetCallbackMeta() (map[string]uint64, error) {
	return make(map[string]uint64, 0), nil
}

func (c *Client) CommitCallback(ibtp *pb.IBTP) error {
	return nil
}

// @ibtp is the original ibtp merged from this appchain
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
	if content.Func != "interchainCharge" {
		return nil, nil
	}

	var bizData []byte
	var err error
	bAbi, ok := c.bizABI[strings.ToLower(content.SrcContractId)]
	if !ok {
		ret.Message = fmt.Sprintf("no abi for contract %s", content.SrcContractId)
	} else {
		bizData, err = c.packFuncArgs(content.Rollback, content.ArgsRb, bAbi)
		if err != nil {
			ret.Message = fmt.Sprintf("pack for ibtp %s func %s and args: %s", ibtp.ID(), content.Func, err)
		}
	}

	// false indicates it is for rollback
	tx, err := c.brokerSession.InvokeInterchain(ibtp.To, ibtp.Index, common.HexToAddress(content.SrcContractId), false, bizData)
	if err != nil {
		return nil, err
	}
	receipt := c.waitForConfirmed(tx.Hash())
	if len(receipt.Logs) == 0 {
		ret.Status = false
		ret.Message = "wrong contract doesn't emit log event"
	}

	return ret, nil
}

func (c *Client) IncreaseInMeta(original *pb.IBTP) (*pb.IBTP, error) {
	ibtp, err := c.generateCallback(original, nil, false)
	if err != nil {
		return nil, err
	}
	errMsg := "ibtp failed in bitxhub"
	success, err := c.InvokeIndexUpdateWithError(original.From, original.Index, original.Category(), errMsg)
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
	blockBytes, err := c.GetInMessage(ibtp.From, ibtp.Index)
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

	bAbi, ok := c.bizABI[strings.ToLower(ct.DstContractId)]
	if !ok {
		return nil, fmt.Errorf("can not find abi for contract %s", ct.DstContractId)
	}
	bizData, err := c.packFuncArgs(ct.Func, ct.Args, bAbi)
	packData, err := c.abi.Pack(InvokeInterchain, ibtp.From, ibtp.Index, common.HexToAddress(ct.DstContractId), ibtp.Category() == pb.IBTP_REQUEST, bizData)
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

func (c *Client) QueryFilterLockStart(appchainIndex int64) int64 {
	res, err := c.escrowsSession.Index2Height(big.NewInt(appchainIndex))
	if err != nil {
		return 0
	}
	return res.Int64()
}

func (c *Client) QueryLockEventByIndex(index int64) *pb.LockEvent {
	var lockCh *pb.LockEvent
	height, er := c.escrowsSession.Index2Height(big.NewInt(index))
	if er != nil {
		return nil
	}
	end := height.Uint64()
	filterOpt := &bind.FilterOpts{
		Start: end,
		End:   &end,
	}

	var (
		iter *contracts.EscrowsLockIterator
		err  error
	)
	if err := retry.Retry(func(attempt uint) error {
		iter, err = c.escrowsSession.Contract.FilterLock(filterOpt)
		if err != nil {
			return err
		}
		return nil
	}, strategy.Wait(1*time.Second)); err != nil {
		logger.Error("Can't get filter mint event", "error", err.Error())
	}
	for iter.Next() {
		if index != iter.Event.AppchainIndex.Int64() {
			continue
		}
		raw := &iter.Event.Raw
		// query this block from ethereum and generate mintEvent and proof for pier
		if err := retry.Retry(func(attempt uint) error {
			block, err := c.ethClient.BlockByNumber(c.ctx, big.NewInt(int64(raw.BlockNumber)))
			if err != nil {
				return err
			}
			receipt, err := c.ethClient.TransactionReceipt(c.ctx, raw.TxHash)
			if err != nil {
				return err
			}
			// construct receipt merkle tree first for proof generate
			receipts := make([]*types.Receipt, 0, len(block.Transactions()))
			for _, tx := range block.Transactions() {
				receipt, err := c.ethClient.TransactionReceipt(c.ctx, tx.Hash())
				if err != nil {
					return err
				}
				receipts = append(receipts, receipt)
			}
			receiptsTrie := new(trie.Trie)
			tReceipts := types.Receipts(receipts)
			types.DeriveSha(tReceipts, receiptsTrie)

			receiptData, err := receipt.MarshalJSON()
			if err != nil {
				return err
			}
			proof, err := c.getProof(receiptsTrie, uint64(receipt.TransactionIndex))
			if err != nil {
				return err
			}
			lockCh = &pb.LockEvent{
				AppchainIndex: iter.Event.AppchainIndex.Uint64(),
				Receipt:       receiptData,
				Proof:         proof,
			}
			return nil
		}, strategy.Wait(1*time.Second)); err != nil {
			logger.Error("Can't retrieve mint event from receipt", "error", err.Error())
		}
	}
	return lockCh
}

func (c *Client) QueryAppchainIndex() int64 {
	res, err := c.escrowsSession.AppchainIndex()
	if err != nil {
		return 0
	}
	return res.Int64()
}
func (c *Client) QueryRelayIndex() int64 {
	res, err := c.escrowsSession.RelayIndex()
	if err != nil {
		return 0
	}
	return res.Int64()
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
