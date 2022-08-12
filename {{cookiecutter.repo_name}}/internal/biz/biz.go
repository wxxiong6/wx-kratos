package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(New{{cookiecutter.service_name}}UseCase)
