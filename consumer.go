//nolint:dupl
package main

import (
	"fmt"

	"github.com/meshplus/pier-client-ethereum/relay"
)

func (c *Client) StartConsumer() error {
	loop := func(interchainCh chan *relay.BrokerThrowInterchainEvent, receiptCh chan *relay.BrokerThrowReceiptEvent) {
		for {
			select {
			case interchainEv := <-interchainCh:
				ibtp, err := c.Convert2IBTP(interchainEv, int64(c.config.Ether.TimeoutHeight))
				if err != nil {
					logger.Warn("convert to IBTP", "src", interchainEv.SrcFullID, "dst", interchainEv.DstFullID, "idx", interchainEv.Index, "err", err.Error())
					continue
				}
				c.eventC <- ibtp
			case receiptEv := <-receiptCh:
				ibtp, err := c.Convert2Receipt(receiptEv)
				if err != nil {
					logger.Warn("convert to IBTP", "src", receiptEv.SrcFullID, "dst", receiptEv.DstFullID, "idx", receiptEv.Index, "err", err.Error())
					continue
				}
				c.eventC <- ibtp
			case <-c.ctx.Done():
				return
			}
		}
	}

	interchainCh := make(chan *relay.BrokerThrowInterchainEvent, 1024)
	receiptCh := make(chan *relay.BrokerThrowReceiptEvent, 1024)
	_, err := c.session.Contract.WatchThrowInterchainEvent(nil, interchainCh)
	if err != nil {
		return fmt.Errorf("watch event: %s", err)
	}
	_, err = c.session.Contract.WatchThrowReceiptEvent(nil, receiptCh)
	if err != nil {
		return fmt.Errorf("watch event: %s", err)
	}
	go loop(interchainCh, receiptCh)

	logger.Info("Consumer started")
	return nil
}
