package sofa

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

// Conn defines a sofa pbrpc client
type Conn interface {
	RequestResponse(context.Context, string, proto.Message, proto.Message) error
}

// ClientConn implements a sofa pbrpc client
type ClientConn struct {
	transport Transport
}

// NewClient returns a new sofa-pbrpc client
func NewClient(addresses []string) (Conn, error) {
	transport, err := NewTCPTransport(addresses)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create TCP transport")
	}

	client := &ClientConn{
		transport: transport,
	}

	return client, nil
}

// RequestResponse sends request and awaits response
func (c *ClientConn) RequestResponse(ctx context.Context, method string, req, resp proto.Message) error {
	return c.transport.RequestResponse(ctx, method, req, resp)
}
