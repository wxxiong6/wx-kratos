package biz

import (
	"context"

	v1 "first/api/{{cookiecutter.api_name}}/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// First is a First model.
type First struct {
	Hello string
}

// FirstRepo is a Greater repo.
type FirstRepo interface {
	Save(context.Context, *First) (*First, error)
	Update(context.Context, *First) (*First, error)
	FindByID(context.Context, int64) (*First, error)
	ListByHello(context.Context, string) ([]*First, error)
	ListAll(context.Context) ([]*First, error)
}

// FirstUsecase is a First usecase.
type FirstUsecase struct {
	repo FirstRepo
	log  *log.Helper
}

// NewFirstUsecase new a First usecase.
func NewFirstUsecase(repo FirstRepo, logger log.Logger) *FirstUsecase {
	return &FirstUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateFirst creates a First, and returns the new First.
func (uc *FirstUsecase) CreateFirst(ctx context.Context, g *First) (*First, error) {
	uc.log.WithContext(ctx).Infof("CreateFirst: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
