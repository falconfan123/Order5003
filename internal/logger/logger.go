package logger

import (
	"go.uber.org/zap"
)

var base *zap.Logger
var sugar *zap.SugaredLogger

func Init() error {
    cfg := zap.NewProductionConfig()
    l, err := cfg.Build(zap.AddCaller())
    if err != nil {
        return err
    }
    base = l.WithOptions(zap.AddCallerSkip(1))
    sugar = base.Sugar()
    return nil
}

func Debug(args ...interface{}) { sugar.Debug(args...) }
func Info(args ...interface{})  { sugar.Info(args...) }
func Error(args ...interface{}) { sugar.Error(args...) }
