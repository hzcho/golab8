package repository

import (
	"context"
	"golab8/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Account struct {
	pool *pgxpool.Pool
}

func NewAccount(pool *pgxpool.Pool) *Account {
	return &Account{
		pool: pool,
	}
}

func (r *Account) Get(ctx context.Context, login string) (model.Account, error) {
	var account model.Account

	query := "select id, login, pass_hash from accounts where login=$1"

	if err := r.pool.QueryRow(ctx, query, login).Scan(&account.Id, &account.Login, &account.Password); err != nil {
		return model.Account{}, err
	}

	return account, nil
}

func (r *Account) Save(ctx context.Context, email, passHash string) (uint64, error) {
	var userID uint64

	query := "insert into accounts (login, pass_hash) values($1, $2) returning id"

	if err := r.pool.QueryRow(ctx, query, email, passHash).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}
