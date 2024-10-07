package repository

import (
	"context"
	"golab8/internal/domain/model"
)

type Account interface {
	Get(ctx context.Context, login string) (model.Account, error)
	Save(ctx context.Context, login, passHash string) (uint64, error)
}
