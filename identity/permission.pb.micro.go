// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: identity/permission.proto

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

// Api Endpoints for PermissionService service

func NewPermissionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{
		{
			Name:    "PermissionService.Create",
			Path:    []string{"/identity/permission"},
			Method:  []string{"POST"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.Read",
			Path:    []string{"/identity/permission/{id}"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.Update",
			Path:    []string{"/identity/permission/{id}"},
			Method:  []string{"PUT"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.Delete",
			Path:    []string{"/identity/permission/{id}"},
			Method:  []string{"DELETE"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.List",
			Path:    []string{"/identity/permissions"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.PageResources",
			Path:    []string{"/identity/permission/{id}/page/resources"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.GrantPageResources",
			Path:    []string{"/identity/permission/{id}/page/resource"},
			Method:  []string{"PUT"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.DeletePageResource",
			Path:    []string{"/identity/permission/{permId}/page/{pageId}/resource"},
			Method:  []string{"DELETE"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.ApiResources",
			Path:    []string{"/identity/permission/{id}/api/resources"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.GrantApiResources",
			Path:    []string{"/identity/permission/{id}/api/resource"},
			Method:  []string{"PUT"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.DeleteApiResource",
			Path:    []string{"/identity/permission/{permId}/api/{apiId}/resource"},
			Method:  []string{"DELETE"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.MenuResources",
			Path:    []string{"/identity/permission/{id}/menu/resources"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.GrantMenuResource",
			Path:    []string{"/identity/permission/{permId}/menu/resource"},
			Method:  []string{"PUT"},
			Handler: "rpc",
		},
		{
			Name:    "PermissionService.DeleteMenuResource",
			Path:    []string{"/identity/permission/{permId}/menu/{menuId}/resource"},
			Method:  []string{"DELETE"},
			Handler: "rpc",
		},
	}
}

// Client API for PermissionService service

type PermissionService interface {
	Create(ctx context.Context, in *PermCreateRequest, opts ...client.CallOption) (*PermCreateResponse, error)
	Read(ctx context.Context, in *PermReadRequest, opts ...client.CallOption) (*PermReadResponse, error)
	Update(ctx context.Context, in *PermUpdateRequest, opts ...client.CallOption) (*PermUpdateResponse, error)
	Delete(ctx context.Context, in *PermDeleteRequest, opts ...client.CallOption) (*PermDeleteResponse, error)
	List(ctx context.Context, in *PermListRequest, opts ...client.CallOption) (*PermListResponse, error)
	// the page resources, `readonly` properties are indirect dependent on
	PageResources(ctx context.Context, in *PermPageResourceRequest, opts ...client.CallOption) (*PermPageResourceResponse, error)
	GrantPageResources(ctx context.Context, in *PermGrantPageResourceRequest, opts ...client.CallOption) (*PermGrantPageResourceResponse, error)
	DeletePageResource(ctx context.Context, in *PermDeletePageResourceRequest, opts ...client.CallOption) (*PermDeletePageResourceResponse, error)
	// the api resources, `readonly` properties are indirect dependent on
	ApiResources(ctx context.Context, in *PermApiResourceRequest, opts ...client.CallOption) (*PermApiResourceResponse, error)
	GrantApiResources(ctx context.Context, in *PermGrantApiResourceRequest, opts ...client.CallOption) (*PermGrantApiResourceResponse, error)
	DeleteApiResource(ctx context.Context, in *PermDeleteApiResourceRequest, opts ...client.CallOption) (*PermDeleteApiResourceResponse, error)
	MenuResources(ctx context.Context, in *PermMenuResourceRequest, opts ...client.CallOption) (*PermMenuResourceRequest, error)
	GrantMenuResource(ctx context.Context, in *PermGrantMenuResourceRequest, opts ...client.CallOption) (*PermGrantMenuResourceResponse, error)
	DeleteMenuResource(ctx context.Context, in *PermDeleteMenuResourceRequest, opts ...client.CallOption) (*PermDeleteMenuResourceResponse, error)
}

type permissionService struct {
	c    client.Client
	name string
}

func NewPermissionService(name string, c client.Client) PermissionService {
	return &permissionService{
		c:    c,
		name: name,
	}
}

func (c *permissionService) Create(ctx context.Context, in *PermCreateRequest, opts ...client.CallOption) (*PermCreateResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.Create", in)
	out := new(PermCreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) Read(ctx context.Context, in *PermReadRequest, opts ...client.CallOption) (*PermReadResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.Read", in)
	out := new(PermReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) Update(ctx context.Context, in *PermUpdateRequest, opts ...client.CallOption) (*PermUpdateResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.Update", in)
	out := new(PermUpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) Delete(ctx context.Context, in *PermDeleteRequest, opts ...client.CallOption) (*PermDeleteResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.Delete", in)
	out := new(PermDeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) List(ctx context.Context, in *PermListRequest, opts ...client.CallOption) (*PermListResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.List", in)
	out := new(PermListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) PageResources(ctx context.Context, in *PermPageResourceRequest, opts ...client.CallOption) (*PermPageResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.PageResources", in)
	out := new(PermPageResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) GrantPageResources(ctx context.Context, in *PermGrantPageResourceRequest, opts ...client.CallOption) (*PermGrantPageResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.GrantPageResources", in)
	out := new(PermGrantPageResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) DeletePageResource(ctx context.Context, in *PermDeletePageResourceRequest, opts ...client.CallOption) (*PermDeletePageResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.DeletePageResource", in)
	out := new(PermDeletePageResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) ApiResources(ctx context.Context, in *PermApiResourceRequest, opts ...client.CallOption) (*PermApiResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.ApiResources", in)
	out := new(PermApiResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) GrantApiResources(ctx context.Context, in *PermGrantApiResourceRequest, opts ...client.CallOption) (*PermGrantApiResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.GrantApiResources", in)
	out := new(PermGrantApiResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) DeleteApiResource(ctx context.Context, in *PermDeleteApiResourceRequest, opts ...client.CallOption) (*PermDeleteApiResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.DeleteApiResource", in)
	out := new(PermDeleteApiResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) MenuResources(ctx context.Context, in *PermMenuResourceRequest, opts ...client.CallOption) (*PermMenuResourceRequest, error) {
	req := c.c.NewRequest(c.name, "PermissionService.MenuResources", in)
	out := new(PermMenuResourceRequest)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) GrantMenuResource(ctx context.Context, in *PermGrantMenuResourceRequest, opts ...client.CallOption) (*PermGrantMenuResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.GrantMenuResource", in)
	out := new(PermGrantMenuResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) DeleteMenuResource(ctx context.Context, in *PermDeleteMenuResourceRequest, opts ...client.CallOption) (*PermDeleteMenuResourceResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.DeleteMenuResource", in)
	out := new(PermDeleteMenuResourceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PermissionService service

type PermissionServiceHandler interface {
	Create(context.Context, *PermCreateRequest, *PermCreateResponse) error
	Read(context.Context, *PermReadRequest, *PermReadResponse) error
	Update(context.Context, *PermUpdateRequest, *PermUpdateResponse) error
	Delete(context.Context, *PermDeleteRequest, *PermDeleteResponse) error
	List(context.Context, *PermListRequest, *PermListResponse) error
	// the page resources, `readonly` properties are indirect dependent on
	PageResources(context.Context, *PermPageResourceRequest, *PermPageResourceResponse) error
	GrantPageResources(context.Context, *PermGrantPageResourceRequest, *PermGrantPageResourceResponse) error
	DeletePageResource(context.Context, *PermDeletePageResourceRequest, *PermDeletePageResourceResponse) error
	// the api resources, `readonly` properties are indirect dependent on
	ApiResources(context.Context, *PermApiResourceRequest, *PermApiResourceResponse) error
	GrantApiResources(context.Context, *PermGrantApiResourceRequest, *PermGrantApiResourceResponse) error
	DeleteApiResource(context.Context, *PermDeleteApiResourceRequest, *PermDeleteApiResourceResponse) error
	MenuResources(context.Context, *PermMenuResourceRequest, *PermMenuResourceRequest) error
	GrantMenuResource(context.Context, *PermGrantMenuResourceRequest, *PermGrantMenuResourceResponse) error
	DeleteMenuResource(context.Context, *PermDeleteMenuResourceRequest, *PermDeleteMenuResourceResponse) error
}

func MicroRegisterPermissionServiceHandler(s server.Server, hdlr PermissionServiceHandler, opts ...server.HandlerOption) error {
	type permissionService interface {
		Create(ctx context.Context, in *PermCreateRequest, out *PermCreateResponse) error
		Read(ctx context.Context, in *PermReadRequest, out *PermReadResponse) error
		Update(ctx context.Context, in *PermUpdateRequest, out *PermUpdateResponse) error
		Delete(ctx context.Context, in *PermDeleteRequest, out *PermDeleteResponse) error
		List(ctx context.Context, in *PermListRequest, out *PermListResponse) error
		PageResources(ctx context.Context, in *PermPageResourceRequest, out *PermPageResourceResponse) error
		GrantPageResources(ctx context.Context, in *PermGrantPageResourceRequest, out *PermGrantPageResourceResponse) error
		DeletePageResource(ctx context.Context, in *PermDeletePageResourceRequest, out *PermDeletePageResourceResponse) error
		ApiResources(ctx context.Context, in *PermApiResourceRequest, out *PermApiResourceResponse) error
		GrantApiResources(ctx context.Context, in *PermGrantApiResourceRequest, out *PermGrantApiResourceResponse) error
		DeleteApiResource(ctx context.Context, in *PermDeleteApiResourceRequest, out *PermDeleteApiResourceResponse) error
		MenuResources(ctx context.Context, in *PermMenuResourceRequest, out *PermMenuResourceRequest) error
		GrantMenuResource(ctx context.Context, in *PermGrantMenuResourceRequest, out *PermGrantMenuResourceResponse) error
		DeleteMenuResource(ctx context.Context, in *PermDeleteMenuResourceRequest, out *PermDeleteMenuResourceResponse) error
	}
	type PermissionService struct {
		permissionService
	}
	h := &permissionServiceHandler{hdlr}
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.Create",
		Path:    []string{"/identity/permission"},
		Method:  []string{"POST"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.Read",
		Path:    []string{"/identity/permission/{id}"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.Update",
		Path:    []string{"/identity/permission/{id}"},
		Method:  []string{"PUT"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.Delete",
		Path:    []string{"/identity/permission/{id}"},
		Method:  []string{"DELETE"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.List",
		Path:    []string{"/identity/permissions"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.PageResources",
		Path:    []string{"/identity/permission/{id}/page/resources"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.GrantPageResources",
		Path:    []string{"/identity/permission/{id}/page/resource"},
		Method:  []string{"PUT"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.DeletePageResource",
		Path:    []string{"/identity/permission/{permId}/page/{pageId}/resource"},
		Method:  []string{"DELETE"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.ApiResources",
		Path:    []string{"/identity/permission/{id}/api/resources"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.GrantApiResources",
		Path:    []string{"/identity/permission/{id}/api/resource"},
		Method:  []string{"PUT"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.DeleteApiResource",
		Path:    []string{"/identity/permission/{permId}/api/{apiId}/resource"},
		Method:  []string{"DELETE"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.MenuResources",
		Path:    []string{"/identity/permission/{id}/menu/resources"},
		Method:  []string{"GET"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.GrantMenuResource",
		Path:    []string{"/identity/permission/{permId}/menu/resource"},
		Method:  []string{"PUT"},
		Handler: "rpc",
	}))
	opts = append(opts, api.WithEndpoint(&api.Endpoint{
		Name:    "PermissionService.DeleteMenuResource",
		Path:    []string{"/identity/permission/{permId}/menu/{menuId}/resource"},
		Method:  []string{"DELETE"},
		Handler: "rpc",
	}))
	return s.Handle(s.NewHandler(&PermissionService{h}, opts...))
}

type permissionServiceHandler struct {
	PermissionServiceHandler
}

func (h *permissionServiceHandler) Create(ctx context.Context, in *PermCreateRequest, out *PermCreateResponse) error {
	return h.PermissionServiceHandler.Create(ctx, in, out)
}

func (h *permissionServiceHandler) Read(ctx context.Context, in *PermReadRequest, out *PermReadResponse) error {
	return h.PermissionServiceHandler.Read(ctx, in, out)
}

func (h *permissionServiceHandler) Update(ctx context.Context, in *PermUpdateRequest, out *PermUpdateResponse) error {
	return h.PermissionServiceHandler.Update(ctx, in, out)
}

func (h *permissionServiceHandler) Delete(ctx context.Context, in *PermDeleteRequest, out *PermDeleteResponse) error {
	return h.PermissionServiceHandler.Delete(ctx, in, out)
}

func (h *permissionServiceHandler) List(ctx context.Context, in *PermListRequest, out *PermListResponse) error {
	return h.PermissionServiceHandler.List(ctx, in, out)
}

func (h *permissionServiceHandler) PageResources(ctx context.Context, in *PermPageResourceRequest, out *PermPageResourceResponse) error {
	return h.PermissionServiceHandler.PageResources(ctx, in, out)
}

func (h *permissionServiceHandler) GrantPageResources(ctx context.Context, in *PermGrantPageResourceRequest, out *PermGrantPageResourceResponse) error {
	return h.PermissionServiceHandler.GrantPageResources(ctx, in, out)
}

func (h *permissionServiceHandler) DeletePageResource(ctx context.Context, in *PermDeletePageResourceRequest, out *PermDeletePageResourceResponse) error {
	return h.PermissionServiceHandler.DeletePageResource(ctx, in, out)
}

func (h *permissionServiceHandler) ApiResources(ctx context.Context, in *PermApiResourceRequest, out *PermApiResourceResponse) error {
	return h.PermissionServiceHandler.ApiResources(ctx, in, out)
}

func (h *permissionServiceHandler) GrantApiResources(ctx context.Context, in *PermGrantApiResourceRequest, out *PermGrantApiResourceResponse) error {
	return h.PermissionServiceHandler.GrantApiResources(ctx, in, out)
}

func (h *permissionServiceHandler) DeleteApiResource(ctx context.Context, in *PermDeleteApiResourceRequest, out *PermDeleteApiResourceResponse) error {
	return h.PermissionServiceHandler.DeleteApiResource(ctx, in, out)
}

func (h *permissionServiceHandler) MenuResources(ctx context.Context, in *PermMenuResourceRequest, out *PermMenuResourceRequest) error {
	return h.PermissionServiceHandler.MenuResources(ctx, in, out)
}

func (h *permissionServiceHandler) GrantMenuResource(ctx context.Context, in *PermGrantMenuResourceRequest, out *PermGrantMenuResourceResponse) error {
	return h.PermissionServiceHandler.GrantMenuResource(ctx, in, out)
}

func (h *permissionServiceHandler) DeleteMenuResource(ctx context.Context, in *PermDeleteMenuResourceRequest, out *PermDeleteMenuResourceResponse) error {
	return h.PermissionServiceHandler.DeleteMenuResource(ctx, in, out)
}
