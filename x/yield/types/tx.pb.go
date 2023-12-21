// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: side/yield/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DepositRecord_Status int32

const (
	// in transfer queue to be sent to the delegation ICA (Hub)
	DepositRecord_TRANSFER_QUEUE DepositRecord_Status = 0
	// transfer in progress (IBC packet sent, ack not received) (to Hub)
	DepositRecord_TRANSFER_FIRST_IN_PROGRESS DepositRecord_Status = 2
	// transfer in progress (IBC packet sent, ack not received) (from Hub to stride)
	DepositRecord_TRANSFER_SECOND_IN_PROGRESS DepositRecord_Status = 3
	// in staking queue on delegation ICA
	DepositRecord_DELEGATION_QUEUE DepositRecord_Status = 1
	// staking in progress (ICA packet sent, ack not received)
	DepositRecord_DELEGATION_IN_PROGRESS DepositRecord_Status = 4
)

var DepositRecord_Status_name = map[int32]string{
	0: "TRANSFER_QUEUE",
	2: "TRANSFER_FIRST_IN_PROGRESS",
	3: "TRANSFER_SECOND_IN_PROGRESS",
	1: "DELEGATION_QUEUE",
	4: "DELEGATION_IN_PROGRESS",
}

var DepositRecord_Status_value = map[string]int32{
	"TRANSFER_QUEUE":              0,
	"TRANSFER_FIRST_IN_PROGRESS":  2,
	"TRANSFER_SECOND_IN_PROGRESS": 3,
	"DELEGATION_QUEUE":            1,
	"DELEGATION_IN_PROGRESS":      4,
}

func (x DepositRecord_Status) String() string {
	return proto.EnumName(DepositRecord_Status_name, int32(x))
}

func (DepositRecord_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_219f732381d2fb9e, []int{2, 0}
}

type DepositRecord_Source int32

const (
	DepositRecord_SIDE DepositRecord_Source = 0
	DepositRecord_HUB  DepositRecord_Source = 1
)

var DepositRecord_Source_name = map[int32]string{
	0: "SIDE",
	1: "HUB",
}

var DepositRecord_Source_value = map[string]int32{
	"SIDE": 0,
	"HUB":  1,
}

func (x DepositRecord_Source) String() string {
	return proto.EnumName(DepositRecord_Source_name, int32(x))
}

func (DepositRecord_Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_219f732381d2fb9e, []int{2, 1}
}

type MsgRegisterHostChain struct {
	ConnectionId      string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	Bech32Prefix      string `protobuf:"bytes,2,opt,name=bech32prefix,proto3" json:"bech32prefix,omitempty"`
	HostDenom         string `protobuf:"bytes,3,opt,name=host_denom,json=hostDenom,proto3" json:"host_denom,omitempty" yaml:"host_denom"`
	IbcDenom          string `protobuf:"bytes,4,opt,name=ibc_denom,json=ibcDenom,proto3" json:"ibc_denom,omitempty" yaml:"ibc_denom"`
	Creator           string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	TransferChannelId string `protobuf:"bytes,6,opt,name=transfer_channel_id,json=transferChannelId,proto3" json:"transfer_channel_id,omitempty" yaml:"transfer_channel_id"`
}

func (m *MsgRegisterHostChain) Reset()         { *m = MsgRegisterHostChain{} }
func (m *MsgRegisterHostChain) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterHostChain) ProtoMessage()    {}
func (*MsgRegisterHostChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_219f732381d2fb9e, []int{0}
}
func (m *MsgRegisterHostChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterHostChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterHostChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterHostChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterHostChain.Merge(m, src)
}
func (m *MsgRegisterHostChain) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterHostChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterHostChain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterHostChain proto.InternalMessageInfo

type MsgRegisterHostChainResponse struct {
}

func (m *MsgRegisterHostChainResponse) Reset()         { *m = MsgRegisterHostChainResponse{} }
func (m *MsgRegisterHostChainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterHostChainResponse) ProtoMessage()    {}
func (*MsgRegisterHostChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_219f732381d2fb9e, []int{1}
}
func (m *MsgRegisterHostChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterHostChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterHostChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterHostChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterHostChainResponse.Merge(m, src)
}
func (m *MsgRegisterHostChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterHostChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterHostChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterHostChainResponse proto.InternalMessageInfo

type DepositRecord struct {
	Id                 uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount             github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	Denom              string                                 `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
	HostChainId        string                                 `protobuf:"bytes,4,opt,name=host_chain_id,json=hostChainId,proto3" json:"host_chain_id,omitempty"`
	Status             DepositRecord_Status                   `protobuf:"varint,6,opt,name=status,proto3,enum=side.yield.DepositRecord_Status" json:"status,omitempty"`
	DepositEpochNumber uint64                                 `protobuf:"varint,7,opt,name=deposit_epoch_number,json=depositEpochNumber,proto3" json:"deposit_epoch_number,omitempty"`
	Source             DepositRecord_Source                   `protobuf:"varint,8,opt,name=source,proto3,enum=side.yield.DepositRecord_Source" json:"source,omitempty"`
}

func (m *DepositRecord) Reset()         { *m = DepositRecord{} }
func (m *DepositRecord) String() string { return proto.CompactTextString(m) }
func (*DepositRecord) ProtoMessage()    {}
func (*DepositRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_219f732381d2fb9e, []int{2}
}
func (m *DepositRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositRecord.Merge(m, src)
}
func (m *DepositRecord) XXX_Size() int {
	return m.Size()
}
func (m *DepositRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositRecord.DiscardUnknown(m)
}

var xxx_messageInfo_DepositRecord proto.InternalMessageInfo

func (m *DepositRecord) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DepositRecord) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DepositRecord) GetHostChainId() string {
	if m != nil {
		return m.HostChainId
	}
	return ""
}

func (m *DepositRecord) GetStatus() DepositRecord_Status {
	if m != nil {
		return m.Status
	}
	return DepositRecord_TRANSFER_QUEUE
}

func (m *DepositRecord) GetDepositEpochNumber() uint64 {
	if m != nil {
		return m.DepositEpochNumber
	}
	return 0
}

func (m *DepositRecord) GetSource() DepositRecord_Source {
	if m != nil {
		return m.Source
	}
	return DepositRecord_SIDE
}

func init() {
	proto.RegisterEnum("side.yield.DepositRecord_Status", DepositRecord_Status_name, DepositRecord_Status_value)
	proto.RegisterEnum("side.yield.DepositRecord_Source", DepositRecord_Source_name, DepositRecord_Source_value)
	proto.RegisterType((*MsgRegisterHostChain)(nil), "side.yield.MsgRegisterHostChain")
	proto.RegisterType((*MsgRegisterHostChainResponse)(nil), "side.yield.MsgRegisterHostChainResponse")
	proto.RegisterType((*DepositRecord)(nil), "side.yield.DepositRecord")
}

func init() { proto.RegisterFile("side/yield/tx.proto", fileDescriptor_219f732381d2fb9e) }

var fileDescriptor_219f732381d2fb9e = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x41, 0x4f, 0xdb, 0x48,
	0x14, 0x8e, 0x93, 0x10, 0xc2, 0x5b, 0x40, 0x61, 0xc8, 0xae, 0xbc, 0x61, 0x65, 0x23, 0x1f, 0x56,
	0x68, 0xa5, 0xb5, 0x17, 0xd8, 0xc3, 0x0a, 0x69, 0x0f, 0x84, 0x18, 0x70, 0x55, 0x42, 0x3b, 0x81,
	0x4b, 0x2f, 0x96, 0x3d, 0x1e, 0x12, 0xab, 0x89, 0x27, 0xf2, 0x4c, 0x2a, 0xf8, 0x07, 0xbd, 0xb5,
	0x3f, 0x81, 0x9f, 0xc3, 0x91, 0x53, 0x55, 0xf5, 0x10, 0x55, 0x70, 0xe9, 0x39, 0xe7, 0x1e, 0x2a,
	0xcf, 0x38, 0x90, 0xa8, 0x48, 0x9c, 0x32, 0xef, 0xfb, 0xbe, 0xf7, 0xe6, 0xbd, 0x2f, 0xcf, 0x03,
	0xeb, 0x3c, 0x8e, 0xa8, 0x73, 0x15, 0xd3, 0x7e, 0xe4, 0x88, 0x4b, 0x7b, 0x98, 0x32, 0xc1, 0x10,
	0x64, 0xa0, 0x2d, 0xc1, 0x46, 0xbd, 0xcb, 0xba, 0x4c, 0xc2, 0x4e, 0x76, 0x52, 0x8a, 0xc6, 0xef,
	0x84, 0xf1, 0x01, 0xe3, 0xbe, 0x22, 0x54, 0x90, 0x53, 0x86, 0x8a, 0x9c, 0x30, 0xe0, 0xd4, 0x79,
	0xb7, 0x1d, 0x52, 0x11, 0x6c, 0x3b, 0x84, 0xc5, 0x89, 0xe2, 0xad, 0x4f, 0x45, 0xa8, 0x9f, 0xf0,
	0x2e, 0xa6, 0xdd, 0x98, 0x0b, 0x9a, 0x1e, 0x33, 0x2e, 0x0e, 0x7a, 0x41, 0x9c, 0xa0, 0xff, 0x61,
	0x85, 0xb0, 0x24, 0xa1, 0x44, 0xc4, 0x2c, 0xf1, 0xe3, 0x48, 0xd7, 0x36, 0xb5, 0xad, 0xa5, 0xa6,
	0x3e, 0x19, 0x9b, 0xf5, 0xab, 0x60, 0xd0, 0xdf, 0xb3, 0xe6, 0x68, 0x0b, 0x2f, 0x3f, 0xc6, 0x5e,
	0x84, 0x2c, 0x58, 0x0e, 0x29, 0xe9, 0xed, 0xee, 0x0c, 0x53, 0x7a, 0x11, 0x5f, 0xea, 0xc5, 0x2c,
	0x1b, 0xcf, 0x61, 0xe8, 0x5f, 0x80, 0x1e, 0xe3, 0xc2, 0x8f, 0x68, 0xc2, 0x06, 0x7a, 0x49, 0xd6,
	0xff, 0x75, 0x32, 0x36, 0xd7, 0x54, 0xfd, 0x47, 0xce, 0xc2, 0x4b, 0x59, 0xd0, 0xca, 0xce, 0x68,
	0x1b, 0x96, 0xe2, 0x90, 0xe4, 0x49, 0x65, 0x99, 0x54, 0x9f, 0x8c, 0xcd, 0x9a, 0x4a, 0x7a, 0xa0,
	0x2c, 0x5c, 0x8d, 0x43, 0xa2, 0x52, 0x74, 0x58, 0x24, 0x29, 0x0d, 0x04, 0x4b, 0xf5, 0x05, 0xd9,
	0xc7, 0x34, 0x44, 0x6d, 0x58, 0x17, 0x69, 0x90, 0xf0, 0x0b, 0x9a, 0xfa, 0xa4, 0x17, 0x24, 0x09,
	0xed, 0x67, 0xb3, 0x56, 0x64, 0x59, 0x63, 0x32, 0x36, 0x1b, 0xaa, 0xec, 0x13, 0x22, 0x0b, 0xaf,
	0x4d, 0xd1, 0x03, 0x05, 0x7a, 0xd1, 0x5e, 0xf5, 0xfd, 0xb5, 0x59, 0xf8, 0x76, 0x6d, 0x16, 0x2c,
	0x03, 0xfe, 0x78, 0xca, 0x57, 0x4c, 0xf9, 0x90, 0x25, 0x9c, 0x5a, 0xdf, 0x4b, 0xb0, 0xd2, 0xa2,
	0x43, 0xc6, 0x63, 0x81, 0x29, 0x61, 0x69, 0x84, 0x56, 0xa1, 0x98, 0xdb, 0x5c, 0xc6, 0xc5, 0x38,
	0x42, 0x87, 0x50, 0x09, 0x06, 0x6c, 0x94, 0x08, 0x65, 0x5e, 0xd3, 0xbe, 0x19, 0x9b, 0x85, 0x2f,
	0x63, 0xf3, 0xcf, 0x6e, 0x2c, 0x7a, 0xa3, 0xd0, 0x26, 0x6c, 0x90, 0xff, 0xd7, 0xf9, 0xcf, 0xdf,
	0x3c, 0x7a, 0xeb, 0x88, 0xab, 0x21, 0xe5, 0xb6, 0x97, 0x08, 0x9c, 0x67, 0xa3, 0x3a, 0x2c, 0xcc,
	0x38, 0x8c, 0x55, 0x80, 0x2c, 0x58, 0x91, 0x06, 0x93, 0xac, 0xab, 0x6c, 0x66, 0x69, 0x25, 0xfe,
	0xa5, 0x37, 0xed, 0xd4, 0x8b, 0xd0, 0x7f, 0x50, 0xe1, 0x22, 0x10, 0x23, 0x2e, 0x0d, 0x59, 0xdd,
	0xd9, 0xb4, 0x1f, 0x57, 0xd1, 0x9e, 0x6b, 0xde, 0xee, 0x48, 0x1d, 0xce, 0xf5, 0xe8, 0x1f, 0xa8,
	0x47, 0x8a, 0xf7, 0xe9, 0x90, 0x91, 0x9e, 0x9f, 0x8c, 0x06, 0x21, 0x4d, 0xf5, 0x45, 0x39, 0x1d,
	0xca, 0x39, 0x37, 0xa3, 0xda, 0x92, 0x91, 0x77, 0xb1, 0x51, 0x4a, 0xa8, 0x5e, 0x7d, 0xf6, 0x2e,
	0xa9, 0xc3, 0xb9, 0xde, 0xfa, 0xa0, 0x41, 0x45, 0x5d, 0x8f, 0x10, 0xac, 0x9e, 0xe1, 0xfd, 0x76,
	0xe7, 0xd0, 0xc5, 0xfe, 0xeb, 0x73, 0xf7, 0xdc, 0xad, 0x15, 0x90, 0x01, 0x8d, 0x07, 0xec, 0xd0,
	0xc3, 0x9d, 0x33, 0xdf, 0x6b, 0xfb, 0xaf, 0xf0, 0xe9, 0x11, 0x76, 0x3b, 0x9d, 0x5a, 0x11, 0x99,
	0xb0, 0xf1, 0xc0, 0x77, 0xdc, 0x83, 0xd3, 0x76, 0x6b, 0x4e, 0x50, 0x42, 0x75, 0xa8, 0xb5, 0xdc,
	0x97, 0xee, 0xd1, 0xfe, 0x99, 0x77, 0xda, 0xce, 0xcb, 0x6a, 0xa8, 0x01, 0xbf, 0xcd, 0xa0, 0xb3,
	0x19, 0x65, 0x6b, 0x03, 0x2a, 0xaa, 0x47, 0x54, 0x85, 0x72, 0xc7, 0x6b, 0x65, 0x6d, 0x2c, 0x42,
	0xe9, 0xf8, 0xbc, 0x59, 0xd3, 0x5e, 0x94, 0xab, 0x0b, 0xb5, 0xca, 0xce, 0x05, 0x94, 0x4e, 0x78,
	0x17, 0xf9, 0xb0, 0xf6, 0xf3, 0xa7, 0x37, 0x37, 0xfa, 0x53, 0x4b, 0xd4, 0xd8, 0x7a, 0x4e, 0x31,
	0x5d, 0xb3, 0x66, 0xeb, 0xe6, 0xce, 0xd0, 0x6e, 0xef, 0x0c, 0xed, 0xeb, 0x9d, 0xa1, 0x7d, 0xbc,
	0x37, 0x0a, 0xb7, 0xf7, 0x46, 0xe1, 0xf3, 0xbd, 0x51, 0x78, 0xf3, 0xd7, 0xcc, 0x1a, 0x65, 0xd5,
	0xe4, 0x7b, 0x40, 0x58, 0x5f, 0x06, 0xce, 0xe5, 0xf4, 0x15, 0xca, 0xd6, 0x29, 0xac, 0x48, 0x72,
	0xf7, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x3c, 0x69, 0x71, 0xa0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RegisterHostChain(ctx context.Context, in *MsgRegisterHostChain, opts ...grpc.CallOption) (*MsgRegisterHostChainResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterHostChain(ctx context.Context, in *MsgRegisterHostChain, opts ...grpc.CallOption) (*MsgRegisterHostChainResponse, error) {
	out := new(MsgRegisterHostChainResponse)
	err := c.cc.Invoke(ctx, "/side.yield.Msg/RegisterHostChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RegisterHostChain(context.Context, *MsgRegisterHostChain) (*MsgRegisterHostChainResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterHostChain(ctx context.Context, req *MsgRegisterHostChain) (*MsgRegisterHostChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterHostChain not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterHostChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterHostChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterHostChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/side.yield.Msg/RegisterHostChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterHostChain(ctx, req.(*MsgRegisterHostChain))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "side.yield.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterHostChain",
			Handler:    _Msg_RegisterHostChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "side/yield/tx.proto",
}

func (m *MsgRegisterHostChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterHostChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterHostChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TransferChannelId) > 0 {
		i -= len(m.TransferChannelId)
		copy(dAtA[i:], m.TransferChannelId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TransferChannelId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.IbcDenom) > 0 {
		i -= len(m.IbcDenom)
		copy(dAtA[i:], m.IbcDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.IbcDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.HostDenom) > 0 {
		i -= len(m.HostDenom)
		copy(dAtA[i:], m.HostDenom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HostDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bech32Prefix) > 0 {
		i -= len(m.Bech32Prefix)
		copy(dAtA[i:], m.Bech32Prefix)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Bech32Prefix)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRegisterHostChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterHostChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterHostChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DepositRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Source != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Source))
		i--
		dAtA[i] = 0x40
	}
	if m.DepositEpochNumber != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.DepositEpochNumber))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.HostChainId) > 0 {
		i -= len(m.HostChainId)
		copy(dAtA[i:], m.HostChainId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.HostChainId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRegisterHostChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Bech32Prefix)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HostDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.IbcDenom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TransferChannelId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterHostChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DepositRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.HostChainId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovTx(uint64(m.Status))
	}
	if m.DepositEpochNumber != 0 {
		n += 1 + sovTx(uint64(m.DepositEpochNumber))
	}
	if m.Source != 0 {
		n += 1 + sovTx(uint64(m.Source))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRegisterHostChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterHostChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterHostChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bech32Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bech32Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TransferChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRegisterHostChainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRegisterHostChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterHostChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DepositRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DepositRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= DepositRecord_Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositEpochNumber", wireType)
			}
			m.DepositEpochNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DepositEpochNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			m.Source = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Source |= DepositRecord_Source(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
