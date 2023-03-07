package repository

import (
	"auth-srv/infra/db"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

type IRepo interface {
	Create(in User) (string, error)
	FindAll() (out []User, err error)
	FindOne(id string) (out User, err error)
}

type repo struct {
	db db.IDb
}

func New(db db.IDb) IRepo {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(in User) (string, error) {
	data := map[string]interface{}{}
	id := uuid.New()

	if err := mapstructure.Decode(in, &data); err != nil {
		return "", err
	}

	return id.String(), r.db.Create(id.String(), data)
}

func (r *repo) FindAll() (out []User, err error) {
	datas, err := r.db.ReadAll()
	if err != nil {
		return
	}

	for _, data := range datas {
		if data != "" {
			var user User
			if err := json.Unmarshal([]byte(data), &user); err != nil {
				return out, err
			}

			out = append(out, user)
		}
	}

	return
}

func (r *repo) FindOne(id string) (out User, err error) {
	return out, r.db.ReadOne(id, &out)
}
