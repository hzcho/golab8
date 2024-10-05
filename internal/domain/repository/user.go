package repository

import (
	"context"
	"golab8/internal/domain/model"
)

type User interface {
	Get(ctx context.Context, filter model.GetUserFilter) ([]model.User, error)
	GetById(ctx context.Context, id uint64) (model.User, error)
	Add(ctx context.Context, user model.User) (uint64, error)
	Update(ctx context.Context, user model.User) (model.User, error)
	Delete(ctx context.Context, id uint64) error
}
