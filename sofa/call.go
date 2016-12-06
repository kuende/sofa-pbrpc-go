package sofa

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

// Invoke sends the RPC request on the wire and returns after response is received.
// Invoke is called by generated code. Also users can call Invoke directly when it
// is really needed in their use cases.
func Invoke(ctx context.Context, method string, args, reply proto.Message, cc Conn, opts ...CallOption) error {
	if err := cc.RequestResponse(ctx, method, args, reply); err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	return nil
}
