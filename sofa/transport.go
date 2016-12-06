package sofa

import (
	"context"

	"github.com/gogo/protobuf/proto"
)

// Transport defines client-server procotocol abstraction
type Transport interface {
	RequestResponse(context.Context, string, proto.Message, proto.Message) error
}
