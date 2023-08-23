package biz

import (
	"context"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New{{cookiecutter.service_name}}UseCase)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

