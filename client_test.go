package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"strings"
	"time"

	"testing"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	contracts "github.com/meshplus/bitxhub-core/eth-contracts/escrows-contracts"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/stretchr/testify/require"
)

func TestM(t *testing.T) {
	//etherCli, err := ethclient.Dial(rinkebyURL)
	//require.Nil(t, err)
	//
	//header, err := etherCli.HeaderByNumber(context.Background(), big.NewInt(100))
	//require.Nil(t, err)
	//fmt.Printf("header 100 is %v\n", header)
	etherCli, err := ethclient.Dial("http://121.41.217.124:8545")
	require.Nil(t, err)
	number, err := etherCli.BlockNumber(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	require.Nil(t, err)

	fmt.Println(number)
	block, _ := etherCli.BlockByNumber(context.Background(), big.NewInt(10540394))
	s, _ := json.Marshal(block.Header())
	fmt.Println(string(s))

	//bytes := common.RightPadBytes([]byte("PIER_ROLE"), 32)
	//fmt.Println(common.Bytes2Hex(bytes))
}

func TestGetPrivKey(t *testing.T) {
	configPath := "./config"
	cfg, _ := UnmarshalConfig(configPath)

	keyPath := filepath.Join(configPath, cfg.Ether.KeyPath)
	keyByte, _ := ioutil.ReadFile(keyPath)

	psdPath := filepath.Join(configPath, cfg.Ether.Password)
	password, _ := ioutil.ReadFile(psdPath)
	unlockedKey, _ := keystore.DecryptKey(keyByte, strings.TrimSpace(string(password)))
	fmt.Println(hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes()))

}

func TestUnpack(t *testing.T) {
	etherCli, _ := ethclient.Dial("wss://ropsten.infura.io/ws/v3/042b7404d74f4f18bbca771786fed781")
	escrowsContract, _ := contracts.NewEscrows(common.HexToAddress("0x956Be099e5Add3d95aaB9D1a7Da5a40eB9d02528"), etherCli)
	escrowsSession := &contracts.EscrowsSession{
		Contract: escrowsContract,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
	}
	c := &Client{ethClient: etherCli, escrowsSession: escrowsSession, ctx: context.Background()}
	var lockCh *pb.LockEvent
	height, er := c.escrowsSession.Index2Height(big.NewInt(1))
	if er != nil {
		return
	}
	end := height.Uint64()
	filterOpt := &bind.FilterOpts{
		Start: end,
		End:   &end,
	}

	var (
		iter *contracts.EscrowsLockIterator
		err  error
	)
	if err := retry.Retry(func(attempt uint) error {
		iter, err = c.escrowsSession.Contract.FilterLock(filterOpt)
		if err != nil {
			return err
		}
		return nil
	}, strategy.Wait(1*time.Second)); err != nil {
		logger.Error("Can't get filter mint event", "error", err.Error())
	}
	for iter.Next() {
		if 1 != iter.Event.AppchainIndex.Int64() {
			continue
		}
		raw := &iter.Event.Raw
		// query this block from ethereum and generate mintEvent and proof for pier
		if err := retry.Retry(func(attempt uint) error {
			block, err := c.ethClient.BlockByNumber(c.ctx, big.NewInt(int64(raw.BlockNumber)))
			if err != nil {
				return err
			}
			receipt, err := c.ethClient.TransactionReceipt(c.ctx, raw.TxHash)
			if err != nil {
				return err
			}
			// construct receipt merkle tree first for proof generate
			receipts := make([]*types.Receipt, 0, len(block.Transactions()))
			for _, tx := range block.Transactions() {
				receipt, err := c.ethClient.TransactionReceipt(c.ctx, tx.Hash())
				if err != nil {
					return err
				}
				receipts = append(receipts, receipt)
			}
			receiptsTrie := new(trie.Trie)
			tReceipts := types.Receipts(receipts)
			types.DeriveSha(tReceipts, receiptsTrie)

			receiptData, err := receipt.MarshalJSON()
			if err != nil {
				return err
			}
			proof, err := c.getProof(receiptsTrie, uint64(receipt.TransactionIndex))
			if err != nil {
				return err
			}
			lockCh = &pb.LockEvent{
				//AppchainIndex: event.AppchainIndex,
				//BlockNumber:   event.BlockNumber,
				Receipt: receiptData,
				Proof:   proof,
			}
			// 验证签名
			nodeList := &light.NodeList{}
			if err := rlp.DecodeBytes(proof, nodeList); err != nil {
				return err
			}
			keyBuf := bytes.Buffer{}
			keyBuf.Reset()
			if err := rlp.Encode(&keyBuf, receipt.TransactionIndex); err != nil {
				return err
			}
			value, err := trie.VerifyProof(block.Header().ReceiptHash, keyBuf.Bytes(), nodeList.NodeSet())
			if err != nil {
				return err
			}
			fmt.Println(value)

			return nil
		}, strategy.Wait(1*time.Second)); err != nil {
			logger.Error("Can't retrieve mint event from receipt", "error", err.Error())
		}
	}
	unpackEscrowsLock(lockCh.Receipt)

}

type ContractAddr struct {
	Addr string `json:"addr"`
}

func unpackEscrowsLock(receiptData []byte) (*contracts.EscrowsLock, error) {
	var receipt types.Receipt
	err := receipt.UnmarshalJSON(receiptData)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var lock *contracts.EscrowsLock

	for _, log := range receipt.Logs {

		if !strings.EqualFold(log.Address.String(), "0x956Be099e5Add3d95aaB9D1a7Da5a40eB9d02528") {
			continue
		}

		if log.Removed {
			continue
		}

		escrows, err := contracts.NewEscrows(common.Address{}, nil)
		if err != nil {
			continue
		}
		lock, _ = escrows.ParseLock(*log)
	}
	if lock == nil {
		return nil, fmt.Errorf("not found the escrow lock event")
	}
	return lock, nil
}
