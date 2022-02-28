package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

	IBTPid := fmt.Sprintf("%s-%s-%d", ev.SrcFullID, ev.DstFullID, ev.Index)
	timestamp, _, err := c.GetTransactionMeta(IBTPid)
	if err != nil {
		return nil, err
	}
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, timestamp)

	return &pb.IBTP{
		From:          ev.SrcFullID,
		To:            ev.DstFullID,
		Index:         ev.Index,
		Type:          pb.IBTP_INTERCHAIN,
		TimeoutHeight: timeoutHeight,
		Proof:         []byte("1"),
		Payload:       pd,
		Extra:         buf,
	}, nil
}

func (c *Client) Convert2Receipt(ev *BrokerThrowReceiptEvent) (*pb.IBTP, error) {
	fullEv, encrypt, err := c.fillReceiptEvent(ev)
	if err != nil {
		return nil, err
	}

	return generateReceipt(fullEv.SrcFullID, fullEv.DstFullID, fullEv.Index, fullEv.Result, fullEv.Typ, encrypt)
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
	if ev.Func == "" {
		fun, args, encrypt, err := c.session.GetOutMessage(pb.GenServicePair(ev.SrcFullID, ev.DstFullID), ev.Index)
		if err != nil {
			return nil, false, err
		}

		ev.Func = fun
		ev.Args = args
		emptyHash := common.Hash{}

		if bytes.Equal(ev.Hash[:], emptyHash[:]) {
			var data []byte
			data = append(data, []byte(fun)...)
			for _, arg := range args {
				data = append(data, []byte(arg)...)
			}

			ev.Hash = common.BytesToHash(crypto.Keccak256(data))
		}

		return ev, encrypt, nil
	}

	return ev, false, nil
}

func (c *Client) fillReceiptEvent(ev *BrokerThrowReceiptEvent) (*BrokerThrowReceiptEvent, bool, error) {
	if ev.Result == nil {
		result, typ, encrypt, err := c.session.GetReceiptMessage(pb.GenServicePair(ev.SrcFullID, ev.DstFullID), ev.Index)
		if err != nil {
			return nil, false, err
		}
		ev.Result = result
		ev.Typ = typ

		emptyHash := common.Hash{}
		if bytes.Equal(ev.Hash[:], emptyHash[:]) {
			var packed []byte
			for _, ele := range result {
				packed = append(packed, ele...)
			}
			ev.Hash = common.BytesToHash(crypto.Keccak256(packed))
		}

		return ev, encrypt, nil
	}

	return ev, false, nil
}
