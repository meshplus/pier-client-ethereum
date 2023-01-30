package main

import (
	"strings"

	"github.com/meshplus/bitxhub-model/pb"
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
	var results [][][]byte
	for _, s := range result.Data {
		results = append(results, s.Data)
	}
	if strings.EqualFold(content.Func, GET_FUNC) && len(result.Data) == 3 && strings.EqualFold(string(results[0][1]), "1") {
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
	return &pb.GetDataResponse{
		Index: req.Index,
		From:  req.From,
		To:    req.To,
	}
}
