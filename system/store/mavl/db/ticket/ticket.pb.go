// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ticket.proto

package ticket

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Ticket struct {
	TicketId string `protobuf:"bytes,1,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	// 0 -> 未成熟 1 -> 可挖矿 2 -> 已挖成功 3-> 已关闭
	Status int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// genesis 创建的私钥比较特殊
	IsGenesis bool `protobuf:"varint,3,opt,name=isGenesis,proto3" json:"isGenesis,omitempty"`
	//创建时间
	CreateTime int64 `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	//挖矿时间
	MinerTime int64 `protobuf:"varint,5,opt,name=minerTime,proto3" json:"minerTime,omitempty"`
	//挖到的币的数目
	MinerValue   int64  `protobuf:"varint,8,opt,name=minerValue,proto3" json:"minerValue,omitempty"`
	MinerAddress string `protobuf:"bytes,6,opt,name=minerAddress,proto3" json:"minerAddress,omitempty"`
	// return wallet
	ReturnAddress string `protobuf:"bytes,7,opt,name=returnAddress,proto3" json:"returnAddress,omitempty"`
	// miner Price
	Price                int64    `protobuf:"varint,9,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ticket) Reset()         { *m = Ticket{} }
func (m *Ticket) String() string { return proto.CompactTextString(m) }
func (*Ticket) ProtoMessage()    {}
func (*Ticket) Descriptor() ([]byte, []int) {
	return fileDescriptor_98a6c21780e82d22, []int{0}
}

func (m *Ticket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticket.Unmarshal(m, b)
}
func (m *Ticket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticket.Marshal(b, m, deterministic)
}
func (m *Ticket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticket.Merge(m, src)
}
func (m *Ticket) XXX_Size() int {
	return xxx_messageInfo_Ticket.Size(m)
}
func (m *Ticket) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticket.DiscardUnknown(m)
}

var xxx_messageInfo_Ticket proto.InternalMessageInfo

func (m *Ticket) GetTicketId() string {
	if m != nil {
		return m.TicketId
	}
	return ""
}

func (m *Ticket) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Ticket) GetIsGenesis() bool {
	if m != nil {
		return m.IsGenesis
	}
	return false
}

func (m *Ticket) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Ticket) GetMinerTime() int64 {
	if m != nil {
		return m.MinerTime
	}
	return 0
}

func (m *Ticket) GetMinerValue() int64 {
	if m != nil {
		return m.MinerValue
	}
	return 0
}

func (m *Ticket) GetMinerAddress() string {
	if m != nil {
		return m.MinerAddress
	}
	return ""
}

func (m *Ticket) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

func (m *Ticket) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func init() {
	proto.RegisterType((*Ticket)(nil), "ticket.Ticket")
}

func init() {
	proto.RegisterFile("ticket.proto", fileDescriptor_98a6c21780e82d22)
}

var fileDescriptor_98a6c21780e82d22 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x49, 0x6b, 0x63, 0x7b, 0xa9, 0x9b, 0x8b, 0x48, 0x10, 0x91, 0x50, 0x5c, 0x64, 0xe5,
	0xc6, 0x27, 0x70, 0x25, 0x6e, 0x43, 0x71, 0x5f, 0xdb, 0xbb, 0x08, 0xda, 0x1f, 0x92, 0xf4, 0x5d,
	0x7c, 0x5c, 0xe9, 0xcd, 0xcc, 0x74, 0xba, 0xcb, 0xf7, 0x9d, 0x73, 0x48, 0x08, 0xd4, 0xd1, 0xf5,
	0x3f, 0x14, 0x5f, 0x17, 0x3f, 0xc7, 0x19, 0x65, 0xa2, 0xe6, 0x2f, 0x03, 0xd9, 0xf2, 0x11, 0x1f,
	0xa1, 0x4c, 0xf2, 0x73, 0x50, 0x42, 0x0b, 0x53, 0xd9, 0x0b, 0xe3, 0x03, 0xc8, 0x10, 0xbb, 0xb8,
	0x06, 0x95, 0x69, 0x61, 0x0a, 0x7b, 0x22, 0x7c, 0x82, 0xca, 0x85, 0x0f, 0x9a, 0x28, 0xb8, 0xa0,
	0x72, 0x2d, 0x4c, 0x69, 0x77, 0x81, 0xcf, 0x00, 0xbd, 0xa7, 0x2e, 0x52, 0xeb, 0x46, 0x52, 0x37,
	0x5a, 0x98, 0xdc, 0x5e, 0x99, 0x6d, 0x3d, 0xba, 0x89, 0x3c, 0xc7, 0x05, 0xc7, 0xbb, 0xd8, 0xd6,
	0x0c, 0x5f, 0xdd, 0xef, 0x4a, 0xaa, 0x4c, 0xeb, 0xdd, 0x60, 0x03, 0x35, 0xd3, 0xfb, 0x30, 0x78,
	0x0a, 0x41, 0x49, 0x7e, 0xf3, 0xc1, 0xe1, 0x0b, 0xdc, 0x79, 0x8a, 0xab, 0x9f, 0xce, 0xa5, 0x5b,
	0x2e, 0x1d, 0x25, 0xde, 0x43, 0xb1, 0x78, 0xd7, 0x93, 0xaa, 0xf8, 0x92, 0x04, 0xdf, 0x92, 0x7f,
	0xea, 0xed, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xb4, 0xe8, 0x04, 0x39, 0x01, 0x00, 0x00,
}
