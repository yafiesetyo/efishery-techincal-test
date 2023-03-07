package main

import (
	"auth-srv/config"
	"auth-srv/router"
	"auth-srv/setup"
	"auth-srv/utils/logger"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	config.Init()
	logger.Init(config.Cfg.Env)
}

func main() {
	handlers := setup.New()
	r := router.Init(handlers)

	srv := http.Server{
		Addr:              fmt.Sprintf(":%s", config.Cfg.Port),
		Handler:           r,
		ReadHeaderTimeout: time.Second * 10,
	}

	idleConsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)

		<-sigint

		log.Info("We received an interrupt signal, shut down")
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Errorf("ERROR STOPPING HTTP AUTH_SRV, err: %v", err.Error())
		}
		close(idleConsClosed)
		log.Info("SUCCESS STOPPING HTTP AUTH_SRV")
	}()
	log.Infof("Listening on port %v", config.Cfg.Port)
	log.Info("SUCCESS RUNNING HTTP AUTH_SRV")
	if err := srv.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		log.Errorf("ERROR START HTTP AUTH_SRV, err: %v", err.Error())
	}

	<-idleConsClosed
}
