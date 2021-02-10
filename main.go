package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"testCDN/service"
)

func main() {
	logger, err := getZapLogger()
	if err != nil {
		log.Fatalf("Logger initialization: %v", err)
	}
	s := service.NewService(logger)
	s.Start()
}

func getZapLogger() (*zap.SugaredLogger, error) {
	var conf zap.Config
	var level zapcore.Level
	conf = zap.NewDevelopmentConfig()
	if err := level.Set("debug"); err != nil {
		return nil, err
	}
	zapLogger, err := conf.Build()
	if err != nil {
		return nil, err
	}
	return zapLogger.Sugar(), nil
}
