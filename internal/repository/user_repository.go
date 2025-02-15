package repository

import (
	"context"

	"github.com/lckrugel/go-basic-api/internal/app/model"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]model.User, error)
	FindById(ctx context.Context, id int32) (model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
	Update(ctx context.Context, user model.User) (model.User, error)
	Delete(ctx context.Context, id int32) error
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByUsername(ctx context.Context, username string) (model.User, error)
}
