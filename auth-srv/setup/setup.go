package setup

import (
	"auth-srv/http"
	"auth-srv/infra/db"
	"auth-srv/repository"
	"auth-srv/usecase"
)

type Handler struct {
	Handler http.IHandler
}

func New() Handler {
	db := db.Init()

	repo := repository.New(db)
	uc := usecase.New(repo)
	handler := http.New(uc)

	return Handler{
		Handler: handler,
	}
}
