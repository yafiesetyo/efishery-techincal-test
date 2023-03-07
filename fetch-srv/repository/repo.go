package repository

import (
	"encoding/json"
	"fetch-srv/config"
	"fetch-srv/infra"
)

type (
	IRepo interface {
		GetList() ([]List, error)
		GetCurrency() (Currency, error)
	}

	repo struct {
		http infra.HttpClientIface
	}
)

func New(http infra.HttpClientIface) IRepo {
	return &repo{
		http: http,
	}
}

func (r *repo) GetList() (out []List, err error) {
	data, err := r.http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list", nil)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &out); err != nil {
		return out, err
	}

	return
}

func (r *repo) GetCurrency() (out Currency, err error) {
	headers := map[string]string{
		"apikey": config.Cfg.FixerAPIKey,
	}

	data, err := r.http.Get("https://api.apilayer.com/fixer/latest?base=IDR&symbols=USD", &headers)
	if err != nil {
		return
	}

	if err := json.Unmarshal(data, &out); err != nil {
		return out, err
	}

	return
}
