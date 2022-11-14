package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	router *gin.Engine
	port   string
	client *Client

	ctx    context.Context
	cancel context.CancelFunc
}

func NewServer(port string, client *Client) (*Server, error) {
	ctx, cancel := context.WithCancel(context.Background())
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	return &Server{
		router: router,
		port:   port,
		client: client,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (s *Server) Start() error {
	s.router.Use(gin.Recovery())
	v1 := s.router.Group("/v1")
	{
		v1.POST("send", s.invoke)
	}

	go func() {
		go func() {
			err := s.router.Run(fmt.Sprintf(":%s", s.port))
			if err != nil {
				panic(err)
			}
		}()
		<-s.ctx.Done()
	}()

	return nil
}

type MockReq struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func (s *Server) invoke(c *gin.Context) {
	req := &MockReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	servicePair := fmt.Sprintf("%s-%s", req.From, req.To)
	s.client.lock.Lock()
	defer s.client.lock.Unlock()
	s.client.interchainInfo.outCounter[servicePair]++
	ev := &BrokerThrowInterchainEvent{
		Index:     s.client.interchainInfo.outCounter[servicePair],
		DstFullID: req.To,
		SrcFullID: req.From,
	}

	ibtp, err := s.client.Convert2IBTP(ev, int64(s.client.config.Mock.TimeoutHeight))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	ibtp.Timestamp = time.Now().UnixNano()
	s.client.eventC <- ibtp
	c.JSON(http.StatusOK, "send successfully!")
}
