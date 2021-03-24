package main

import (
	"fmt"
	"time"

	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) generateCallback(original *pb.IBTP, args [][]byte, status bool) (result *pb.IBTP, err error) {
	if original == nil {
		return nil, fmt.Errorf("got nil ibtp To generate receipt: %w", err)
	}
	pd := &pb.Payload{}
	if err := pd.Unmarshal(original.Payload); err != nil {
		return nil, fmt.Errorf("ibtp payload unmarshal: %w", err)
	}

	originalContent := &pb.Content{}
	if err := originalContent.Unmarshal(pd.Content); err != nil {
		return nil, fmt.Errorf("ibtp payload unmarshal: %w", err)
	}

	content := &pb.Content{
		SrcContractId: originalContent.DstContractId,
		DstContractId: originalContent.SrcContractId,
	}

	if status {
		content.Func = originalContent.Callback
		content.Args = append(originalContent.ArgsCb, args...)
	} else {
		content.Func = originalContent.Rollback
		content.Args = originalContent.ArgsRb
	}

	b, err := content.Marshal()
	if err != nil {
		return nil, err
	}
	retPd := &pb.Payload{
		Content: b,
	}

	pdb, err := retPd.Marshal()
	if err != nil {
		return nil, err
	}

	typ := pb.IBTP_RECEIPT_SUCCESS
	if original.Type == pb.IBTP_INTERCHAIN {
		if !status {
			typ = pb.IBTP_RECEIPT_FAILURE
		}
	} else {
		typ = pb.IBTP_ASSET_EXCHANGE_RECEIPT
	}

	return &pb.IBTP{
		From:      original.From,
		To:        original.To,
		Index:     original.Index,
		Type:      typ,
		Timestamp: time.Now().UnixNano(),
		Proof:     original.Proof,
		Payload:   pdb,
		Version:   original.Version,
	}, nil
}
