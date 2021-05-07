package main

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/light"
)

const (
	defaultCap = 20
)

type headerPool struct {
	batchCh      chan []*types.Header
	recvHeaderCh chan *types.Header

	headersSet []*types.Header
	currentNum uint64
}

func newHeaderBuffer() *headerPool {
	return &headerPool{
		headersSet:   make([]*types.Header, 0, defaultCap),
		batchCh:      make(chan []*types.Header, defaultCap),
		recvHeaderCh: make(chan *types.Header, defaultCap),
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
				c.headerPool.headersSet = make([]*types.Header, 0, defaultCap)
				c.metaC <- batch
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

	pool := newHeaderBuffer()
	c.headerPool = pool

	for {
		select {
		case <-ticker.C:
			// get latest blockchain height and got all near-finalized headers into pool
			latestHeight, err := c.ethClient.BlockNumber(c.ctx)
			if err != nil {
				logger.Error("get most recent height", "error", err.Error())
				continue
			}
			for i := c.headerPool.currentNum + 1; i <= latestHeight-Threshold; i++ {
				header, err := light.GetHeaderByNumber(c.ctx, nil, c.headerPool.currentNum)
				if err != nil {
					return
				}
				pool.recvHeaderCh <- header
				c.headerPool.currentNum++
			}
		case <-c.ctx.Done():
			ticker.Stop()
			return
		}
	}
}
