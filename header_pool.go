package main

import (
	"math/big"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	defaultCap = 200
)

type headerPool struct {
	batchCh      chan []*types.Header
	recvHeaderCh chan *types.Header

	headersSet []*types.Header
	currentNum uint64
}

func newHeaderPool(currentNum uint64) *headerPool {
	return &headerPool{
		headersSet:   make([]*types.Header, 0, defaultCap),
		batchCh:      make(chan []*types.Header, defaultCap),
		recvHeaderCh: make(chan *types.Header, defaultCap),
		currentNum:   currentNum,
	}
}

func (b *headerPool) append(header *types.Header) {
	b.headersSet = append(b.headersSet, header)
}

// postHeaders listen on blockchain headersSet periodically and post headers if not empty
func (c *Client) postHeaders() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case header := <-c.headerPool.recvHeaderCh:
			c.headerPool.append(header)
		case <-ticker.C:
			// check if there are any headers in buffer;
			// if so, post a new batch of block headers; else return
			if len(c.headerPool.headersSet) != 0 {
				batch := c.headerPool.headersSet
				c.filterLog(batch)
				c.headerPool.headersSet = make([]*types.Header, 0, defaultCap)
				//data, _ := json.Marshal(batch)
				//c.metaC <- &pb.UpdateMeta{
				//	Meta:      data,
				//	EndHeader: batch[len(batch)-1].Number.Uint64(),
				//}
			}
		case <-c.ctx.Done():
			ticker.Stop()
			return
		}
	}
}

// listen on block headers in ethereum periodically
func (c *Client) listenHeader() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// get latest blockchain height and got all finalized headers into pool
			latestHeight, err := c.ethClient.BlockNumber(c.ctx)
			if err != nil {
				logger.Error("get most recent height", "error", err.Error())
				continue
			}
			for i := c.headerPool.currentNum; i <= latestHeight; i++ { //Threshold
				c.headerPool.currentNum++
				var header *types.Header
				if err := retry.Retry(func(attempt uint) error {
					var err error
					header, err = c.ethClient.HeaderByNumber(c.ctx, big.NewInt(int64(c.headerPool.currentNum)))
					if err != nil {
						return err
					}
					return nil
				}, strategy.Wait(1*time.Second)); err != nil {
					logger.Error("Can't receive head from ethclient", "error", err.Error())
				}

				c.headerPool.recvHeaderCh <- header
			}
		case <-c.ctx.Done():
			ticker.Stop()
			return
		}
	}
}
