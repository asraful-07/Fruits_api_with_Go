package user

import (
	"fruits-api/domain"
	userHandler "fruits-api/rest/handlers/user"
)


type Service interface {
	userHandler.Service //embedding
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
