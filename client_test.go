package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestM(t *testing.T) {
	//etherCli, err := ethclient.Dial(rinkebyURL)
	//require.Nil(t, err)
	//
	//header, err := etherCli.HeaderByNumber(context.Background(), big.NewInt(100))
	//require.Nil(t, err)
	//fmt.Printf("header 100 is %v\n", header)

	bytes := common.RightPadBytes([]byte("PIER_ROLE"), 32)
	fmt.Println(common.Bytes2Hex(bytes))
}
