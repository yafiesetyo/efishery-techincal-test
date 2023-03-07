package usecase

import (
	"auth-srv/repository"
	"time"
)

type User struct {
	Phone    string
	Name     string
	Role     string
	Password string
}

func (u User) ToEntity() repository.User {
	return repository.User{
		Phone:     u.Phone,
		Name:      u.Name,
		Role:      u.Role,
		Password:  u.Password,
		CreatedAt: time.Now().String(),
	}
}
