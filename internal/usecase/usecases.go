package usecase

import (
	"golab8/internal/config"
	"golab8/internal/domain/usecase"
	"golab8/internal/repository"

	"github.com/sirupsen/logrus"
)

type Usecases struct {
	usecase.User
	usecase.Auth
}

func NewUsecases(repos *repository.Repositories, log *logrus.Logger, cfg *config.Config) *Usecases {
	return &Usecases{
		User: NewUser(repos.User, log),
		Auth: NewAuth(repos.Account, repos.Admin, log, &cfg.Auth),
	}
}
