package group

import "golab8/internal/usecase"

type Groups struct {
	User
	Auth
	Middleware
}

func NewGroups(usecases *usecase.Usecases) *Groups {
	return &Groups{
		User:       *NewUser(usecases.User),
		Auth:       *NewAuth(usecases.Auth),
		Middleware: *NewMiddleware(usecases.Auth),
	}
}
