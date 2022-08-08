package main

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) Convert2DirectIBTP(ev *BrokerDirectThrowInterchainEvent, timeoutHeight int64) (*pb.IBTP, error) {
	fullEv, encrypt, err := c.fillDirectInterchainEvent(ev)
	if err != nil {
		return nil, err
	}
	pd, err := encodeDirectPayload(fullEv, encrypt)
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

func (c *Client) Convert2DirectReceipt(ev *BrokerDirectThrowReceiptEvent) (*pb.IBTP, error) {
	fullEv, encrypt, err := c.fillDirectReceiptEvent(ev)
	if err != nil {
		return nil, err
	}

	return generateDirectReceipt(fullEv.SrcFullID, fullEv.DstFullID, fullEv.Index, fullEv.Result, fullEv.Typ, encrypt)
}

func encodeDirectPayload(ev *BrokerDirectThrowInterchainEvent, encrypt bool) ([]byte, error) {
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

func (c *Client) fillDirectInterchainEvent(ev *BrokerDirectThrowInterchainEvent) (*BrokerDirectThrowInterchainEvent, bool, error) {
	if ev.Func == "" {
		fun, args, encrypt, err := c.sessionDirect.GetOutMessage(pb.GenServicePair(ev.SrcFullID, ev.DstFullID), ev.Index)
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

func (c *Client) fillDirectReceiptEvent(ev *BrokerDirectThrowReceiptEvent) (*BrokerDirectThrowReceiptEvent, bool, error) {
	if ev.Result == nil {
		result, typ, encrypt, err := c.sessionDirect.GetReceiptMessage(pb.GenServicePair(ev.SrcFullID, ev.DstFullID), ev.Index)
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

func generateDirectReceipt(from, to string, idx uint64, data [][]byte, typ uint64, encrypt bool) (*pb.IBTP, error) {
	result := &pb.Result{Data: data}
	content, err := result.Marshal()
	if err != nil {
		return nil, err
	}

	var packed []byte
	for _, ele := range data {
		packed = append(packed, ele...)
	}

	payload := pb.Payload{
		Encrypted: encrypt,
		Content:   content,
		Hash:      crypto.Keccak256(packed),
	}

	pd, err := payload.Marshal()
	if err != nil {
		return nil, err
	}

	return &pb.IBTP{
		From:          from,
		To:            to,
		Index:         idx,
		Type:          pb.IBTP_Type(typ),
		TimeoutHeight: 0,
		Proof:         []byte("1"),
		Payload:       pd,
	}, nil
}
