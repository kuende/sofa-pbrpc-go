package sofa

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/kuende/sofa-pbrpc-go/generated/protocol"
	"github.com/pkg/errors"
)

var (
	// ErrTimeoutExceeded is returned if no response is returned before timeout
	ErrTimeoutExceeded = errors.New("RPC timeout exceeded")

	magicString    = "SOFA"
	defaultTimeout = 3 * time.Second
)

// TCPTransport implements Sofa pbrpc over TCP
type TCPTransport struct {
	socket        net.Conn
	sequenceID    uint64
	sequenceIDMtx sync.Mutex

	Logger         Logger
	DefaultTimeout time.Duration

	inFlightRequests map[uint64]chan []byte
	inFlightMtx      sync.Mutex
}

// TCPTransportOptions defines optional configuration for TCPTransport
type TCPTransportOptions struct {
	Timeout *time.Duration
	Logger  Logger
}

// NewTCPTransport returns a new TCPTransport
func NewTCPTransport(addresses []string, options *TCPTransportOptions) (Transport, error) {
	// FIXME implement seed provider
	address := addresses[0]

	sock, err := net.Dial("tcp", address)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial tcp connection")
	}

	transport := &TCPTransport{
		socket:           sock,
		Logger:           NewLogger(),
		DefaultTimeout:   defaultTimeout,
		inFlightRequests: make(map[uint64]chan []byte),
	}

	transport.AddOptions(options)

	go transport.readLoop()

	return transport, nil
}

// AddOptions sets transport options
func (t *TCPTransport) AddOptions(options *TCPTransportOptions) {
	if options == nil {
		return
	}

	if options.Timeout != nil {
		t.DefaultTimeout = *options.Timeout
	}

	if options.Logger != nil {
		t.Logger = options.Logger
	}
}

// RequestResponse sends request and awaits response
func (t *TCPTransport) RequestResponse(ctx context.Context, method string, req, resp proto.Message) error {
	meta, err := t.makeMeta(method)

	if err != nil {
		return err
	}

	defer t.unregisterRequestID(meta)
	respChan, err := t.sendRequest(meta, req)

	if err != nil {
		return err
	}

	if err := t.awaitResponse(respChan, resp); err != nil {
		return err
	}

	return nil
}

func (t *TCPTransport) sendRequest(meta *protocol.RpcMeta, req proto.Message) (chan []byte, error) {
	request, err := t.makeRequest(meta, req)

	if err != nil {
		return nil, errors.Wrap(err, "failed to create request bytes")
	}

	ch, err := t.registerRequestID(meta)
	if err != nil {
		return nil, errors.Wrap(err, "failed to register request")
	}

	bytesCount, err := t.socket.Write(request)

	if err != nil {
		return nil, err
	}

	if bytesCount != len(request) {
		return nil, fmt.Errorf("Expected to send %d bytes, %d sent", len(request), bytesCount)
	}

	return ch, nil
}

func (t *TCPTransport) readLoop() {
	for {
		header := make([]byte, 24)

		count, err := t.socket.Read(header)

		if err != nil {
			t.Logger.Errorf("Failed to read from socket: %v", err)
			continue
		}

		if count != 24 {
			t.Logger.Errorf("Expected to read 24 bytes, got %d bytes", count)
			continue
		}

		if string(header[0:4]) != magicString {
			t.Logger.Errorf("Expected message to start with SOFA, started with %s", string(header[0:4]))
			continue
		}

		metaSize := binary.LittleEndian.Uint32(header[4:8])
		dataSize := binary.LittleEndian.Uint64(header[8:16])
		messageSize := binary.LittleEndian.Uint64(header[16:24])

		t.Logger.Debugf("Got: metaSize %d, dataSize %d, messageSize %d", metaSize, dataSize, messageSize)

		if messageSize != uint64(metaSize)+dataSize {
			t.Logger.Errorf("Expected messageSize to equal metaSize+dataSize, %d != %d+%d", messageSize, metaSize, dataSize)
			continue
		}

		metaBuf := make([]byte, metaSize)
		count, err = t.socket.Read(metaBuf)
		if err != nil {
			t.Logger.Errorf("Failed to read meta %v", err)
			continue
		}
		if uint32(count) != metaSize {
			t.Logger.Errorf("Expected to read %d bytes, got %d", metaSize, count)
			continue
		}

		dataBuf := make([]byte, dataSize)
		count, err = t.socket.Read(dataBuf)
		if err != nil {
			t.Logger.Errorf("Failed to read data %v", err)
			continue
		}
		if uint64(count) != dataSize {
			t.Logger.Errorf("Expected to read %d bytes, got %d", dataSize, count)
			continue
		}

		go t.processResponse(metaBuf, dataBuf)
	}
}

func (t *TCPTransport) processResponse(metaBytes, dataBytes []byte) {
	meta := &protocol.RpcMeta{}
	if err := proto.Unmarshal(metaBytes, meta); err != nil {
		t.Logger.Errorf("Failed to read meta: %v", err)
		return
	}

	ch, err := t.getResponseChan(meta)
	if err != nil {
		t.Logger.Errorf("Failed to get response chan")
		return
	}

	ch <- dataBytes
}

func (t *TCPTransport) registerRequestID(meta *protocol.RpcMeta) (chan []byte, error) {
	sequenceID := meta.SequenceId
	if sequenceID == nil {
		return nil, errors.New("sequenceID cannot be nil")
	}

	ch := make(chan []byte)

	t.inFlightMtx.Lock()
	defer t.inFlightMtx.Unlock()

	_, found := t.inFlightRequests[*sequenceID]
	if found {
		return nil, errors.New("sequenceID already present")
	}

	t.inFlightRequests[*sequenceID] = ch

	return ch, nil
}

func (t *TCPTransport) getResponseChan(meta *protocol.RpcMeta) (chan []byte, error) {
	sequenceID := meta.SequenceId
	if sequenceID == nil {
		return nil, errors.New("sequenceID cannot be nil")
	}

	t.inFlightMtx.Lock()
	defer t.inFlightMtx.Unlock()

	ch, found := t.inFlightRequests[*sequenceID]
	if !found {
		return nil, fmt.Errorf("failed to get in flight request for sequence id %d", *sequenceID)
	}

	return ch, nil
}

func (t *TCPTransport) unregisterRequestID(meta *protocol.RpcMeta) {
	sequenceID := meta.SequenceId
	if sequenceID == nil {
		t.Logger.Errorf("sequenceID cannot be nil")
		return
	}

	t.inFlightMtx.Lock()
	defer t.inFlightMtx.Unlock()

	_, found := t.inFlightRequests[*sequenceID]
	if !found {
		t.Logger.Errorf("sequendeID %d not found", *sequenceID)
		return
	}

	delete(t.inFlightRequests, *sequenceID)
}

func (t *TCPTransport) awaitResponse(ch chan []byte, resp proto.Message) error {
	for {
		select {
		case b := <-ch:
			if err := proto.Unmarshal(b, resp); err != nil {
				return err
			}

			return nil
		case <-time.After(t.DefaultTimeout):
			return ErrTimeoutExceeded
		}
	}
}

func (t *TCPTransport) makeRequest(meta *protocol.RpcMeta, req proto.Message) ([]byte, error) {
	metaBytestring, err := proto.Marshal(meta)
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

func (t *TCPTransport) makeMeta(method string) (*protocol.RpcMeta, error) {
	meta := &protocol.RpcMeta{
		Type:                         rpcMetaTypeP(protocol.RpcMeta_REQUEST),
		SequenceId:                   proto.Uint64(t.NextSequenceID()),
		Method:                       proto.String(method),
		ServerTimeout:                proto.Int64(3000), // FIXME, make it configurable
		CompressType:                 compressTypeP(protocol.CompressType_CompressTypeNone),
		ExpectedResponseCompressType: compressTypeP(protocol.CompressType_CompressTypeNone),
	}
	return meta, nil
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
