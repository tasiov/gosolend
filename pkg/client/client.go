package client

import (
	"time"

	"github.com/tasiov/gosolend/internal/common"
	"github.com/tasiov/gosolend/pkg/service/obligations"
	"github.com/tasiov/gosolend/pkg/service/reserves"
)

type Client struct {
	endpoint    string
	timeout     time.Duration
	reserves    *reserves.Service
	obligations *obligations.Service
}

func NewClient(endpoint string, opts ...Option) *Client {
	c := &Client{
		endpoint: endpoint,
		timeout:  time.Duration(common.DefaultTimeout) * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.reserves = reserves.NewService(c.endpoint, c.timeout)
	c.obligations = obligations.NewService(c.endpoint, c.timeout)

	return c
}
