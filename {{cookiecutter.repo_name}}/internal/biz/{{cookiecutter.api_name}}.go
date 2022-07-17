package biz

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.module_name}}/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type {{cookiecutter.service_name}} struct {
}

type {{cookiecutter.service_name}}Repo interface {
	List(context.Context, int32, string) ([]*{{cookiecutter.service_name}}, error)
	Get(context.Context, int64) (*{{cookiecutter.service_name}}, error)
	Create(context.Context, *{{cookiecutter.service_name}}) (*{{cookiecutter.service_name}}, error)
	Update(context.Context, *{{cookiecutter.service_name}}, []string) (*{{cookiecutter.service_name}}, error)
	Delete(context.Context, int64) error
}

type {{cookiecutter.service_name}}UseCase struct {
	repo {{cookiecutter.service_name}}Repo
	log  *log.Helper
}


func (uc *{{cookiecutter.service_name}}UseCase) List(ctx context.Context, pageSize int32, pageToken string) ([]*{{cookiecutter.service_name}}, error) {
	return uc.repo.List(ctx, pageSize, pageToken)
}

func New{{cookiecutter.service_name}}UseCase(repo {{cookiecutter.service_name}}Repo, logger log.Logger) *{{cookiecutter.service_name}}UseCase {
	return &{{cookiecutter.service_name}}UseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/{{cookiecutter.service_name}}"))}
}

func (uc *{{cookiecutter.service_name}}UseCase) Create(ctx context.Context, g *{{cookiecutter.service_name}}) (*{{cookiecutter.service_name}}, error) {
	return uc.repo.Create(ctx, g)
}

func (uc *{{cookiecutter.service_name}}UseCase) Update(ctx context.Context, g *{{cookiecutter.service_name}}, fm []string) (*{{cookiecutter.service_name}}, error) {
	return uc.repo.Update(ctx, g, fm)
}

func (uc *{{cookiecutter.service_name}}UseCase) Get(ctx context.Context, id int64) (*{{cookiecutter.service_name}}, error) {
	return uc.repo.Get(ctx, id)
}

func (uc *{{cookiecutter.service_name}}UseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}


