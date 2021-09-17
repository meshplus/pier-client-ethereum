package main

import (
	"fmt"
)

func (c *Client) StartConsumer() error {
	loop := func(interchainCh chan *BrokerThrowInterchainEvent, receiptCh chan *BrokerThrowReceiptEvent) {
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

	interchainCh := make(chan *BrokerThrowInterchainEvent, 1024)
	receiptCh := make(chan *BrokerThrowReceiptEvent, 1024)
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
