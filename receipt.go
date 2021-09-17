package main

import (
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) generateCallback(original *pb.IBTP, data [][]byte, status bool) (result *pb.IBTP, err error) {
	payload := &pb.Payload{}
	if err := payload.Unmarshal(original.Payload); err != nil {
		return nil, err
	}

	return generateReceipt(original.From, original.To, original.Index, data, status, payload.Encrypted)
}

func generateReceipt(from, to string, idx uint64, data [][]byte, status, encrypt bool) (*pb.IBTP, error) {
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

	typ := pb.IBTP_RECEIPT_SUCCESS
	if !status {
		typ = pb.IBTP_RECEIPT_FAILURE
	}

	return &pb.IBTP{
		From:          from,
		To:            to,
		Index:         idx,
		Type:          typ,
		TimeoutHeight: 0,
		Proof:         []byte("1"),
		Payload:       pd,
	}, nil
}

// TODO
func unpackToBytesArray(data []byte) ([][]byte, error) {
	return nil, nil
}
