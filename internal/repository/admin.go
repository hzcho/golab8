package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Admin struct {
	pool *pgxpool.Pool
}

func NewAdmin(pool *pgxpool.Pool) *Admin {
	return &Admin{
		pool: pool,
	}
}

func (a *Admin) Get(id uint64) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM admins WHERE account_id = $1)`

	err := a.pool.QueryRow(context.Background(), query, id).Scan(&exists)
	if err != nil {
		log.Printf("Error checking admin existence: %v", err)
		return false
	}

	return exists
}
