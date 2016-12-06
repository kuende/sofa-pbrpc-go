package sofa

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/kuende/sofa-pbrpc-go/generated/protocol"
	"github.com/pkg/errors"
)

// TCPTransport implements Sofa pbrpc over TCP
type TCPTransport struct {
	socket        net.Conn
	sequenceID    uint64
	sequenceIDMtx sync.Mutex
}

// NewTCPTransport returns a new TCPTransport
func NewTCPTransport(addresses []string) (Transport, error) {
	// FIXME implement seed provider
	address := addresses[0]

	sock, err := net.Dial("tcp", address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial tcp connection")
	}

	transport := &TCPTransport{
		socket: sock,
	}

	return transport, nil
}

// RequestResponse sends request and awaits response
func (t *TCPTransport) RequestResponse(ctx context.Context, method string, req, resp proto.Message) error {
	request, err := t.makeRequest(method, req)

	if err != nil {
		return err
	}

	bytesCount, err := t.socket.Write(request)

	if err != nil {
		return err
	}

	if bytesCount != len(request) {
		return fmt.Errorf("Expected to send %d bytes, %d sent", len(request), bytesCount)
	}

	return nil
}

func (t *TCPTransport) makeRequest(method string, req proto.Message) ([]byte, error) {
	metaBytestring, err := t.makeMeta(method)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode meta")
	}

	messageBytestring, err := proto.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode message")
	}

	buf := bytes.NewBuffer([]byte("SOFA"))

	if err := binary.Write(buf, binary.LittleEndian, uint32(len(metaBytestring))); err != nil {
		return nil, errors.Wrap(err, "failed to write meta length")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint64(len(messageBytestring))); err != nil {
		return nil, errors.Wrap(err, "failed to write message length")
	}
	if err := binary.Write(buf, binary.LittleEndian, uint64(len(metaBytestring)+len(messageBytestring))); err != nil {
		return nil, errors.Wrap(err, "failed to write meta + message length")
	}

	if _, err := buf.Write(metaBytestring); err != nil {
		return nil, errors.Wrap(err, "failed to write meta")
	}
	if _, err := buf.Write(messageBytestring); err != nil {
		return nil, errors.Wrap(err, "failed to write message")
	}

	return buf.Bytes(), nil
}

func (t *TCPTransport) makeMeta(method string) ([]byte, error) {
	meta := &protocol.RpcMeta{
		Type:                         rpcMetaTypeP(protocol.RpcMeta_REQUEST),
		SequenceId:                   proto.Uint64(t.NextSequenceID()),
		Method:                       proto.String(method),
		ServerTimeout:                proto.Int64(3000), // FIXME, make it configurable
		CompressType:                 compressTypeP(protocol.CompressType_CompressTypeNone),
		ExpectedResponseCompressType: compressTypeP(protocol.CompressType_CompressTypeNone),
	}
	return proto.Marshal(meta)
}

// NextSequenceID returns next sequence ID defined by the protocol
func (t *TCPTransport) NextSequenceID() uint64 {
	t.sequenceIDMtx.Lock()
	defer t.sequenceIDMtx.Unlock()
	t.sequenceID++
	return t.sequenceID
}

func rpcMetaTypeP(x protocol.RpcMeta_Type) *protocol.RpcMeta_Type {
	return &x
}

func compressTypeP(x protocol.CompressType) *protocol.CompressType {
	return &x
}
