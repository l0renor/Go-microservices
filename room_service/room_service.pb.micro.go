// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: room_service.proto

package room_service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Room_Service service

type Room_Service interface {
	CreateRoom(ctx context.Context, in *CreateRoomMsg, opts ...client.CallOption) (*CreateRoomResponseMsg, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomMsg, opts ...client.CallOption) (*DeleteRoomResponseMsg, error)
	GetRoom(ctx context.Context, in *GetRoomMsg, opts ...client.CallOption) (*GetRoomResponseMsg, error)
	GetRooms(ctx context.Context, in *GetRoomsMsg, opts ...client.CallOption) (*GetRoomsResponseMsg, error)
}

type room_Service struct {
	c    client.Client
	name string
}

func NewRoom_Service(name string, c client.Client) Room_Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "room_service"
	}
	return &room_Service{
		c:    c,
		name: name,
	}
}

func (c *room_Service) CreateRoom(ctx context.Context, in *CreateRoomMsg, opts ...client.CallOption) (*CreateRoomResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Room_Service.CreateRoom", in)
	out := new(CreateRoomResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *room_Service) DeleteRoom(ctx context.Context, in *DeleteRoomMsg, opts ...client.CallOption) (*DeleteRoomResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Room_Service.DeleteRoom", in)
	out := new(DeleteRoomResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *room_Service) GetRoom(ctx context.Context, in *GetRoomMsg, opts ...client.CallOption) (*GetRoomResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Room_Service.GetRoom", in)
	out := new(GetRoomResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *room_Service) GetRooms(ctx context.Context, in *GetRoomsMsg, opts ...client.CallOption) (*GetRoomsResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Room_Service.GetRooms", in)
	out := new(GetRoomsResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Room_Service service

type Room_ServiceHandler interface {
	CreateRoom(context.Context, *CreateRoomMsg, *CreateRoomResponseMsg) error
	DeleteRoom(context.Context, *DeleteRoomMsg, *DeleteRoomResponseMsg) error
	GetRoom(context.Context, *GetRoomMsg, *GetRoomResponseMsg) error
	GetRooms(context.Context, *GetRoomsMsg, *GetRoomsResponseMsg) error
}

func RegisterRoom_ServiceHandler(s server.Server, hdlr Room_ServiceHandler, opts ...server.HandlerOption) error {
	type room_Service interface {
		CreateRoom(ctx context.Context, in *CreateRoomMsg, out *CreateRoomResponseMsg) error
		DeleteRoom(ctx context.Context, in *DeleteRoomMsg, out *DeleteRoomResponseMsg) error
		GetRoom(ctx context.Context, in *GetRoomMsg, out *GetRoomResponseMsg) error
		GetRooms(ctx context.Context, in *GetRoomsMsg, out *GetRoomsResponseMsg) error
	}
	type Room_Service struct {
		room_Service
	}
	h := &room_ServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&Room_Service{h}, opts...))
}

type room_ServiceHandler struct {
	Room_ServiceHandler
}

func (h *room_ServiceHandler) CreateRoom(ctx context.Context, in *CreateRoomMsg, out *CreateRoomResponseMsg) error {
	return h.Room_ServiceHandler.CreateRoom(ctx, in, out)
}

func (h *room_ServiceHandler) DeleteRoom(ctx context.Context, in *DeleteRoomMsg, out *DeleteRoomResponseMsg) error {
	return h.Room_ServiceHandler.DeleteRoom(ctx, in, out)
}

func (h *room_ServiceHandler) GetRoom(ctx context.Context, in *GetRoomMsg, out *GetRoomResponseMsg) error {
	return h.Room_ServiceHandler.GetRoom(ctx, in, out)
}

func (h *room_ServiceHandler) GetRooms(ctx context.Context, in *GetRoomsMsg, out *GetRoomsResponseMsg) error {
	return h.Room_ServiceHandler.GetRooms(ctx, in, out)
}
