package http

import (
	"fetch-srv/middleware"
	"fetch-srv/usecase"
	"fetch-srv/utils/httplib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	IHandler interface {
		Mount(g *gin.RouterGroup)
	}

	handler struct {
		uc usecase.IUsecase
	}
)

func New(uc usecase.IUsecase) IHandler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) Mount(g *gin.RouterGroup) {
	list := g.Group("/list")

	list.Use(middleware.ValidateToken())
	list.GET("", h.GetList)
	list.GET("/aggregate", h.GetAggregate)
}

func (h *handler) GetAggregate(c *gin.Context) {
	if strings.ToLower(c.GetString("role")) != "admin" {
		httplib.WriteResponse(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	res, err := h.uc.FetchAggregate()
	if err != nil {
		httplib.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	httplib.WriteResponse(c, http.StatusOK, "Ok", res)
}

func (h *handler) GetList(c *gin.Context) {
	res, err := h.uc.Fetch()
	if err != nil {
		httplib.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var resp []List
	for _, r := range res {
		var l List
		resp = append(resp, l.ToResponse(r))
	}
	httplib.WriteResponse(c, http.StatusOK, "Ok", resp)
}
