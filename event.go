package main

import (
	"strings"
	"time"

	"github.com/cloudflare/cfssl/log"
	"github.com/meshplus/bitxhub-model/pb"
)

func Convert2IBTP(ev *BrokerThrowEvent, from string, ibtpType pb.IBTP_Type) *pb.IBTP {
	pd, err := encryptPayload(ev)
	if err != nil {
		log.Fatalf("Get ibtp payload :%s", err)
	}
	return &pb.IBTP{
		From:      from,
		To:        strings.ToLower(ev.To.String()),
		Index:     ev.Index,
		Type:      ibtpType,
		Timestamp: time.Now().UnixNano(),
		Proof:     []byte("1"),
		Payload:   pd,
	}
}

func encryptPayload(ev *BrokerThrowEvent) ([]byte, error) {
	args := make([][]byte, 0)
	as := strings.Split(ev.Args, ",")
	for _, a := range as {
		args = append(args, []byte(a))
	}
	content := &pb.Content{
		SrcContractId: ev.Fid.String(),
		DstContractId: ev.Tid,
		Func:          ev.Func,
		Args:          args,
		Callback:      ev.Callback,
	}
	data, err := content.Marshal()
	if err != nil {
		return nil, err
	}

	ibtppd := &pb.Payload{
		Content: data,
	}
	return ibtppd.Marshal()
}
