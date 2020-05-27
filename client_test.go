package main

import (
	"testing"
	"time"

	"github.com/meshplus/bitxhub-model/pb"
	"github.com/stretchr/testify/require"
)

const (
	from     = "0x3c1adfbe8b4dd45e8e9bcad6216a236416dcd6a9"
	to       = "0xe02d8fdacd59020d7f292ab3278d13674f5c404d"
	dstAddr  = "0x6aF7CC8B895e5FCe9e9B3571B437A01795E3B93e"
	fromAddr = "0x55f34bd04b2ea8047b46c2d891b1c2a0aff15102"
)

func TestClient_Start(t *testing.T) {
	cli, err := NewClient("/Users/taoyongxing/.pier/ether", from, nil)
	require.Nil(t, err)

	require.Nil(t, cli.Start())
	ibtp := getIBTP(t, 2, pb.IBTP_INTERCHAIN)

	res, err := cli.SubmitIBTP(ibtp)
	require.Nil(t, err)
	logger.Info("returned ibtp:", res)
	time.Sleep(1000 * time.Second)
}

func Test(t *testing.T) {
	cli, err := NewClient("/Users/taoyongxing/.pier/ether", from, nil)
	require.Nil(t, err)

	require.Nil(t, cli.Start())
	//ibtp := getIBTP(t, 1, pb.IBTP_INTERCHAIN)

	res, err := cli.GetInMeta()
	require.Nil(t, err)
	logger.Info("returned ibtp:", res)
	time.Sleep(1000 * time.Second)
}

func getIBTP(t *testing.T, index uint64, typ pb.IBTP_Type) *pb.IBTP {
	ct := &pb.Content{
		SrcContractId: fromAddr,
		DstContractId: dstAddr,
		Func:          "interchainCharge",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "interchainConfirm",
	}
	c, err := ct.Marshal()
	require.Nil(t, err)

	pd := pb.Payload{
		Encrypted: false,
		Content:   c,
	}
	ibtppd, err := pd.Marshal()
	require.Nil(t, err)

	return &pb.IBTP{
		From:      from,
		To:        to,
		Payload:   ibtppd,
		Index:     index,
		Type:      typ,
		Timestamp: time.Now().UnixNano(),
	}
}
