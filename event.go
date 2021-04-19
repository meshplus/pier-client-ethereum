package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/cloudflare/cfssl/log"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/meshplus/bitxid"
)

func Convert2IBTP(ev *BrokerThrowEvent, srcMethod string, ibtpType pb.IBTP_Type) *pb.IBTP {
	pd, err := encryptPayload(ev)
	if err != nil {
		log.Fatalf("Get ibtp payload :%s", err)
	}

	return &pb.IBTP{
		From:      srcMethod,
		To:        string(bitxid.DID(ev.DestDID).GetChainDID()),
		Index:     ev.Index,
		Type:      ibtpType,
		Timestamp: time.Now().UnixNano(),
		Proof:     []byte("1"),
		Payload:   pd,
	}
}

func handleArgs(args string) [][]byte {
	argsBytes := make([][]byte, 0)
	as := strings.Split(args, ",")
	for _, a := range as {
		argsBytes = append(argsBytes, []byte(a))
	}
	return argsBytes
}

func encryptPayload(ev *BrokerThrowEvent) ([]byte, error) {
	funcs := strings.Split(ev.Funcs, ",")
	if len(funcs) != 3 {
		return nil, fmt.Errorf("expected 3 functions, cur: %s", ev.Funcs)
	}

	content := &pb.Content{
		SrcContractId: ev.Fid.String(),
		DstContractId: bitxid.DID(ev.DestDID).GetAddress(),
		Func:          funcs[0],
		Args:          handleArgs(ev.Args),
		Callback:      funcs[1],
		ArgsCb:        handleArgs(ev.Argscb),
		Rollback:      funcs[2],
		ArgsRb:        handleArgs(ev.Argsrb),
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
