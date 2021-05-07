package main

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestM(t *testing.T) {
	etherCli, err := ethclient.Dial(rinkebyURL)
	require.Nil(t, err)

	header, err := etherCli.HeaderByNumber(context.Background(), big.NewInt(100))
	require.Nil(t, err)
	fmt.Printf("header 100 is %v\n", header)
}
