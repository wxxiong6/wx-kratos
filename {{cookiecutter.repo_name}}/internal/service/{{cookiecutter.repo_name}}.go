package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

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

func (s *{{cookiecutter.service_name}}Service) List{{cookiecutter.service_name}}(ctx context.Context, in *v1.List{{cookiecutter.service_name}}Request) (*v1.List{{cookiecutter.service_name}}Response, error) {
    res :=  &v1.List{{cookiecutter.service_name}}Response {}
	_, err := s.uc.List{{cookiecutter.service_name}}(ctx, in.PageSize, in.PageToken)
	if err != nil {
		return res, err
	}

	rs := make([]*v1.{{cookiecutter.service_name}}, 0)
	//for _, x := range rv {
	//	rs = append(rs, &v1.First{})
	//	s.uc.Log.Debug(x)
	//}

	return &v1.List{{cookiecutter.service_name}}Response {
		{{cookiecutter.service_name}}s: rs,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Create{{cookiecutter.service_name}}(ctx context.Context, in *v1.Create{{cookiecutter.service_name}}Request) (*v1.{{cookiecutter.service_name}}, error) {
	res := &v1.{{cookiecutter.service_name}}{}
	_, err := s.uc.Create{{cookiecutter.service_name}}(ctx, &biz.{{cookiecutter.service_name}}{})
	if err != nil {
		return res, err
	}

	return &v1.{{cookiecutter.service_name}} {
	//Id: x.Id,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Get{{cookiecutter.service_name}}(ctx context.Context, in *v1.Get{{cookiecutter.service_name}}Request) (*v1.{{cookiecutter.service_name}}, error) {
	res := &v1.{{cookiecutter.service_name}}{}
	var id int64
	_, err := s.uc.Get{{cookiecutter.service_name}}(ctx, id)
	if err != nil {
		return res, err
	}

	return &v1.{{cookiecutter.service_name}}{
		//Id: x.Id,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Update{{cookiecutter.service_name}}(ctx context.Context, in *v1.Update{{cookiecutter.service_name}}Request) (*v1.{{cookiecutter.service_name}}, error) {
	res := &v1.{{cookiecutter.service_name}}{}

	_, err := s.uc.Update{{cookiecutter.service_name}}(ctx, &biz.{{cookiecutter.service_name}}{}, in.UpdateMask.Paths)
	if err != nil {
		return res, err
	}

	return &v1.{{cookiecutter.service_name}}{
	//Id: x.Id,
	}, nil
}

func (s *{{cookiecutter.service_name}}Service) Delete{{cookiecutter.service_name}}(ctx context.Context, in *v1.Delete{{cookiecutter.service_name}}Request)  (*emptypb.Empty, error) {
	res := &emptypb.Empty{}
	var id int64
	//id ,_ := strconv.ParseInt(in.Id, 10, 64)
	err := s.uc.Delete(ctx, id)
	if err != nil {
		return res, err
	}

	return res, err
}
