package main

import (
	"bytes"
	"encoding/binary"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) Convert2IBTP(ev *BrokerThrowInterchainEvent, timeoutHeight int64) (*pb.IBTP, error) {
	fullEv, encrypt, err := c.fillInterchainEvent(ev)
	if err != nil {
		return nil, err
	}
	pd, err := encodePayload(fullEv, encrypt)
	if err != nil {
		return nil, err
	}

	return &pb.IBTP{
		From:          ev.SrcFullID,
		To:            ev.DstFullID,
		Index:         ev.Index,
		Type:          pb.IBTP_INTERCHAIN,
		TimeoutHeight: timeoutHeight,
		Proof:         []byte("1"),
		Payload:       pd,
	}, nil
}

func encodePayload(ev *BrokerThrowInterchainEvent, encrypt bool) ([]byte, error) {
	var args [][]byte
	for _, arg := range ev.Args {
		args = append(args, []byte(arg))
	}
	content := &pb.Content{
		Func: ev.Func,
		Args: args,
	}
	data, err := content.Marshal()
	if err != nil {
		return nil, err
	}

	ibtppd := &pb.Payload{
		Encrypted: encrypt,
		Content:   data,
		Hash:      ev.Hash[:],
	}
	return ibtppd.Marshal()
}

func (c *Client) fillInterchainEvent(ev *BrokerThrowInterchainEvent) (*BrokerThrowInterchainEvent, bool, error) {
	ev.Func = "interchainCharge"
	var args [][]byte
	args = append(args, []byte("alice"))
	args = append(args, []byte("bob"))
	args = append(args, IntToBytes(10))
	ev.Args = args
	emptyHash := common.Hash{}

	if bytes.Equal(ev.Hash[:], emptyHash[:]) {
		var data []byte
		data = append(data, []byte(ev.Func)...)
		for _, arg := range args {
			data = append(data, []byte(arg)...)
		}

		ev.Hash = common.BytesToHash(crypto.Keccak256(data))
	}

	return ev, false, nil
}

func IntToBytes(n int) []byte {
	x := uint64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
