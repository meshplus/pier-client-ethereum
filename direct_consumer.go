package main

import "fmt"

func (c *Client) StartDirectConsumer() error {
	loop := func(interchainCh chan *BrokerDirectThrowInterchainEvent, receiptCh chan *BrokerDirectThrowReceiptEvent) {
		for {
			select {
			case interchainEv := <-interchainCh:
				ibtp, err := c.Convert2DirectIBTP(interchainEv, int64(c.config.Ether.TimeoutHeight))
				if err != nil {
					logger.Warn("convert to IBTP", "src", interchainEv.SrcFullID, "dst", interchainEv.DstFullID, "idx", interchainEv.Index, "err", err.Error())
					continue
				}
				c.eventC <- ibtp
			case receiptEv := <-receiptCh:
				ibtp, err := c.Convert2DirectReceipt(receiptEv)
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

	interchainCh := make(chan *BrokerDirectThrowInterchainEvent, 1024)
	receiptCh := make(chan *BrokerDirectThrowReceiptEvent, 1024)
	_, err := c.sessionDirect.Contract.WatchThrowInterchainEvent(nil, interchainCh)
	if err != nil {
		return fmt.Errorf("watch event: %s", err)
	}
	_, err = c.sessionDirect.Contract.WatchThrowReceiptEvent(nil, receiptCh)
	if err != nil {
		return fmt.Errorf("watch event: %s", err)
	}
	go loop(interchainCh, receiptCh)

	logger.Info("Consumer started")
	return nil
}
