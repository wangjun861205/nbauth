package client

import (
	"github.com/wangjun861205/nbauth/api"
	"google.golang.org/grpc"
)

type Client struct {
	api.AuthServiceClient
	conn *grpc.ClientConn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func NewClient(addr string) (*Client, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{api.NewAuthServiceClient(conn), conn}, nil
}
