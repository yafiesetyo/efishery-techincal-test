package setup

import (
	httpHandler "fetch-srv/http"
	"fetch-srv/infra"
	"fetch-srv/repository"
	"fetch-srv/usecase"
	"net/http"
)

type Handlers struct {
	Handler httpHandler.IHandler
}

func Init() Handlers {
	httpClient := infra.NewHttpClient(http.Client{})

	repo := repository.New(httpClient)
	uc := usecase.New(repo)
	handler := httpHandler.New(uc)

	return Handlers{
		Handler: handler,
	}
}
