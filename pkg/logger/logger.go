package logger

import (
	"github.com/zhf0439/im-server/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global *zap.Logger

func Init(cfg config.LogConfig) error {
	zapCfg := zap.NewDevelopmentConfig()
	zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	level := zapcore.DebugLevel
	if err := level.Set(cfg.Level); err != nil {
		return err
	}
	zapCfg.Level = zap.NewAtomicLevelAt(level)

	l, err := zapCfg.Build()
	if err != nil {
		return err
	}
	global = l
	return nil
}

func L() *zap.Logger {
	if global == nil {
		global = zap.NewNop()
	}
	return global
}

func Sync() {
	if global != nil {
		_ = global.Sync()
	}
}
