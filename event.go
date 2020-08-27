package main

import (
	"fmt"
	"strconv"
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

	// TODO: generate proof for init and redeem
	var extra []byte
	if ev.Func == "interchainAssetExchangeInit" {
		ibtpType = pb.IBTP_ASSET_EXCHANGE_INIT
		extra, err = generateExtra(ev.Args, ibtpType)
		if err != nil {
			log.Fatalf("generate extra for asset exchange init :%s", err)
		}
	} else if ev.Func == "interchainAssetExchangeRedeem" {
		ibtpType = pb.IBTP_ASSET_EXCHANGE_REDEEM
		extra = []byte(ev.Args)
	} else if ev.Func == "interchainAssetExchangeRefund" {
		ibtpType = pb.IBTP_ASSET_EXCHANGE_REFUND
		extra = []byte(ev.Args)
	}

	return &pb.IBTP{
		From:      from,
		To:        ev.To.String(),
		Index:     ev.Index,
		Type:      ibtpType,
		Timestamp: time.Now().UnixNano(),
		Proof:     []byte("1"),
		Payload:   pd,
		Extra:     extra,
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
