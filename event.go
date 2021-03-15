package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bitxhub/bitxid"
	"github.com/cloudflare/cfssl/log"
	"github.com/meshplus/bitxhub-model/pb"
)

func Convert2IBTP(ev *BrokerThrowEvent, srcMethod string, ibtpType pb.IBTP_Type) *pb.IBTP {
	pd, err := encryptPayload(ev)
	if err != nil {
		log.Fatalf("Get ibtp payload :%s", err)
	}

	return &pb.IBTP{
		From:      srcMethod,
		To:        bitxid.DID(ev.DestDID).GetMethod(),
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
		ArgsCb:        handleArgs(ev.Args),
		Rollback:      funcs[2],
		ArgsRb:        handleArgs(ev.Args),
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

func generateExtra(args string, typ pb.IBTP_Type) ([]byte, error) {
	as := strings.Split(args, ",")

	if typ == pb.IBTP_ASSET_EXCHANGE_INIT {
		if len(as) != 8 {
			return nil, fmt.Errorf("incorrect args count for asset exchange init")
		}

		assetOnSrc, err := strconv.ParseUint(as[4], 10, 64)
		if err != nil {
			return nil, err
		}

		assetOnDst, err := strconv.ParseUint(as[7], 10, 64)
		if err != nil {
			return nil, err
		}

		aei := &pb.AssetExchangeInfo{
			Id:            as[1],
			SenderOnSrc:   as[2],
			ReceiverOnSrc: as[3],
			AssetOnSrc:    assetOnSrc,
			SenderOnDst:   as[4],
			ReceiverOnDst: as[5],
			AssetOnDst:    assetOnDst,
		}

		return aei.Marshal()
	}

	return nil, nil
}
