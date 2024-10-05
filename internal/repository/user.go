package repository

import (
	"context"
	"fmt"
	"golab8/internal/domain/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	pool *pgxpool.Pool
}

func NewUser(pool *pgxpool.Pool) *User {
	return &User{
		pool: pool,
	}
}

func (u *User) Get(ctx context.Context, filter model.GetUserFilter) ([]model.User, error) {
	query := "SELECT id, name, age FROM users WHERE 1=1"
	var args []interface{}
	argID := 1

	if filter.Name != "" {
		query += fmt.Sprintf(" AND name=$%d", argID)
		args = append(args, filter.Name)
		argID++
	}

	if filter.Age > 0 {
		query += fmt.Sprintf(" AND age=$%d", argID)
		args = append(args, filter.Age)
		argID++
	}

	query += fmt.Sprintf(" LIMIT $%d", argID)
	args = append(args, filter.Limit)
	argID++

	query += fmt.Sprintf(" OFFSET $%d", argID)
	args = append(args, filter.Limit*filter.Page)

	rows, err := u.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) GetById(ctx context.Context, id uint64) (model.User, error) {
	query := "select id, name, age from users where id=$1"

	user := model.User{}
	err := u.pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *User) Add(ctx context.Context, user model.User) (uint64, error) {
	query := "insert into users (name, age) values($1, $2) returning id"

	var id uint64
	err := u.pool.QueryRow(ctx, query, user.Name, user.Age).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *User) Update(ctx context.Context, user model.User) (model.User, error) {
	query := "UPDATE users SET "
	var args []interface{}
	argID := 1

	if user.Name != "" {
		query += fmt.Sprintf("name=$%d,", argID)
		args = append(args, user.Name)
		argID++
	}
	if user.Age >= 0 {
		query += fmt.Sprintf("age=$%d,", argID)
		args = append(args, user.Age)
		argID++
	}

	if len(args) > 0 {
		query = query[:len(query)-1]
	}

	query += fmt.Sprintf(" WHERE id=$%d RETURNING id, name, age", argID)
	args = append(args, user.ID)

	row := u.pool.QueryRow(ctx, query, args...)

	updatedUser := model.User{}
	if err := row.Scan(
		&updatedUser.ID,
		&updatedUser.Name,
		&updatedUser.Age,
	); err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}

func (u *User) Delete(ctx context.Context, id uint64) error {
	query := "DELETE FROM users WHERE id=$1"

	_, err := u.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
