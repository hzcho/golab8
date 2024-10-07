package usecase

import (
	"context"
	"errors"
	"golab8/internal/config"
	"golab8/internal/domain/model"
	"golab8/internal/domain/repository"
	"golab8/internal/token"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	accountRepo repository.Account
	log         *logrus.Logger
	ttl         time.Duration
	tokenKey    string
}

func NewAuth(accountRepo repository.Account, log *logrus.Logger, cfg *config.Auth) *Auth {
	return &Auth{
		accountRepo: accountRepo,
		log:         log,
		ttl:         cfg.TTL,
		tokenKey:    cfg.TokenKey,
	}
}

func (a *Auth) CreateAccount(ctx context.Context, email, password string) (uint64, error) {
	log := a.log.WithField("op", "internal/usecase/auth/CreateAccount")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	userId, err := a.accountRepo.Save(ctx, email, string(passHash))
	if err != nil {
		log.Error(err.Error())
		return 0, err
	}

	return userId, nil
}

func (a *Auth) GenerateToken(ctx context.Context, login, password string) (string, error) {
	log := a.log.WithField("op", "internal/usecase/auth/GenerateToken")

	user, err := a.accountRepo.Get(ctx, login)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}
	if user == (model.Account{}) {
		err := errors.New("the user does not exist")

		log.Error(err.Error())

		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token, err := token.New(a.tokenKey, token.TokenClaims{
		Login: login,
		Exp:   time.Now().Add(a.ttl).Unix(),
	})

	return token, err
}

func (a *Auth) VerifyToken(ctx context.Context, tkn string) (string, error) {
	log := a.log.WithField("op", "internal/usecase/auth/VerifyToken")

	claims, err := token.ExtractClaims(a.tokenKey, tkn)
	if err != nil {
		log.Error(err.Error())
		return "", err
	}

	return claims.Login, nil
}
