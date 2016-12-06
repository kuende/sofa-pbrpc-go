package sofa

import (
	"context"
	"errors"
	"time"

	"github.com/gogo/protobuf/proto"
)

var (
	// ErrTimeoutExceeded is returned if no response is returned before timeout
	ErrTimeoutExceeded = errors.New("RPC timeout exceeded")
	defaultTimeout     = 3 * time.Second
)

// Transport defines client-server procotocol abstraction
type Transport interface {
	RequestResponse(context.Context, string, proto.Message, proto.Message) error
}
