// Code generated by protoc-gen-go. DO NOT EDIT.
// source: room_service.proto

package room_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateRoomMsg struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NrOfSeats            int32    `protobuf:"varint,2,opt,name=nrOfSeats,proto3" json:"nrOfSeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomMsg) Reset()         { *m = CreateRoomMsg{} }
func (m *CreateRoomMsg) String() string { return proto.CompactTextString(m) }
func (*CreateRoomMsg) ProtoMessage()    {}
func (*CreateRoomMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{0}
}

func (m *CreateRoomMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomMsg.Unmarshal(m, b)
}
func (m *CreateRoomMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomMsg.Marshal(b, m, deterministic)
}
func (m *CreateRoomMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomMsg.Merge(m, src)
}
func (m *CreateRoomMsg) XXX_Size() int {
	return xxx_messageInfo_CreateRoomMsg.Size(m)
}
func (m *CreateRoomMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomMsg proto.InternalMessageInfo

func (m *CreateRoomMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRoomMsg) GetNrOfSeats() int32 {
	if m != nil {
		return m.NrOfSeats
	}
	return 0
}

type CreateRoomResponseMsg struct {
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomResponseMsg) Reset()         { *m = CreateRoomResponseMsg{} }
func (m *CreateRoomResponseMsg) String() string { return proto.CompactTextString(m) }
func (*CreateRoomResponseMsg) ProtoMessage()    {}
func (*CreateRoomResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{1}
}

func (m *CreateRoomResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomResponseMsg.Unmarshal(m, b)
}
func (m *CreateRoomResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomResponseMsg.Marshal(b, m, deterministic)
}
func (m *CreateRoomResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomResponseMsg.Merge(m, src)
}
func (m *CreateRoomResponseMsg) XXX_Size() int {
	return xxx_messageInfo_CreateRoomResponseMsg.Size(m)
}
func (m *CreateRoomResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomResponseMsg proto.InternalMessageInfo

func (m *CreateRoomResponseMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteRoomMsg struct {
	Id                   int32    `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoomMsg) Reset()         { *m = DeleteRoomMsg{} }
func (m *DeleteRoomMsg) String() string { return proto.CompactTextString(m) }
func (*DeleteRoomMsg) ProtoMessage()    {}
func (*DeleteRoomMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{2}
}

func (m *DeleteRoomMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoomMsg.Unmarshal(m, b)
}
func (m *DeleteRoomMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoomMsg.Marshal(b, m, deterministic)
}
func (m *DeleteRoomMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoomMsg.Merge(m, src)
}
func (m *DeleteRoomMsg) XXX_Size() int {
	return xxx_messageInfo_DeleteRoomMsg.Size(m)
}
func (m *DeleteRoomMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoomMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoomMsg proto.InternalMessageInfo

func (m *DeleteRoomMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteRoomResponseMsg struct {
	Success              bool     `protobuf:"varint,5,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRoomResponseMsg) Reset()         { *m = DeleteRoomResponseMsg{} }
func (m *DeleteRoomResponseMsg) String() string { return proto.CompactTextString(m) }
func (*DeleteRoomResponseMsg) ProtoMessage()    {}
func (*DeleteRoomResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{3}
}

func (m *DeleteRoomResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRoomResponseMsg.Unmarshal(m, b)
}
func (m *DeleteRoomResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRoomResponseMsg.Marshal(b, m, deterministic)
}
func (m *DeleteRoomResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRoomResponseMsg.Merge(m, src)
}
func (m *DeleteRoomResponseMsg) XXX_Size() int {
	return xxx_messageInfo_DeleteRoomResponseMsg.Size(m)
}
func (m *DeleteRoomResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRoomResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRoomResponseMsg proto.InternalMessageInfo

func (m *DeleteRoomResponseMsg) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type GetRoomMsg struct {
	Id                   int32    `protobuf:"varint,6,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoomMsg) Reset()         { *m = GetRoomMsg{} }
func (m *GetRoomMsg) String() string { return proto.CompactTextString(m) }
func (*GetRoomMsg) ProtoMessage()    {}
func (*GetRoomMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{4}
}

func (m *GetRoomMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomMsg.Unmarshal(m, b)
}
func (m *GetRoomMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomMsg.Marshal(b, m, deterministic)
}
func (m *GetRoomMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomMsg.Merge(m, src)
}
func (m *GetRoomMsg) XXX_Size() int {
	return xxx_messageInfo_GetRoomMsg.Size(m)
}
func (m *GetRoomMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomMsg proto.InternalMessageInfo

func (m *GetRoomMsg) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetRoomResponseMsg struct {
	Room                 *RoomData `protobuf:"bytes,7,opt,name=room,proto3" json:"room,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetRoomResponseMsg) Reset()         { *m = GetRoomResponseMsg{} }
func (m *GetRoomResponseMsg) String() string { return proto.CompactTextString(m) }
func (*GetRoomResponseMsg) ProtoMessage()    {}
func (*GetRoomResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{5}
}

func (m *GetRoomResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomResponseMsg.Unmarshal(m, b)
}
func (m *GetRoomResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomResponseMsg.Marshal(b, m, deterministic)
}
func (m *GetRoomResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomResponseMsg.Merge(m, src)
}
func (m *GetRoomResponseMsg) XXX_Size() int {
	return xxx_messageInfo_GetRoomResponseMsg.Size(m)
}
func (m *GetRoomResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomResponseMsg proto.InternalMessageInfo

func (m *GetRoomResponseMsg) GetRoom() *RoomData {
	if m != nil {
		return m.Room
	}
	return nil
}

type GetRoomsMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRoomsMsg) Reset()         { *m = GetRoomsMsg{} }
func (m *GetRoomsMsg) String() string { return proto.CompactTextString(m) }
func (*GetRoomsMsg) ProtoMessage()    {}
func (*GetRoomsMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{6}
}

func (m *GetRoomsMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomsMsg.Unmarshal(m, b)
}
func (m *GetRoomsMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomsMsg.Marshal(b, m, deterministic)
}
func (m *GetRoomsMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomsMsg.Merge(m, src)
}
func (m *GetRoomsMsg) XXX_Size() int {
	return xxx_messageInfo_GetRoomsMsg.Size(m)
}
func (m *GetRoomsMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomsMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomsMsg proto.InternalMessageInfo

type GetRoomsResponseMsg struct {
	Rooms                []*RoomData `protobuf:"bytes,8,rep,name=Rooms,proto3" json:"Rooms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetRoomsResponseMsg) Reset()         { *m = GetRoomsResponseMsg{} }
func (m *GetRoomsResponseMsg) String() string { return proto.CompactTextString(m) }
func (*GetRoomsResponseMsg) ProtoMessage()    {}
func (*GetRoomsResponseMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{7}
}

func (m *GetRoomsResponseMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRoomsResponseMsg.Unmarshal(m, b)
}
func (m *GetRoomsResponseMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRoomsResponseMsg.Marshal(b, m, deterministic)
}
func (m *GetRoomsResponseMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRoomsResponseMsg.Merge(m, src)
}
func (m *GetRoomsResponseMsg) XXX_Size() int {
	return xxx_messageInfo_GetRoomsResponseMsg.Size(m)
}
func (m *GetRoomsResponseMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRoomsResponseMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GetRoomsResponseMsg proto.InternalMessageInfo

func (m *GetRoomsResponseMsg) GetRooms() []*RoomData {
	if m != nil {
		return m.Rooms
	}
	return nil
}

type RoomData struct {
	Name                 string   `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32    `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	NrOfSeats            int32    `protobuf:"varint,11,opt,name=nrOfSeats,proto3" json:"nrOfSeats,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomData) Reset()         { *m = RoomData{} }
func (m *RoomData) String() string { return proto.CompactTextString(m) }
func (*RoomData) ProtoMessage()    {}
func (*RoomData) Descriptor() ([]byte, []int) {
	return fileDescriptor_36c4a5b3f0e57253, []int{8}
}

func (m *RoomData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomData.Unmarshal(m, b)
}
func (m *RoomData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomData.Marshal(b, m, deterministic)
}
func (m *RoomData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomData.Merge(m, src)
}
func (m *RoomData) XXX_Size() int {
	return xxx_messageInfo_RoomData.Size(m)
}
func (m *RoomData) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomData.DiscardUnknown(m)
}

var xxx_messageInfo_RoomData proto.InternalMessageInfo

func (m *RoomData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoomData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RoomData) GetNrOfSeats() int32 {
	if m != nil {
		return m.NrOfSeats
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateRoomMsg)(nil), "CreateRoomMsg")
	proto.RegisterType((*CreateRoomResponseMsg)(nil), "CreateRoomResponseMsg")
	proto.RegisterType((*DeleteRoomMsg)(nil), "DeleteRoomMsg")
	proto.RegisterType((*DeleteRoomResponseMsg)(nil), "DeleteRoomResponseMsg")
	proto.RegisterType((*GetRoomMsg)(nil), "GetRoomMsg")
	proto.RegisterType((*GetRoomResponseMsg)(nil), "GetRoomResponseMsg")
	proto.RegisterType((*GetRoomsMsg)(nil), "GetRoomsMsg")
	proto.RegisterType((*GetRoomsResponseMsg)(nil), "GetRoomsResponseMsg")
	proto.RegisterType((*RoomData)(nil), "RoomData")
}

func init() { proto.RegisterFile("room_service.proto", fileDescriptor_36c4a5b3f0e57253) }

var fileDescriptor_36c4a5b3f0e57253 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xc2, 0x30,
	0x14, 0xc6, 0x5b, 0xff, 0xb6, 0xa7, 0xea, 0xc5, 0x71, 0x4a, 0x28, 0x0e, 0x4b, 0x6e, 0xd6, 0x9b,
	0x85, 0x4d, 0xc1, 0xfb, 0x31, 0x61, 0x37, 0x93, 0x41, 0x7c, 0x00, 0xe9, 0xf4, 0x4c, 0x84, 0x69,
	0xa4, 0xe9, 0xf6, 0xba, 0x7b, 0x95, 0xd1, 0xd6, 0x9a, 0x76, 0xeb, 0x5d, 0xf3, 0xe5, 0xe4, 0xd7,
	0xf0, 0xfb, 0x02, 0x18, 0x2b, 0x75, 0xdc, 0x68, 0x8a, 0xbf, 0x0f, 0x5b, 0x12, 0xe7, 0x58, 0x25,
	0x8a, 0x3f, 0x41, 0xff, 0x39, 0xa6, 0x28, 0x21, 0xa9, 0xd4, 0x71, 0xa5, 0xf7, 0x88, 0xd0, 0x3a,
	0x45, 0x47, 0x62, 0x76, 0x60, 0x87, 0xae, 0xcc, 0xbe, 0x71, 0x02, 0xee, 0x29, 0x7e, 0xfb, 0x58,
	0x53, 0x94, 0x68, 0xd6, 0x08, 0xec, 0xb0, 0x2d, 0x4d, 0xc0, 0xef, 0x60, 0x64, 0x10, 0x92, 0xf4,
	0x59, 0x9d, 0x34, 0xa5, 0xa8, 0x01, 0x34, 0x0e, 0x3b, 0xd6, 0xcc, 0xe6, 0x1b, 0x87, 0x1d, 0x9f,
	0x42, 0x7f, 0x49, 0x9f, 0x64, 0xfe, 0x95, 0x0f, 0xb4, 0xae, 0x03, 0x8f, 0x30, 0x32, 0x03, 0x65,
	0x12, 0x83, 0xae, 0xfe, 0xda, 0x6e, 0x49, 0x6b, 0xd6, 0x0e, 0xec, 0xd0, 0x91, 0xc5, 0x92, 0x4f,
	0x00, 0x5e, 0x28, 0xa9, 0x02, 0x3b, 0x57, 0xe0, 0x1c, 0xf0, 0xb2, 0x5b, 0xa6, 0xdd, 0x42, 0x2b,
	0x35, 0xc1, 0xba, 0x81, 0x1d, 0x7a, 0x33, 0x57, 0xa4, 0xfb, 0xcb, 0x28, 0x89, 0x64, 0x16, 0xf3,
	0x3e, 0x78, 0x97, 0x43, 0x7a, 0xa5, 0xf7, 0x7c, 0x01, 0xc3, 0x62, 0x59, 0x86, 0x4c, 0xa1, 0x9d,
	0x65, 0xcc, 0x09, 0x9a, 0x55, 0x4a, 0x9e, 0xf3, 0x57, 0x70, 0x8a, 0xe8, 0x2a, 0xd5, 0x2d, 0x49,
	0xcd, 0xef, 0x0a, 0xc5, 0x5d, 0xab, 0x92, 0xbd, 0x3f, 0x92, 0x67, 0x3f, 0x36, 0xf4, 0x52, 0xdc,
	0x66, 0x9d, 0xd7, 0x87, 0x0b, 0x00, 0x63, 0x1d, 0x07, 0xa2, 0xd2, 0xa2, 0x3f, 0x16, 0xb5, 0x95,
	0x70, 0x2b, 0x3d, 0x67, 0x1c, 0xe3, 0x40, 0x54, 0x1a, 0xf1, 0xc7, 0xa2, 0xb6, 0x00, 0x6e, 0xe1,
	0x3d, 0x74, 0x2f, 0x1a, 0xd0, 0x13, 0x46, 0xb9, 0x3f, 0x14, 0xff, 0x0d, 0x73, 0x0b, 0x1f, 0xc0,
	0x29, 0xac, 0x61, 0x4f, 0x94, 0x7c, 0xfa, 0x37, 0xa2, 0x46, 0x27, 0xb7, 0xde, 0x3b, 0xd9, 0x83,
	0x9c, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x02, 0xa9, 0xb6, 0x25, 0xa6, 0x02, 0x00, 0x00,
}
