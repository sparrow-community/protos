// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: config/config.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Api Endpoints for Source service

func NewSourceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Source service

type SourceService interface {
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*wrapperspb.BoolValue, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Source_WatchService, error)
}

type sourceService struct {
	c    client.Client
	name string
}

func NewSourceService(name string, c client.Client) SourceService {
	return &sourceService{
		c:    c,
		name: name,
	}
}

func (c *sourceService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Source.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceService) Write(ctx context.Context, in *WriteRequest, opts ...client.CallOption) (*wrapperspb.BoolValue, error) {
	req := c.c.NewRequest(c.name, "Source.Write", in)
	out := new(wrapperspb.BoolValue)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourceService) Watch(ctx context.Context, in *WatchRequest, opts ...client.CallOption) (Source_WatchService, error) {
	req := c.c.NewRequest(c.name, "Source.Watch", &WatchRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &sourceServiceWatch{stream}, nil
}

type Source_WatchService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*WatchResponse, error)
}

type sourceServiceWatch struct {
	stream client.Stream
}

func (x *sourceServiceWatch) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *sourceServiceWatch) Close() error {
	return x.stream.Close()
}

func (x *sourceServiceWatch) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceServiceWatch) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceServiceWatch) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceServiceWatch) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Source service

type SourceHandler interface {
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Write(context.Context, *WriteRequest, *wrapperspb.BoolValue) error
	Watch(context.Context, *WatchRequest, Source_WatchStream) error
}

func RegisterSourceHandler(s server.Server, hdlr SourceHandler, opts ...server.HandlerOption) error {
	type source interface {
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Write(ctx context.Context, in *WriteRequest, out *wrapperspb.BoolValue) error
		Watch(ctx context.Context, stream server.Stream) error
	}
	type Source struct {
		source
	}
	h := &sourceHandler{hdlr}
	return s.Handle(s.NewHandler(&Source{h}, opts...))
}

type sourceHandler struct {
	SourceHandler
}

func (h *sourceHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.SourceHandler.Read(ctx, in, out)
}

func (h *sourceHandler) Write(ctx context.Context, in *WriteRequest, out *wrapperspb.BoolValue) error {
	return h.SourceHandler.Write(ctx, in, out)
}

func (h *sourceHandler) Watch(ctx context.Context, stream server.Stream) error {
	m := new(WatchRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.SourceHandler.Watch(ctx, m, &sourceWatchStream{stream})
}

type Source_WatchStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*WatchResponse) error
}

type sourceWatchStream struct {
	stream server.Stream
}

func (x *sourceWatchStream) Close() error {
	return x.stream.Close()
}

func (x *sourceWatchStream) Context() context.Context {
	return x.stream.Context()
}

func (x *sourceWatchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *sourceWatchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *sourceWatchStream) Send(m *WatchResponse) error {
	return x.stream.Send(m)
}
