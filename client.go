package main

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/pier/pkg/plugins"
)

type Client struct {
	config         *Config
	ctx            context.Context
	cancel         context.CancelFunc
	eventC         chan *pb.IBTP
	interchainInfo *Interchain
	lock           *sync.RWMutex
}

type Interchain struct {
	outCounter      map[string]uint64
	callbackCounter map[string]uint64
	inCounter       map[string]uint64
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

func (c *Client) Initialize(configPath string, extra []byte, mode string) error {
	cfg, err := UnmarshalConfig(configPath)
	if err != nil {
		return fmt.Errorf("unmarshal config for plugin :%w", err)
	}

	logger.Info("Basic appchain info",
		"bxhId", cfg.Mock.BxhId,
		"chainId", cfg.Mock.ChainId,
		"serviceList", cfg.Mock.ServiceList)

	c.config = cfg
	c.interchainInfo = &Interchain{
		outCounter:      make(map[string]uint64),
		inCounter:       make(map[string]uint64),
		callbackCounter: make(map[string]uint64),
	}
	c.lock = &sync.RWMutex{}
	c.eventC = make(chan *pb.IBTP, 1024)
	c.ctx, c.cancel = context.WithCancel(context.Background())

	server, err := NewServer(c.config.Mock.Port, c)
	if err != nil {
		return err
	}

	if err := server.Start(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Start() error {
	return nil
}

func (c *Client) Stop() error {
	c.cancel()
	return nil
}

func (c *Client) GetIBTPCh() chan *pb.IBTP {
	return c.eventC
}

func (c *Client) Name() string {
	return c.config.Mock.Name
}

func (c *Client) Type() string {
	return EtherType
}

// SubmitIBTP submit interchain ibtp. It will unwrap the ibtp and execute
// the function inside the ibtp. If any execution results returned, pass
// them to other modules.
func (c *Client) SubmitIBTP(from string, index uint64, serviceID string, ibtpType pb.IBTP_Type, content *pb.Content, proof *pb.BxhProof, isEncrypted bool) (*pb.SubmitIBTPResponse, error) {
	ret := &pb.SubmitIBTPResponse{Status: true}

	to := fmt.Sprintf("%s:%s:%s", c.config.Mock.BxhId, c.config.Mock.ChainId, serviceID)
	var result [][]byte
	result = append(result, []byte("0"))
	receipt, _ := generateReceipt(from, to, index, result, uint64(pb.IBTP_RECEIPT_SUCCESS), false)
	c.eventC <- receipt

	servicePair := fmt.Sprintf("%s-%s", from, to)
	c.lock.Lock()
	c.interchainInfo.inCounter[servicePair]++
	c.lock.Unlock()

	return ret, nil
}

func (c *Client) SubmitReceipt(to string, index uint64, serviceID string, ibtpType pb.IBTP_Type, result *pb.Result, proof *pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	ret := &pb.SubmitIBTPResponse{Status: true}

	from := fmt.Sprintf("%s:%s:%s", c.config.Mock.BxhId, c.config.Mock.ChainId, serviceID)
	servicePair := fmt.Sprintf("%s-%s", from, to)
	c.lock.Lock()
	c.interchainInfo.callbackCounter[servicePair]++
	c.lock.Unlock()

	return ret, nil
}

func (c *Client) SubmitIBTPBatch(from []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, content []*pb.Content, proof []*pb.BxhProof, isEncrypted []bool) (*pb.SubmitIBTPResponse, error) {
	ret := &pb.SubmitIBTPResponse{Status: true}

	for idx, src := range from {
		to := fmt.Sprintf("%s:%s:%s", c.config.Mock.BxhId, c.config.Mock.ChainId, serviceID[idx])
		var result [][]byte
		result = append(result, []byte("0"))
		receipt, _ := generateReceipt(src, to, index[idx], result, uint64(pb.IBTP_RECEIPT_SUCCESS), false)
		c.eventC <- receipt

		servicePair := fmt.Sprintf("%s-%s", src, to)
		c.lock.Lock()
		c.interchainInfo.inCounter[servicePair]++
		c.lock.Unlock()
	}

	return ret, nil
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

	return c.Convert2IBTP(ev, int64(c.config.Mock.TimeoutHeight))
}

// GetInMessage gets the execution results from contract by from-index key
func (c *Client) GetReceiptMessage(servicePair string, idx uint64) (*pb.IBTP, error) {
	var result [][]byte
	result = append(result, []byte("0"))

	srcServiceID, dstServiceID, err := pb.ParseServicePair(servicePair)
	if err != nil {
		return nil, err
	}

	return generateReceipt(srcServiceID, dstServiceID, idx, result, uint64(pb.IBTP_RECEIPT_SUCCESS), false)
}

// GetInMeta queries contract about how many interchain txs have been
// executed on this appchain for different source chains.
func (c *Client) GetInMeta() (map[string]uint64, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.interchainInfo.inCounter, nil
}

// GetOutMeta queries contract about how many interchain txs have been
// sent out on this appchain to different destination chains.
func (c *Client) GetOutMeta() (map[string]uint64, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.interchainInfo.outCounter, nil
}

// GetCallbackMeta queries contract about how many callback functions have been
// executed on this appchain from different destination chains.
func (c *Client) GetCallbackMeta() (map[string]uint64, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.interchainInfo.callbackCounter, nil
}

func (c *Client) GetDstRollbackMeta() (map[string]uint64, error) {
	return make(map[string]uint64), nil
}

func (c *Client) GetChainID() (string, string, error) {
	return c.config.Mock.BxhId, c.config.Mock.ChainId, nil
}

func (c *Client) GetServices() ([]string, error) {
	var services []string
	for _, service := range c.config.Mock.ServiceList {
		serviceFullId := fmt.Sprintf("%s:%s:%s", c.config.Mock.BxhId, c.config.Mock.ChainId, service)
		services = append(services, serviceFullId)
	}

	return services, nil
}

func (c *Client) GetDirectTransactionMeta(IBTPid string) (uint64, uint64, uint64, error) {
	panic("implement me")
}

func (c *Client) GetAppchainInfo(chainID string) (string, []byte, string, error) {
	panic("implement me")
}

func (c *Client) GetOffChainData(request *pb.GetDataRequest) (*pb.GetDataResponse, error) {
	panic("implement me")
}

func (c *Client) GetOffChainDataReq() chan *pb.GetDataRequest {
	panic("implement me")
}

func (c *Client) SubmitOffChainData(response *pb.GetDataResponse) error {
	panic("implement me")
}

func (c *Client) GetUpdateMeta() chan *pb.UpdateMeta {
	panic("implement me")
}

func (c *Client) SubmitReceiptBatch(to []string, index []uint64, serviceID []string, ibtpType []pb.IBTP_Type, result []*pb.Result, proof []*pb.BxhProof) (*pb.SubmitIBTPResponse, error) {
	panic("implement me")
}
