// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: identity/user.proto

package identity

import (
	fmt "fmt"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "UserService.SignUp",
			Path:    []string{"/identity/user:signUp"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "UserService.SignIn",
			Path:    []string{"/identity/user:signIn"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
	}
}

// Client API for UserService service

type UserService interface {
	SignUp(ctx context.Context, in *UserSignUpRequest, opts ...client.CallOption) (*UserSignUpResponse, error)
	SignIn(ctx context.Context, in *UserSignInRequest, opts ...client.CallOption) (*UserSignInResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SignUp(ctx context.Context, in *UserSignUpRequest, opts ...client.CallOption) (*UserSignUpResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.SignUp", in)
	out := new(UserSignUpResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SignIn(ctx context.Context, in *UserSignInRequest, opts ...client.CallOption) (*UserSignInResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.SignIn", in)
	out := new(UserSignInResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	SignUp(context.Context, *UserSignUpRequest, *UserSignUpResponse) error
	SignIn(context.Context, *UserSignInRequest, *UserSignInResponse) error
}

func MicroRegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		SignUp(ctx context.Context, in *UserSignUpRequest, out *UserSignUpResponse) error
		SignIn(ctx context.Context, in *UserSignInRequest, out *UserSignInResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UserService.SignUp",
		Path:    []string{"/identity/user:signUp"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "UserService.SignIn",
		Path:    []string{"/identity/user:signIn"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) SignUp(ctx context.Context, in *UserSignUpRequest, out *UserSignUpResponse) error {
	return h.UserServiceHandler.SignUp(ctx, in, out)
}

func (h *userServiceHandler) SignIn(ctx context.Context, in *UserSignInRequest, out *UserSignInResponse) error {
	return h.UserServiceHandler.SignIn(ctx, in, out)
}