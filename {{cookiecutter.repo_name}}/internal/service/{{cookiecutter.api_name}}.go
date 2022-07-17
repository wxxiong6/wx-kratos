package service

import (
	"context"

	v1 "{{cookiecutter.module_name}}/api/{{cookiecutter.module_name}}/v1"
	"{{cookiecutter.module_name}}/internal/biz"
)

type {{cookiecutter.service_name}}Service struct {
	v1.Unimplemented{{cookiecutter.service_name}}ServiceServer

	uc *biz.{{cookiecutter.service_name}}UseCase
}

func New{{cookiecutter.service_name}}Service(uc *biz.{{cookiecutter.service_name}}UseCase) *{{cookiecutter.service_name}}Service {
	return &{{cookiecutter.service_name}}Service{uc: uc}
}

func (s *{{cookiecutter.service_name}}Service) List(ctx context.Context, in *v1.List{{cookiecutter.service_name}}Request) (*v1.List{{cookiecutter.service_name}}Response, error) {
	rv, err := s.uc.List(ctx, PageSize, PageToken)
	if err != nil {
		return nil, err
	}

	rs := make([]*v1.{{cookiecutter.service_name}}, 0)
	//for _, x := range rv {
	//	rs = append(rs, &v1.First{})
	//	s.uc.Log.Debug(x)
	//}

	return &v1.List{{cookiecutter.service_name}}Response {
		{{cookiecutter.service_name}}: rs,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Create(ctx context.Context, in *v1.Create{{cookiecutter.service_name}}Request) (*v1.{{cookiecutter.service_name}}, error) {
	_, err := s.uc.Create(ctx, &biz.First{})
	if err != nil {
		return nil, err
	}

	return &v1.First {
	//Id: x.Id,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Get(ctx context.Context, in *v1.Get{{cookiecutter.service_name}}Request) (*v1.{{cookiecutter.service_name}}, error) {
	_, err := s.uc.Get(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetOrderReply{
		//Id: x.Id,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Update(ctx context.Context, in *v1.Update{{cookiecutter.service_name}}Request) (*v1.Empty, error) {
	return s.uc.Update(ctx, in)
}

func (s *{{cookiecutter.service_name}}Service) Delete(ctx context.Context, in *v1.Delete{{cookiecutter.service_name}}Request) (*v1.Delete{{cookiecutter.service_name}}Response, error) {
	_, err := s.uc.Create(ctx, &biz.First{})
	if err != nil {
		return err
	}

	return nil
}
