package usecase

import (
	"auth-srv/config"
	"auth-srv/repository"
	"auth-srv/utils/logger"
	"auth-srv/utils/str"
	"errors"
	"time"

	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type IUsecase interface {
	Login(in User) (token string, err error)
	Create(in User) (pwd string, err error)
}

type usecase struct {
	repo repository.IRepo
}

func New(repo repository.IRepo) IUsecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Create(in User) (pwd string, err error) {
	logCtx := fmt.Sprintf("%T.Create", *u)

	datas, err := u.repo.FindAll()
	if err != nil {
		logger.Error(logCtx, "repo.FindAll error: %v", err)
		return
	}

	for _, data := range datas {
		if data.Phone == in.Phone {
			return "", errors.New("Phone number already registered")
		}
	}

	pwd = str.GeneratePassword()
	in.Password = pwd

	if _, err := u.repo.Create(in.ToEntity()); err != nil {
		logger.Error(logCtx, "repo.Create error: %v", err)
		return "", err
	}

	return
}

func (u *usecase) Login(in User) (token string, err error) {
	logCtx := fmt.Sprintf("%T.Login", *u)

	datas, err := u.repo.FindAll()
	if err != nil {
		logger.Error(logCtx, "repo.FindAll error: %v", err)
		return
	}

	for _, d := range datas {
		if d.Phone == in.Phone && d.Password == in.Password {
			token, err := u.generateToken(d)
			if err != nil {
				logger.Error(logCtx, "u.generateToken error: %v", err)
				return "", err
			}

			return token, nil
		}
	}

	return "", errors.New("Phone or password doesnt match")
}

func (u *usecase) generateToken(in repository.User) (out string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = in.Name
	claims["phone"] = in.Phone
	claims["role"] = in.Role
	claims["created_at"] = in.CreatedAt
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.Cfg.JWT.AccessTokenExp)).Unix()

	t, err := token.SignedString([]byte(config.Cfg.JWT.Secret))
	if err != nil {
		return
	}

	return t, nil
}
