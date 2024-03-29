// Code generated by protoc-gen-gogo.
// source: rpc_meta.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	rpc_meta.proto
	rpc_option.proto

It has these top-level messages:
	RpcMeta
*/
package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Message type.
type RpcMeta_Type int32

const (
	RpcMeta_REQUEST  RpcMeta_Type = 0
	RpcMeta_RESPONSE RpcMeta_Type = 1
)

var RpcMeta_Type_name = map[int32]string{
	0: "REQUEST",
	1: "RESPONSE",
}
var RpcMeta_Type_value = map[string]int32{
	"REQUEST":  0,
	"RESPONSE": 1,
}

func (x RpcMeta_Type) Enum() *RpcMeta_Type {
	p := new(RpcMeta_Type)
	*p = x
	return p
}
func (x RpcMeta_Type) String() string {
	return proto.EnumName(RpcMeta_Type_name, int32(x))
}
func (x *RpcMeta_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpcMeta_Type_value, data, "RpcMeta_Type")
	if err != nil {
		return err
	}
	*x = RpcMeta_Type(value)
	return nil
}
func (RpcMeta_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorRpcMeta, []int{0, 0} }

type RpcMeta struct {
	Type *RpcMeta_Type `protobuf:"varint,1,req,name=type,enum=protocol.RpcMeta_Type" json:"type,omitempty"`
	// Message sequence id.
	SequenceId *uint64 `protobuf:"varint,2,req,name=sequence_id,json=sequenceId" json:"sequence_id,omitempty"`
	// Method full name.
	// For example: "test.HelloService.GreetMethod"
	Method *string `protobuf:"bytes,100,opt,name=method" json:"method,omitempty"`
	// Server timeout in milli-seconds.
	ServerTimeout *int64 `protobuf:"varint,101,opt,name=server_timeout,json=serverTimeout" json:"server_timeout,omitempty"`
	// Set as true if the call is failed.
	Failed *bool `protobuf:"varint,200,opt,name=failed" json:"failed,omitempty"`
	// The error code if the call is failed.
	ErrorCode *int32 `protobuf:"varint,201,opt,name=error_code,json=errorCode" json:"error_code,omitempty"`
	// The error reason if the call is failed.
	Reason *string `protobuf:"bytes,202,opt,name=reason" json:"reason,omitempty"`
	// Set the request/response compress type.
	CompressType *CompressType `protobuf:"varint,300,opt,name=compress_type,json=compressType,enum=protocol.CompressType" json:"compress_type,omitempty"`
	// Set the response compress type of user expected.
	ExpectedResponseCompressType *CompressType `protobuf:"varint,301,opt,name=expected_response_compress_type,json=expectedResponseCompressType,enum=protocol.CompressType" json:"expected_response_compress_type,omitempty"`
	XXX_unrecognized             []byte        `json:"-"`
}

func (m *RpcMeta) Reset()                    { *m = RpcMeta{} }
func (m *RpcMeta) String() string            { return proto.CompactTextString(m) }
func (*RpcMeta) ProtoMessage()               {}
func (*RpcMeta) Descriptor() ([]byte, []int) { return fileDescriptorRpcMeta, []int{0} }

func (m *RpcMeta) GetType() RpcMeta_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return RpcMeta_REQUEST
}

func (m *RpcMeta) GetSequenceId() uint64 {
	if m != nil && m.SequenceId != nil {
		return *m.SequenceId
	}
	return 0
}

func (m *RpcMeta) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *RpcMeta) GetServerTimeout() int64 {
	if m != nil && m.ServerTimeout != nil {
		return *m.ServerTimeout
	}
	return 0
}

func (m *RpcMeta) GetFailed() bool {
	if m != nil && m.Failed != nil {
		return *m.Failed
	}
	return false
}

func (m *RpcMeta) GetErrorCode() int32 {
	if m != nil && m.ErrorCode != nil {
		return *m.ErrorCode
	}
	return 0
}

func (m *RpcMeta) GetReason() string {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return ""
}

func (m *RpcMeta) GetCompressType() CompressType {
	if m != nil && m.CompressType != nil {
		return *m.CompressType
	}
	return CompressType_CompressTypeNone
}

func (m *RpcMeta) GetExpectedResponseCompressType() CompressType {
	if m != nil && m.ExpectedResponseCompressType != nil {
		return *m.ExpectedResponseCompressType
	}
	return CompressType_CompressTypeNone
}

func init() {
	proto.RegisterType((*RpcMeta)(nil), "protocol.RpcMeta")
	proto.RegisterEnum("protocol.RpcMeta_Type", RpcMeta_Type_name, RpcMeta_Type_value)
}

func init() { proto.RegisterFile("rpc_meta.proto", fileDescriptorRpcMeta) }

var fileDescriptorRpcMeta = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x8f, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x71, 0x5a, 0xfa, 0x73, 0xfb, 0xa3, 0xca, 0x43, 0xb1, 0x10, 0xa2, 0xa1, 0x12, 0x52,
	0xc4, 0x90, 0x81, 0x99, 0xad, 0xca, 0xc0, 0xc0, 0x9f, 0x5b, 0x56, 0xac, 0xc8, 0xbe, 0x88, 0x48,
	0x4d, 0x6c, 0x6c, 0x17, 0xd1, 0x37, 0x62, 0x80, 0xf7, 0x00, 0x9e, 0x0a, 0x35, 0x49, 0xa5, 0x76,
	0x61, 0xba, 0x3a, 0xe7, 0xbb, 0xe7, 0x5e, 0x1d, 0x18, 0x5a, 0x23, 0x45, 0x8e, 0x3e, 0x8d, 0x8d,
	0xd5, 0x5e, 0xd3, 0x4e, 0x39, 0xa4, 0x5e, 0x1e, 0x8f, 0x36, 0x44, 0x1b, 0x9f, 0xe9, 0xa2, 0x62,
	0xd3, 0x8f, 0x06, 0xb4, 0xb9, 0x91, 0x37, 0xe8, 0x53, 0x7a, 0x01, 0x4d, 0xbf, 0x36, 0xc8, 0x48,
	0x18, 0x44, 0xc3, 0xcb, 0x71, 0xbc, 0x8d, 0xc5, 0xf5, 0x42, 0xbc, 0x58, 0x1b, 0xe4, 0xe5, 0x0e,
	0x9d, 0x40, 0xcf, 0xe1, 0xeb, 0x0a, 0x0b, 0x89, 0x22, 0x53, 0x2c, 0x08, 0x83, 0xa8, 0xc9, 0x61,
	0x6b, 0x5d, 0x2b, 0x3a, 0x86, 0x56, 0x8e, 0xfe, 0x45, 0x2b, 0xa6, 0x42, 0x12, 0x75, 0x79, 0xad,
	0xe8, 0x39, 0x0c, 0x1d, 0xda, 0x37, 0xb4, 0xc2, 0x67, 0x39, 0xea, 0x95, 0x67, 0x18, 0x92, 0xa8,
	0xc1, 0x07, 0x95, 0xbb, 0xa8, 0x4c, 0x7a, 0x04, 0xad, 0xe7, 0x34, 0x5b, 0xa2, 0x62, 0xdf, 0x24,
	0x24, 0x51, 0x87, 0xd7, 0x92, 0x9e, 0x02, 0xa0, 0xb5, 0xda, 0x0a, 0xa9, 0x15, 0xb2, 0x9f, 0x0d,
	0x3c, 0xe4, 0xdd, 0xd2, 0x9a, 0x69, 0x85, 0x9b, 0xa0, 0xc5, 0xd4, 0xe9, 0x82, 0xfd, 0x92, 0xea,
	0x71, 0x25, 0xe9, 0x15, 0x0c, 0xa4, 0xce, 0x8d, 0x45, 0xe7, 0x44, 0x59, 0xf3, 0x33, 0x08, 0xc9,
	0x7e, 0xcf, 0x59, 0xcd, 0xcb, 0x9e, 0x7d, 0xb9, 0xa3, 0xe8, 0x13, 0x4c, 0xf0, 0xdd, 0xa0, 0xf4,
	0xa8, 0x84, 0x45, 0x67, 0x74, 0xe1, 0x50, 0xec, 0xdf, 0xfb, 0xfa, 0xff, 0xde, 0xc9, 0x36, 0xcf,
	0xeb, 0xf8, 0x2e, 0x9d, 0x9e, 0x41, 0xb3, 0xfc, 0xd3, 0x83, 0x36, 0x4f, 0x1e, 0x1e, 0x93, 0xf9,
	0x62, 0x74, 0x40, 0xfb, 0xd0, 0xe1, 0xc9, 0xfc, 0xfe, 0xee, 0x76, 0x9e, 0x8c, 0xc8, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd9, 0x53, 0x90, 0xf2, 0xd7, 0x01, 0x00, 0x00,
}
