package http

import (
	"github.com/gin-gonic/gin"
)

type (
	createResp struct {
		Password string `json:"password"`
	}

	loginResp struct {
		Token string `json:"token"`
	}

	claimsResp struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		CreatedAt string `json:"created_at"`
	}
)

func (c createResp) Populate(password string) createResp {
	return createResp{
		Password: password,
	}
}

func (l loginResp) Populate(token string) loginResp {
	return loginResp{
		Token: token,
	}
}

func (l claimsResp) Populate(c *gin.Context) claimsResp {
	return claimsResp{
		Name:      c.GetString("name"),
		Phone:     c.GetString("phone"),
		Role:      c.GetString("role"),
		CreatedAt: c.GetString("createdAt"),
	}
}
