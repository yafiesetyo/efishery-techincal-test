package router

import (
	"auth-srv/middleware"
	"auth-srv/setup"

	"github.com/gin-gonic/gin"
)

func Init(in setup.Handler) *gin.Engine {
	g := gin.Default()

	g.Use(middleware.Logger())
	api := g.Group("/api")
	v1 := api.Group("/v1")

	in.Handler.Mount(v1)

	return g
}
