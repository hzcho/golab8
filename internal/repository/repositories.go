package repository

import (
	"golab8/internal/domain/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repositories struct {
	repository.User
	repository.Account
	repository.Admin
}

func NewRepositories(pool *pgxpool.Pool) *Repositories {
	return &Repositories{
		User:    NewUser(pool),
		Account: NewAccount(pool),
		Admin:   NewAdmin(pool),
	}
}
