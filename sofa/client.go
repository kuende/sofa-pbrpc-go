package sofa

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/kuende/sofa-pbrpc-go/sofa/discovery"
	"github.com/pkg/errors"
)

// Conn defines a sofa pbrpc client
type Conn interface {
	RequestResponse(context.Context, string, proto.Message, proto.Message) error
}

// ClientConn implements a sofa pbrpc client
type ClientConn struct {
	transport Transport
	provider  discovery.Provider
}

// ClientOpts defines a list of options for a client
type ClientOpts struct {
	TransportType     string
	DiscoveryProvider discovery.Provider
}

// NewClient returns a new sofa-pbrpc client
func NewClient(clientOpts ...ClientOpts) (Conn, error) {
	var opts ClientOpts
	if len(clientOpts) > 1 {
		return nil, errors.New("NewClient only accepts one option struct")
	} else if len(clientOpts) == 1 {
		opts = clientOpts[0]
	} else {
		opts = ClientOpts{
			TransportType: TCPTransportType,
		}
	}

	if opts.DiscoveryProvider == nil {
		return nil, errors.New("Please send opts.DiscoveryProvider")
	}

	hosts, err := opts.DiscoveryProvider.Hosts()

	if err != nil {
		return nil, errors.Wrap(err, "failed to get initial hosts")
	}

	transport, err := newClientTransport(hosts, opts)

	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize transport")
	}

	client := &ClientConn{
		transport: transport,
		provider:  opts.DiscoveryProvider,
	}

	return client, nil
}

// RequestResponse sends request and awaits response
func (c *ClientConn) RequestResponse(ctx context.Context, method string, req, resp proto.Message) error {
	return c.transport.RequestResponse(ctx, method, req, resp)
}

func newClientTransport(hosts []string, opts ClientOpts) (Transport, error) {
	transportType := opts.TransportType

	if len(hosts) == 0 {
		return nil, errors.New("hosts list cannot be empty")
	}

	switch transportType {
	case TCPTransportType:
		// FIXME: send tcp transport options from client opts
		return NewTCPTransport(hosts, &TCPTransportOptions{})
	case HTTPTransportType:
		return NewHTTPTransport(hosts[0])
	default:
		return nil, fmt.Errorf("undefined transport %s", transportType)
	}
}
