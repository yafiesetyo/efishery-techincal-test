package http

import "auth-srv/usecase"

type (
	create struct {
		Name  string `json:"name" binding:"required"`
		Phone string `json:"phone" binding:"required"`
		Role  string `json:"role" binding:"required"`
	}

	login struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
)

func (c create) ToEntity() usecase.User {
	return usecase.User{
		Phone: c.Phone,
		Name:  c.Name,
		Role:  c.Role,
	}
}

func (c login) ToEntity() usecase.User {
	return usecase.User{
		Phone:    c.Phone,
		Password: c.Password,
	}
}
