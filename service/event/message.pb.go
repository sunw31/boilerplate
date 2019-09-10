// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/event/message.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Event is Entity.
type Event struct {
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// origin is an identity of the operator who entered the event.
	Origin               string       `protobuf:"bytes,2,opt,name=origin,proto3" json:"origin,omitempty"`
	TimePoint            []*TimePoint `protobuf:"bytes,3,rep,name=time_point,json=timePoint,proto3" json:"time_point,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73432f895dbcfa7, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Event) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Event) GetTimePoint() []*TimePoint {
	if m != nil {
		return m.TimePoint
	}
	return nil
}

// TimePoint is Entity.
type TimePoint struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Tag                  string               `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimePoint) Reset()         { *m = TimePoint{} }
func (m *TimePoint) String() string { return proto.CompactTextString(m) }
func (*TimePoint) ProtoMessage()    {}
func (*TimePoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a73432f895dbcfa7, []int{1}
}

func (m *TimePoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimePoint.Unmarshal(m, b)
}
func (m *TimePoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimePoint.Marshal(b, m, deterministic)
}
func (m *TimePoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimePoint.Merge(m, src)
}
func (m *TimePoint) XXX_Size() int {
	return xxx_messageInfo_TimePoint.Size(m)
}
func (m *TimePoint) XXX_DiscardUnknown() {
	xxx_messageInfo_TimePoint.DiscardUnknown(m)
}

var xxx_messageInfo_TimePoint proto.InternalMessageInfo

func (m *TimePoint) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *TimePoint) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "michilu.boilerplate.service.event.Event")
	proto.RegisterType((*TimePoint)(nil), "michilu.boilerplate.service.event.TimePoint")
}

func init() { proto.RegisterFile("service/event/message.proto", fileDescriptor_a73432f895dbcfa7) }

var fileDescriptor_a73432f895dbcfa7 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xc1, 0x4e, 0xb3, 0x40,
	0x14, 0x85, 0xff, 0xa1, 0x29, 0x7f, 0x98, 0xba, 0x9a, 0x85, 0x21, 0x98, 0x58, 0xec, 0x8a, 0x85,
	0x9d, 0x49, 0x6a, 0x62, 0x5c, 0x37, 0x71, 0x6f, 0xd0, 0x85, 0x71, 0x53, 0x07, 0x18, 0xa7, 0x37,
	0x32, 0x0c, 0x81, 0x0b, 0x7d, 0x11, 0xdf, 0xcf, 0xc4, 0x27, 0x31, 0x0c, 0x2d, 0x26, 0x6e, 0xdc,
	0xc1, 0xb9, 0xe7, 0x7c, 0x73, 0x0e, 0xbd, 0x68, 0x55, 0xd3, 0x43, 0xae, 0x84, 0xea, 0x55, 0x85,
	0xc2, 0xa8, 0xb6, 0x95, 0x5a, 0xf1, 0xba, 0xb1, 0x68, 0xd9, 0x95, 0x81, 0x7c, 0x0f, 0x65, 0xc7,
	0x33, 0x0b, 0xa5, 0x6a, 0xea, 0x52, 0xa2, 0xe2, 0xc7, 0x00, 0x77, 0x81, 0xe8, 0x56, 0x03, 0xee,
	0xbb, 0x8c, 0xe7, 0xd6, 0x08, 0x73, 0x00, 0x7c, 0xb7, 0x07, 0xa1, 0xed, 0xda, 0xe5, 0xd7, 0xbd,
	0x2c, 0xa1, 0x90, 0x68, 0x9b, 0x56, 0x4c, 0x9f, 0x23, 0x3a, 0x5a, 0x6a, 0x6b, 0x75, 0xa9, 0x84,
	0xfb, 0xcb, 0xba, 0x37, 0x81, 0x60, 0x54, 0x8b, 0xd2, 0xd4, 0xa3, 0x61, 0xf5, 0x41, 0xe8, 0xfc,
	0x7e, 0x78, 0x82, 0x9d, 0x53, 0x0f, 0x8a, 0x90, 0xc4, 0x24, 0x39, 0xdb, 0xfa, 0x5f, 0x9f, 0x4b,
	0xaf, 0xfe, 0x97, 0x7a, 0x50, 0xb0, 0x4b, 0xea, 0xdb, 0x06, 0x34, 0x54, 0xa1, 0x17, 0x93, 0x24,
	0x18, 0x6f, 0xcf, 0x24, 0x3d, 0xaa, 0xec, 0x91, 0xd2, 0x01, 0xba, 0xab, 0x2d, 0x54, 0x18, 0xce,
	0xe2, 0x59, 0xb2, 0xd8, 0x5c, 0xf3, 0x3f, 0x27, 0xf1, 0x27, 0x30, 0xea, 0x61, 0xc8, 0x8c, 0xc4,
	0x57, 0x92, 0x06, 0x78, 0x92, 0x56, 0x3b, 0x1a, 0x4c, 0x77, 0x76, 0x47, 0x83, 0xa9, 0xb6, 0x2b,
	0xb8, 0xd8, 0x44, 0x7c, 0x1c, 0xc6, 0x4f, 0xc3, 0x1c, 0xce, 0x39, 0xd2, 0x1f, 0x33, 0x0b, 0xe9,
	0x0c, 0xa5, 0xfe, 0x55, 0x7c, 0x90, 0xb6, 0xff, 0x5f, 0xe6, 0xae, 0x46, 0xe6, 0x3b, 0xc2, 0xcd,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x82, 0x3c, 0xf6, 0xa2, 0x01, 0x00, 0x00,
}
