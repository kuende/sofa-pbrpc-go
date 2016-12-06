package main

import (
	"context"

	"github.com/kuende/sofa-pbrpc-go/sofa"
)

// Client API for EchoServer service

// EchoServerClient implements RPC client for EchoServer
type EchoServerClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...sofa.CallOption) (*EchoResponse, error)
}

type echoServerClient struct {
	cc sofa.Conn
}

// NewEchoServerClient returns a new EchoServer client
func NewEchoServerClient(cc sofa.Conn) EchoServerClient {
	return &echoServerClient{cc}
}

func (c *echoServerClient) Echo(ctx context.Context, in *EchoRequest, opts ...sofa.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := sofa.Invoke(ctx, "sofa.pbrpc.test.EchoServer.Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
