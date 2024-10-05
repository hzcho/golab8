package usecase

import (
	"context"
	"golab8/internal/domain/model"
	"golab8/internal/domain/repository"

	"github.com/sirupsen/logrus"
)

type User struct {
	userRepo repository.User
	log      *logrus.Logger
}

func NewUser(userRepo repository.User, log *logrus.Logger) *User {
	return &User{
		userRepo: userRepo,
		log:      log,
	}
}

func (u *User) Get(ctx context.Context, filter model.GetUserFilter) ([]model.User, error) {
	log := u.log.WithField("op", "internal/usecase/user/Get")

	if filter.Page < 0 {
		filter.Page = 0
	}
	if filter.Limit <= 0 {
		filter.Limit = 10
	}

	users, err := u.userRepo.Get(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return users, nil
}

func (u *User) GetById(ctx context.Context, id uint64) (model.User, error) {
	log := u.log.WithField("op", "internal/usecase/user/GetById")

	user, err := u.userRepo.GetById(ctx, id)
	if err != nil {
		log.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (u *User) Add(ctx context.Context, addUser model.AddUser) (uint64, error) {
	log := u.log.WithField("op", "internal/usecase/user/Add")

	user := model.User{
		Name: addUser.Name,
		Age:  addUser.Age,
	}

	id, err := u.userRepo.Add(ctx, user)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return id, nil
}

func (u *User) Update(ctx context.Context, updateUser model.UpdateUser) (model.User, error) {
	log := u.log.WithField("op", "internal/usecase/user/Update")

	user := model.User{
		ID:   updateUser.ID,
		Name: updateUser.Name,
		Age:  updateUser.Age,
	}

	user, err := u.userRepo.Update(ctx, user)
	if err != nil {
		log.Error(err)
		return model.User{}, err
	}

	return user, nil
}

func (u *User) Delete(ctx context.Context, id uint64) (bool, error) {
	log := u.log.WithField("op", "internal/usecase/user/Delete")

	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, err
}
