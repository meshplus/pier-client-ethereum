package main

import (
	"fmt"

	"github.com/meshplus/bitxhub-model/pb"
)

func (c *Client) StartConsumer() error {
	loop := func(ch chan *BrokerThrowEvent) {
		for {
			select {
			case ev := <-ch:
				c.eventC <- Convert2IBTP(ev, c.selfMethod, pb.IBTP_INTERCHAIN)
			case <-c.ctx.Done():
				return
			}
		}
	}

	evCh := make(chan *BrokerThrowEvent, 1024)
	_, err := c.session.Contract.WatchThrowEvent(nil, evCh)
	if err != nil {
		return fmt.Errorf("watch event: %s", err)
	}
	go loop(evCh)

	logger.Info("Consumer started")
	return nil
}
