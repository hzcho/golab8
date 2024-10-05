package usecase

import (
	"golab8/internal/domain/usecase"
	"golab8/internal/repository"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.User
}

func NewUsecases(repos *repository.Repositories, log *logrus.Logger) *Usecases {
	return &Usecases{
		User: NewUser(repos.User, log),
	}
}
