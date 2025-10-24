package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func Init() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	Logger = l
}
