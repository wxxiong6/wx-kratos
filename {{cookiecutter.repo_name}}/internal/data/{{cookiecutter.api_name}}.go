package data

import (
	"context"

	"{{cookiecutter.module_name}}/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type {{cookiecutter.repo_name}}Repo struct {
	data *Data
	log  *log.Helper
}

func New{{cookiecutter.service_name}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.service_name}}Repo {
	return &{{cookiecutter.repo_name}}Repo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "{{cookiecutter.service_name}}-service/data")),
	}
}

func (r *{{cookiecutter.repo_name}}Repo) List(ctx context.Context, pageSize int32, pageToken string) ([]*biz.{{cookiecutter.service_name}}, error) {
	return nil, nil
}

func (r *{{cookiecutter.repo_name}}Repo) Create(ctx context.Context, g *biz.{{cookiecutter.service_name}}) (*biz.{{cookiecutter.service_name}}, error) {
	return g, nil
}

func (r *{{cookiecutter.repo_name}}Repo) Update(ctx context.Context, g *biz.{{cookiecutter.service_name}}, fm []string) (*biz.{{cookiecutter.service_name}}, error) {
	return g, nil
}

func (r *{{cookiecutter.repo_name}}Repo) Get(ctx context.Context, id int64) (*biz.{{cookiecutter.service_name}}, error) {
	return nil, nil
}

func (r *{{cookiecutter.repo_name}}Repo) Delete(ctx context.Context, id int64) error {
return nil, nil
}

