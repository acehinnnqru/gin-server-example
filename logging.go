package main

import (
	"github.com/acehinnnqru/gin-server-example/config"
	log "github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLogger() {
	lc := config.AppConfig.Logging

	lvl, err := log.ParseLevel(lc.Level)
	if err != nil {
		log.SetLevel(log.InfoLevel)
	}
	log.SetLevel(lvl)

	log.SetFormatter(&log.JSONFormatter{})

	// rotate log file
	logRotation := &lumberjack.Logger{
		Filename:   lc.Filename,
		MaxSize:    lc.MaxSize,
		MaxBackups: lc.MaxBackups,
		MaxAge:     lc.MaxAge,
		Compress:   lc.Compress,
	}

	log.SetOutput(logRotation)
	log.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		log.PanicLevel,
		log.FatalLevel,
		log.ErrorLevel,
		log.WarnLevel,
		log.InfoLevel,
	)))
}
