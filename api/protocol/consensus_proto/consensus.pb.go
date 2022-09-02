// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consensus.proto

package consensus_pb

import (
	fmt "fmt"
	github_com_amazechain_amc_common_types "github.com/amazechain/amc/common/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type PBSigners struct {
	Signer               []*PBSigner `protobuf:"bytes,1,rep,name=signer,proto3" json:"signer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PBSigners) Reset()         { *m = PBSigners{} }
func (m *PBSigners) String() string { return proto.CompactTextString(m) }
func (*PBSigners) ProtoMessage()    {}
func (*PBSigners) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{0}
}
func (m *PBSigners) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBSigners.Unmarshal(m, b)
}
func (m *PBSigners) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBSigners.Marshal(b, m, deterministic)
}
func (m *PBSigners) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBSigners.Merge(m, src)
}
func (m *PBSigners) XXX_Size() int {
	return xxx_messageInfo_PBSigners.Size(m)
}
func (m *PBSigners) XXX_DiscardUnknown() {
	xxx_messageInfo_PBSigners.DiscardUnknown(m)
}

var xxx_messageInfo_PBSigners proto.InternalMessageInfo

func (m *PBSigners) GetSigner() []*PBSigner {
	if m != nil {
		return m.Signer
	}
	return nil
}

type PBSigner struct {
	Public               string                                         `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Address              github_com_amazechain_amc_common_types.Address `protobuf:"bytes,2,opt,name=address,proto3,customtype=github.com/amazechain/amc/common/types.Address" json:"address"`
	Version              string                                         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *PBSigner) Reset()         { *m = PBSigner{} }
func (m *PBSigner) String() string { return proto.CompactTextString(m) }
func (*PBSigner) ProtoMessage()    {}
func (*PBSigner) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{1}
}
func (m *PBSigner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBSigner.Unmarshal(m, b)
}
func (m *PBSigner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBSigner.Marshal(b, m, deterministic)
}
func (m *PBSigner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBSigner.Merge(m, src)
}
func (m *PBSigner) XXX_Size() int {
	return xxx_messageInfo_PBSigner.Size(m)
}
func (m *PBSigner) XXX_DiscardUnknown() {
	xxx_messageInfo_PBSigner.DiscardUnknown(m)
}

var xxx_messageInfo_PBSigner proto.InternalMessageInfo

func (m *PBSigner) GetPublic() string {
	if m != nil {
		return m.Public
	}
	return ""
}

func (m *PBSigner) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type PBVote struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBVote) Reset()         { *m = PBVote{} }
func (m *PBVote) String() string { return proto.CompactTextString(m) }
func (*PBVote) ProtoMessage()    {}
func (*PBVote) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{2}
}
func (m *PBVote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBVote.Unmarshal(m, b)
}
func (m *PBVote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBVote.Marshal(b, m, deterministic)
}
func (m *PBVote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBVote.Merge(m, src)
}
func (m *PBVote) XXX_Size() int {
	return xxx_messageInfo_PBVote.Size(m)
}
func (m *PBVote) XXX_DiscardUnknown() {
	xxx_messageInfo_PBVote.DiscardUnknown(m)
}

var xxx_messageInfo_PBVote proto.InternalMessageInfo

type PBPoaInfo struct {
	Public               string   `protobuf:"bytes,1,opt,name=public,proto3" json:"public,omitempty"`
	Sign                 []byte   `protobuf:"bytes,2,opt,name=sign,proto3" json:"sign,omitempty"`
	Type                 int64    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBPoaInfo) Reset()         { *m = PBPoaInfo{} }
func (m *PBPoaInfo) String() string { return proto.CompactTextString(m) }
func (*PBPoaInfo) ProtoMessage()    {}
func (*PBPoaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f0f2c53b3de771, []int{3}
}
func (m *PBPoaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBPoaInfo.Unmarshal(m, b)
}
func (m *PBPoaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBPoaInfo.Marshal(b, m, deterministic)
}
func (m *PBPoaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBPoaInfo.Merge(m, src)
}
func (m *PBPoaInfo) XXX_Size() int {
	return xxx_messageInfo_PBPoaInfo.Size(m)
}
func (m *PBPoaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PBPoaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PBPoaInfo proto.InternalMessageInfo

func (m *PBPoaInfo) GetPublic() string {
	if m != nil {
		return m.Public
	}
	return ""
}

func (m *PBPoaInfo) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *PBPoaInfo) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*PBSigners)(nil), "consensus.pb.PBSigners")
	proto.RegisterType((*PBSigner)(nil), "consensus.pb.PBSigner")
	proto.RegisterType((*PBVote)(nil), "consensus.pb.PBVote")
	proto.RegisterType((*PBPoaInfo)(nil), "consensus.pb.PBPoaInfo")
}

func init() { proto.RegisterFile("consensus.proto", fileDescriptor_56f0f2c53b3de771) }

var fileDescriptor_56f0f2c53b3de771 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0xad, 0x95, 0xce, 0x4c, 0x1c, 0x10, 0x82, 0x0c, 0xc1, 0xcd, 0x94, 0xae, 0xba, 0x4a,
	0x40, 0xc1, 0x8d, 0x2b, 0xbb, 0x13, 0x37, 0x25, 0x82, 0xfb, 0x34, 0x13, 0x3b, 0x01, 0x93, 0x5b,
	0x7a, 0x5b, 0x41, 0x1f, 0xc2, 0xe7, 0xf2, 0x19, 0x5c, 0xcc, 0xb3, 0x48, 0xd2, 0x29, 0xcc, 0xc6,
	0xdd, 0x77, 0x4f, 0x4e, 0xce, 0xfd, 0x21, 0x57, 0x1a, 0x3c, 0x1a, 0x8f, 0x23, 0xf2, 0xae, 0x87,
	0x01, 0xe8, 0xfa, 0x44, 0x68, 0x6e, 0xae, 0x5b, 0x68, 0x21, 0x3e, 0x88, 0x40, 0x93, 0xa7, 0x78,
	0x20, 0xab, 0xba, 0x7a, 0xb1, 0xad, 0x37, 0x3d, 0x52, 0x4e, 0x32, 0x8c, 0xc8, 0x92, 0x3c, 0x2d,
	0x2f, 0x6f, 0x37, 0xfc, 0x34, 0x81, 0xcf, 0x46, 0x79, 0x74, 0x15, 0xdf, 0x09, 0x59, 0xce, 0x22,
	0xdd, 0x90, 0xac, 0x1b, 0x9b, 0x77, 0xab, 0x59, 0x92, 0x27, 0xe5, 0x4a, 0x1e, 0x2b, 0x5a, 0x93,
	0x85, 0xda, 0xed, 0x7a, 0x83, 0xc8, 0xce, 0xf3, 0xa4, 0x5c, 0x57, 0xf7, 0x3f, 0x87, 0xed, 0xd9,
	0xef, 0x61, 0xcb, 0x5b, 0x3b, 0xec, 0xc7, 0x86, 0x6b, 0x70, 0x42, 0x39, 0xf5, 0x65, 0xf4, 0x5e,
	0x59, 0x2f, 0x94, 0xd3, 0x42, 0x83, 0x73, 0xe0, 0xc5, 0xf0, 0xd9, 0x19, 0xe4, 0x8f, 0xd3, 0x6f,
	0x39, 0xc7, 0x50, 0x46, 0x16, 0x1f, 0xa6, 0x47, 0x0b, 0x9e, 0xa5, 0xb1, 0xd5, 0x5c, 0x16, 0x4b,
	0x92, 0xd5, 0xd5, 0x2b, 0x0c, 0xa6, 0x78, 0x0e, 0x7b, 0xd5, 0xa0, 0x9e, 0xfc, 0x1b, 0xfc, 0x3b,
	0x1a, 0x25, 0x17, 0x61, 0x93, 0x69, 0x2e, 0x19, 0x39, 0x68, 0xa1, 0x6d, 0x4c, 0x4e, 0x65, 0xe4,
	0x26, 0x8b, 0xb7, 0xba, 0xfb, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x95, 0xdc, 0xc8, 0xd4, 0x62, 0x01,
	0x00, 0x00,
}
