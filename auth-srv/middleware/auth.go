package middleware

import (
	"auth-srv/config"
	"auth-srv/utils/httplib"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var (
	errAccessTokenEmpty    = errors.New("token empty")
	errInvalidToken        = errors.New("invalid token")
	errInvalidRefreshToken = errors.New("invalid refresh token")
	errRefreshTokenEmpty   = errors.New("refresh token empty")
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if strings.TrimSpace(authorization) == "" {
			httplib.WriteResponse(ctx, http.StatusUnauthorized, errAccessTokenEmpty.Error(), nil)
			ctx.Abort()
			return
		}

		strBearer := strings.Split(authorization, " ")
		if len(strBearer) <= 1 {
			httplib.WriteResponse(ctx, http.StatusUnauthorized, errAccessTokenEmpty.Error(), nil)
			ctx.Abort()
			return
		}

		tokenStr := strBearer[1]
		if strings.TrimSpace(tokenStr) == "" {
			httplib.WriteResponse(ctx, http.StatusUnauthorized, errAccessTokenEmpty.Error(), nil)
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errInvalidToken
			}
			return []byte(config.Cfg.JWT.Secret), nil
		})
		if err != nil {
			httplib.WriteResponse(ctx, http.StatusUnauthorized, err.Error(), nil)
			ctx.Abort()
			return
		}
		if !token.Valid {
			httplib.WriteResponse(ctx, http.StatusUnauthorized, errInvalidToken.Error(), nil)
			ctx.Abort()
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			httplib.WriteResponse(ctx, http.StatusUnauthorized, errInvalidToken.Error(), nil)
			ctx.Abort()
			return
		}

		ctx.Set("name", payload["name"])
		ctx.Set("phone", payload["phone"])
		ctx.Set("role", payload["role"])
		ctx.Set("createdAt", payload["created_at"])

		ctx.Next()
	}
}
