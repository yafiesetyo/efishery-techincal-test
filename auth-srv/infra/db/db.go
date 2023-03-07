package db

import (
	"auth-srv/utils/logger"
	"fmt"
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
)

type IDb interface {
	Create(id string, data map[string]interface{}) error
	ReadAll() (out []string, err error)
	ReadOne(id string, out interface{}) (err error)
	Delete(id string) error
}

type db struct {
	drv *scribble.Driver
}

func Init() IDb {
	drv, err := scribble.New("./jsondb", nil)
	if err != nil {
		log.Fatal("failed to init database, error", err)
	}
	return &db{
		drv: drv,
	}
}

func (d *db) Create(id string, data map[string]interface{}) error {
	logCtx := fmt.Sprintf("%T.Create", *d)

	if err := d.drv.Write("user", id, data); err != nil {
		logger.Error(logCtx, "scribble.Write error: %v", err)
		return err
	}

	return nil
}

func (d *db) ReadAll() (out []string, err error) {
	logCtx := fmt.Sprintf("%T.ReadAll", *d)

	out, err = d.drv.ReadAll("user")
	if err != nil {
		logger.Error(logCtx, "scribble.ReadAll error: %v", err)
	}

	return
}

func (d *db) ReadOne(id string, out interface{}) (err error) {
	logCtx := fmt.Sprintf("%T.ReadOne", *d)

	if err = d.drv.Read("user", id, &out); err != nil {
		logger.Error(logCtx, "scribble.ReadOne error: %v", err)
	}

	return
}

func (d *db) Delete(id string) error {
	logCtx := fmt.Sprintf("%T.Delete", *d)

	if err := d.drv.Delete("user", id); err != nil {
		logger.Error(logCtx, "scribble.Delete error: %v", err)
		return err
	}

	return nil
}
