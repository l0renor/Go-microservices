// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/reservation_service.proto

package api

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

// Client API for Reservation_Service service

type Reservation_Service interface {
}

type reservation_Service struct {
	c    client.Client
	name string
}

func NewReservation_Service(name string, c client.Client) Reservation_Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "reservation_service"
	}
	return &reservation_Service{
		c:    c,
		name: name,
	}
}

// Server API for Reservation_Service service

type Reservation_ServiceHandler interface {
}

func RegisterReservation_ServiceHandler(s server.Server, hdlr Reservation_ServiceHandler, opts ...server.HandlerOption) error {
	type reservation_Service interface {
	}
	type Reservation_Service struct {
		reservation_Service
	}
	h := &reservation_ServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&Reservation_Service{h}, opts...))
}

type reservation_ServiceHandler struct {
	Reservation_ServiceHandler
}