package main

import (
	"bytes"
	"math/big"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/light"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
	contracts "github.com/meshplus/bitxhub-core/eth-contracts"
	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) listenLock() {
	for {
		select {
		case log := <-c.logCh:
			// query this block from ethereum and generate mintEvent and proof for pier
			if err := retry.Retry(func(attempt uint) error {
				block, err := c.ethClient.BlockByNumber(c.ctx, big.NewInt(int64(log.BlockNumber)))
				if err != nil {
					return err
				}
				receipt, err := c.ethClient.TransactionReceipt(c.ctx, log.TxHash)
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
				proof, err := c.getProof(receiptsTrie, uint64(log.Index))
				if err != nil {
					return err
				}
				c.lockCh <- &pb.LockEvent{
					Receipt: receiptData,
					Proof:       proof,
				}
				return nil
			}, strategy.Wait(1*time.Second)); err != nil {
				logger.Error("Can't retrieve mint event from receipt", "error", err.Error())
			}
		case filterOpt := <-c.filterOptCh:
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
				c.logCh <- &iter.Event.Raw
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) getProof(receiptsTrie *trie.Trie, index uint64) ([]byte, error) {
	keybuf := new(bytes.Buffer)
	err := rlp.Encode(keybuf, index)
	if err != nil {
		return nil, err
	}
	nodeSet := light.NewNodeSet()
	err = receiptsTrie.Prove(keybuf.Bytes(), 0, nodeSet)
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(nodeSet.NodeList())
}

func (c *Client) filterLog(batch []*types.Header) {
	start := batch[0].Number
	end := batch[len(batch)-1].Number.Uint64()
	filterOpt := &bind.FilterOpts{
		Start: start.Uint64(),
		End:   &end,
	}
	c.filterOptCh <- filterOpt
}
