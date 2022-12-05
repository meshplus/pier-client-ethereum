package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

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
	group, err := c.fillGroup(fullEv.Group)
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
		Group:         group,
	}, nil
}

func (c *Client) fillGroup(evGroup []string) (*pb.StringUint64Map, error) {
	if len(evGroup) == 0 {
		return nil, nil
	}
	keys := make([]string, len(evGroup))
	values := make([]uint64, len(evGroup))
	for i, str := range evGroup {
		group := strings.Split(str, "-")
		if len(group) != 2 {
			return nil, fmt.Errorf("input group err: %s", str)
		}
		keys[i] = group[0]
		v, err := strconv.Atoi(group[1])
		if err != nil {
			return nil, fmt.Errorf("convert group value %s err:%s", group[i], err)
		}
		values[i] = uint64(v)
	}
	return &pb.StringUint64Map{
		Keys: keys,
		Vals: values,
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
		fun, args, encrypt, group, err := c.session.GetOutMessage(pb.GenServicePair(ev.SrcFullID, ev.DstFullID), ev.Index)
		if err != nil {
			return nil, false, err
		}

		ev.Func = fun
		ev.Args = args
		ev.Group = group
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
