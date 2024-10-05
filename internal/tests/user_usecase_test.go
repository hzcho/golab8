package usecase

import (
	"context"
	"golab8/internal/domain/model"
	mocks "golab8/internal/domain/repository/mock"
	"golab8/internal/usecase"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	userRepoMock := mocks.NewUser(t)
	userUsecase := usecase.NewUser(userRepoMock, logrus.New())

	ctx := context.Background()
	filter := model.GetUserFilter{
		Name:  "test",
		Age:   -67,
		Page:  0,
		Limit: 99,
	}

	userRepoMock.On("Get", ctx, filter).
		Once().
		Return(nil, nil)

	_, err := userUsecase.Get(ctx, filter)

	assert.NoError(t, err)
}

func TestGetById(t *testing.T) {
	userRepoMock := mocks.NewUser(t)
	userUsecase := usecase.NewUser(userRepoMock, logrus.New())

	ctx := context.Background()
	var id uint64

	userRepoMock.On("GetById", ctx, id).
		Once().
		Return(model.User{}, nil)

	_, err := userUsecase.GetById(ctx, id)

	assert.NoError(t, err)
}

func TestAdd(t *testing.T) {
	userRepoMock := mocks.NewUser(t)
	userUsecase := usecase.NewUser(userRepoMock, logrus.New())

	ctx := context.Background()
	addUser := model.AddUser{
		Name: "poperdolilo",
		Age:  18,
	}
	user := model.User{
		Name: addUser.Name,
		Age:  addUser.Age,
	}

	userRepoMock.On("Add", ctx, user).
		Once().
		Return(uint64(0), nil)

	_, err := userUsecase.Add(ctx, addUser)

	assert.NoError(t, err)
}
