package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Initialize() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var err error
	logger, err = config.Build()
	if err != nil {
		return nil, fmt.Errorf("cannot initialize Zap logger")
	}

	return logger, nil
}

func GetInstance() *zap.Logger {
	return logger
}
