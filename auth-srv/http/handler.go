package http

import (
	"auth-srv/middleware"
	"auth-srv/usecase"
	"auth-srv/utils/httplib"
	"net/http"

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
	u := g.Group("/users")

	u.POST("/login", h.Login)
	u.POST("", h.Create)

	u.Use(middleware.ValidateToken())
	u.GET("/claims", h.GetClaims)
}

func (h *handler) Create(ctx *gin.Context) {
	var req create

	if err := ctx.ShouldBindJSON(&req); err != nil {
		httplib.WriteResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	pwd, err := h.uc.Create(req.ToEntity())
	if err != nil {
		httplib.WriteResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var resp createResp
	httplib.WriteResponse(ctx, http.StatusCreated, "Created", resp.Populate(pwd))
}

func (h *handler) Login(ctx *gin.Context) {
	var req login

	if err := ctx.ShouldBindJSON(&req); err != nil {
		httplib.WriteResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	token, err := h.uc.Login(req.ToEntity())
	if err != nil {
		httplib.WriteResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	var resp loginResp
	httplib.WriteResponse(ctx, http.StatusOK, "Ok", resp.Populate(token))
}

func (h *handler) GetClaims(ctx *gin.Context) {
	var resp claimsResp
	httplib.WriteResponse(ctx, http.StatusOK, "Ok", resp.Populate(ctx))
}
