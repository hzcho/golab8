package group

import "golab8/internal/domain/usecase"

type Groups struct {
	User
}

func NewGroups(userUseCase usecase.User) *Groups {
	return &Groups{
		User: *NewUser(userUseCase),
	}
}
