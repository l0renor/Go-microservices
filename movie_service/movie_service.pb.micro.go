// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: movie_service.proto

package movie_service

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

// Client API for Movie_Service service

type Movie_Service interface {
	CreateMovie(ctx context.Context, in *CreateMovieMsg, opts ...client.CallOption) (*CreateMovieResponseMsg, error)
	DeleteMovie(ctx context.Context, in *DeleteMovieMsg, opts ...client.CallOption) (*DeleteMovieResponseMsg, error)
	GetMovie(ctx context.Context, in *GetMovieMsg, opts ...client.CallOption) (*GetMovieResponseMsg, error)
	GetMovies(ctx context.Context, in *GetMoviesMsg, opts ...client.CallOption) (*GetMoviesResponseMsg, error)
}

type movie_Service struct {
	c    client.Client
	name string
}

func NewMovie_Service(name string, c client.Client) Movie_Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "movie_service"
	}
	return &movie_Service{
		c:    c,
		name: name,
	}
}

func (c *movie_Service) CreateMovie(ctx context.Context, in *CreateMovieMsg, opts ...client.CallOption) (*CreateMovieResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Movie_Service.CreateMovie", in)
	out := new(CreateMovieResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movie_Service) DeleteMovie(ctx context.Context, in *DeleteMovieMsg, opts ...client.CallOption) (*DeleteMovieResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Movie_Service.DeleteMovie", in)
	out := new(DeleteMovieResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movie_Service) GetMovie(ctx context.Context, in *GetMovieMsg, opts ...client.CallOption) (*GetMovieResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Movie_Service.GetMovie", in)
	out := new(GetMovieResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movie_Service) GetMovies(ctx context.Context, in *GetMoviesMsg, opts ...client.CallOption) (*GetMoviesResponseMsg, error) {
	req := c.c.NewRequest(c.name, "Movie_Service.GetMovies", in)
	out := new(GetMoviesResponseMsg)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Movie_Service service

type Movie_ServiceHandler interface {
	CreateMovie(context.Context, *CreateMovieMsg, *CreateMovieResponseMsg) error
	DeleteMovie(context.Context, *DeleteMovieMsg, *DeleteMovieResponseMsg) error
	GetMovie(context.Context, *GetMovieMsg, *GetMovieResponseMsg) error
	GetMovies(context.Context, *GetMoviesMsg, *GetMoviesResponseMsg) error
}

func RegisterMovie_ServiceHandler(s server.Server, hdlr Movie_ServiceHandler, opts ...server.HandlerOption) error {
	type movie_Service interface {
		CreateMovie(ctx context.Context, in *CreateMovieMsg, out *CreateMovieResponseMsg) error
		DeleteMovie(ctx context.Context, in *DeleteMovieMsg, out *DeleteMovieResponseMsg) error
		GetMovie(ctx context.Context, in *GetMovieMsg, out *GetMovieResponseMsg) error
		GetMovies(ctx context.Context, in *GetMoviesMsg, out *GetMoviesResponseMsg) error
	}
	type Movie_Service struct {
		movie_Service
	}
	h := &movie_ServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&Movie_Service{h}, opts...))
}

type movie_ServiceHandler struct {
	Movie_ServiceHandler
}

func (h *movie_ServiceHandler) CreateMovie(ctx context.Context, in *CreateMovieMsg, out *CreateMovieResponseMsg) error {
	return h.Movie_ServiceHandler.CreateMovie(ctx, in, out)
}

func (h *movie_ServiceHandler) DeleteMovie(ctx context.Context, in *DeleteMovieMsg, out *DeleteMovieResponseMsg) error {
	return h.Movie_ServiceHandler.DeleteMovie(ctx, in, out)
}

func (h *movie_ServiceHandler) GetMovie(ctx context.Context, in *GetMovieMsg, out *GetMovieResponseMsg) error {
	return h.Movie_ServiceHandler.GetMovie(ctx, in, out)
}

func (h *movie_ServiceHandler) GetMovies(ctx context.Context, in *GetMoviesMsg, out *GetMoviesResponseMsg) error {
	return h.Movie_ServiceHandler.GetMovies(ctx, in, out)
}
