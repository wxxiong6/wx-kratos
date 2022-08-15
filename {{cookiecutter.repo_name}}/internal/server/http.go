package server

import (
	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.api_name}}/v1"
	"{{cookiecutter.module_name}}/internal/conf"
	"{{cookiecutter.module_name}}/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

func NewHTTPServer(c *conf.Server, {{cookiecutter.repo_name}} *service.{{cookiecutter.service_name}}Service, logger log.Logger) *http.Server {
	openAPIHandler := openapiv2.NewHandler()
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/q/", openAPIHandler)
	v1.Register{{cookiecutter.service_name}}ServiceHTTPServer(srv, {{cookiecutter.repo_name}})
	return srv
}
