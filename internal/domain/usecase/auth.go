package usecase

import (
	"context"
)

type Auth interface {
	CreateAccount(ctx context.Context, login, password string) (uint64, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
	VerifyToken(ctx context.Context, tkn string) (string, error)
}
