package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Init(env string) {
	log.SetFormatter(&log.JSONFormatter{})
	switch strings.ToLower(env) {
	case "dev":
		log.SetLevel(log.DebugLevel)
	case "prod":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.SetOutput(os.Stdout)
}

func getEntry(ctxName string) *log.Entry {
	return log.WithFields(log.Fields{
		"context": ctxName,
	})
}

func Info(ctxName, format string, args ...interface{}) {
	getEntry(ctxName).Infof(format, args)
}

func Warn(ctxName, format string, args ...interface{}) {
	getEntry(ctxName).Warnf(format, args)
}

func Debug(ctxName, format string, args ...interface{}) {
	getEntry(ctxName).Debugf(format, args)
}

func Error(ctxName, format string, args ...interface{}) {
	getEntry(ctxName).Errorf(format, args)
}
