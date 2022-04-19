package main

import (
	"github.com/meshplus/bitxhub-model/pb"
	"strings"
)

const (
	PUSH_FUNC = "interchainPush"
	GET_FUNC  = "interchainGet"
)

func CheckInterchainOffChain(content *pb.Content) bool {
	if strings.EqualFold(content.Func, PUSH_FUNC) && len(content.Args) == 4 && strings.EqualFold(string(content.Args[2]), "1") {
		return true
	}
	return false
}

func CheckReceiptOffChain(ibtp *pb.IBTP, result *pb.Result) (bool, error) {
	pd := pb.Payload{}
	if err := pd.Unmarshal(ibtp.Payload); err != nil {
		return false, err
	}

	content := pb.Content{}
	if err := content.Unmarshal(pd.Content); err != nil {
		return false, err
	}

	if strings.EqualFold(content.Func, GET_FUNC) && len(result.Data) == 3 && strings.EqualFold(string(result.Data[1]), "1") {
		return true, nil
	}

	return false, nil
}

func constructReq(index uint64, from, to string, url []byte) *pb.GetDataRequest {
	return &pb.GetDataRequest{
		Index: index,
		From:  from,
		To:    to,
		Req:   url,
	}
}

func constructResp(req *pb.GetDataRequest) *pb.GetDataResponse {
	if req.IsSrc {
		return &pb.GetDataResponse{
			Index: req.Index,
			From:  req.To,
			To:    req.From,
		}
	}
	return &pb.GetDataResponse{
		Index: req.Index,
		From:  req.From,
		To:    req.To,
	}
}
