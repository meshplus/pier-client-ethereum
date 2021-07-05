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

type PreLockEvent struct {
	Receipt       []byte
	Proof         []byte
	AppchainIndex uint64
	BlockNumber   uint64
}

func (c *Client) postLock() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	loop := func(ch <-chan *PreLockEvent) {
		for {
			select {
			case event, ok := <-ch:
				if !ok {
					logger.Warn("preLockEvent handle error")
					return
				}
				if c.headerPool.currentNum-event.BlockNumber > 20 {
					c.lockCh <- &pb.LockEvent{
						AppchainIndex: event.AppchainIndex,
						Receipt:       event.Receipt,
						Proof:         event.Proof,
						BlockNumber:   event.BlockNumber,
					}
				} else {
					// not enough attach current height will redo in preLockCh
					// todo set a queue
					c.preLockCh <- event
				}
			case <-c.ctx.Done():
				return
			}
		}
	}
	for {
		select {
		case <-ticker.C:
			ch := c.preLockCh
			loop(ch)
		case <-c.ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func (c *Client) listenLock() {
	for {
		select {
		case log := <-c.logCh:
			// query this block from ethereum and generate mintEvent and proof for pier
			if err := retry.Retry(func(attempt uint) error {
				block, err := c.ethClient.BlockByNumber(c.ctx, big.NewInt(int64(log.Raw.BlockNumber)))
				if err != nil {
					return err
				}
				receipt, err := c.ethClient.TransactionReceipt(c.ctx, log.Raw.TxHash)
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
				c.preLockCh <- &PreLockEvent{
					AppchainIndex: log.AppchainIndex.Uint64(),
					Receipt:       receiptData,
					Proof:         proof,
					BlockNumber:   log.Raw.BlockNumber,
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
				c.logCh <- iter.Event
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
	logger.Info("current filter log batch info", "start", start.String())
	logger.Info("current filter log batch info", "end", end)
	filterOpt := &bind.FilterOpts{
		Start: start.Uint64(),
		End:   &end,
	}
	c.filterOptCh <- filterOpt
}
