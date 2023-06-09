// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: logger/logger.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Logger service

func NewLoggerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Logger service

type LoggerService interface {
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error)
}

type loggerService struct {
	c    client.Client
	name string
}

func NewLoggerService(name string, c client.Client) LoggerService {
	return &loggerService{
		c:    c,
		name: name,
	}
}

func (c *loggerService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*WriteResponse, error) {
	req := c.c.NewRequest(c.name, "Logger.Write", in)
	out := new(WriteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Logger service

type LoggerHandler interface {
	Write(context.Context, *WriteRequest, *WriteResponse) error
}

func RegisterLoggerHandler(s server.Server, hdlr LoggerHandler, opts ...server.HandlerOption) error {
	type logger interface {
		Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error
	}
	type Logger struct {
		logger
	}
	h := &loggerHandler{hdlr}
	return s.Handle(s.NewHandler(&Logger{h}, opts...))
}

type loggerHandler struct {
	LoggerHandler
}

func (h *loggerHandler) Write(ctx context.Context, in *WriteRequest, out *WriteResponse) error {
	return h.LoggerHandler.Write(ctx, in, out)
}
