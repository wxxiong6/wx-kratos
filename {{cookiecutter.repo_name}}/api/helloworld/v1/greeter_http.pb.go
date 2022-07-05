// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type {{cookiecutter.service_name}}HTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func Register{{cookiecutter.service_name}}HTTPServer(s *http.Server, srv {{cookiecutter.service_name}}HTTPServer) {
	r := s.Route("/")
	r.GET("/helloworld/{name}", _{{cookiecutter.service_name}}_SayHello0_HTTP_Handler(srv))
}

func _{{cookiecutter.service_name}}_SayHello0_HTTP_Handler(srv {{cookiecutter.service_name}}HTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HelloRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/helloworld.v1.{{cookiecutter.service_name}}/SayHello")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayHello(ctx, req.(*HelloRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HelloReply)
		return ctx.Result(200, reply)
	}
}

type {{cookiecutter.service_name}}HTTPClient interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...http.CallOption) (rsp *HelloReply, err error)
}

type {{cookiecutter.service_name}}HTTPClientImpl struct {
	cc *http.Client
}

func New{{cookiecutter.service_name}}HTTPClient(client *http.Client) {{cookiecutter.service_name}}HTTPClient {
	return &{{cookiecutter.service_name}}HTTPClientImpl{client}
}

func (c *{{cookiecutter.service_name}}HTTPClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http.CallOption) (*HelloReply, error) {
	var out HelloReply
	pattern := "/helloworld/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/helloworld.v1.{{cookiecutter.service_name}}/SayHello"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
