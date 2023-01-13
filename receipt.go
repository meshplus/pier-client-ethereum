package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) generateCallback(original *pb.IBTP, data [][][]byte, typ uint64, multiStatus []bool) (result *pb.IBTP, err error) {
	payload := &pb.Payload{}
	if err := payload.Unmarshal(original.Payload); err != nil {
		return nil, err
	}

	return generateReceipt(original.From, original.To, original.Index, data, typ, payload.Encrypted, multiStatus)
}

func generateReceipt(from, to string, idx uint64, data [][][]byte, typ uint64, encrypt bool, multiStatus []bool) (*pb.IBTP, error) {
	//result := &pb.Result{Data: data}
	var result []*pb.ResultRes
	for _, s := range data {
		res := &pb.ResultRes{Data: s}
		result = append(result, res)
	}
	results := &pb.Result{Data: result, MultiStatus: multiStatus}
	content, err := results.Marshal()
	if err != nil {
		return nil, err
	}

	var packed []byte
	for _, ele := range data {
		for _, val := range ele {
			packed = append(packed, val...)
		}
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

// TODO
func unpackToBytesArray(data []byte) ([][]byte, error) {
	return nil, nil
}
