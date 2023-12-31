package user

import (
	"context"

	"github.com/letsvote/api/internal/model"
	"github.com/letsvote/api/internal/repository"
)

// The Controller interface provides specification related to order functionality.
type Controller interface {
	Register(ctx context.Context, input model.User) (model.User, error)
	Login(ctx context.Context, input model.User) (model.User, error)
}

type impl struct {
	repo repository.Registry
}

// New returns an implementation instance satisfying controller impl
func New(repo repository.Registry) Controller {
	return impl{
		repo: repo,
	}
}
