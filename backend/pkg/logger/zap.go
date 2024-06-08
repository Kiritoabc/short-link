package logger

import (
	"go.uber.org/zap"
	"sync"
)

var logger *zap.Logger

var once sync.Once

func Init() {
	once.Do(initLogger)
}

// init logger
func initLogger() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
}

// GetLogger
// 性能模式日志 例
//
//	logger.Info("failed to fetch URL",
//	// Structured context as strongly typed Field values.
//	zap.String("url", url),
//	zap.Int("attempt", 3),
//	zap.Duration("backoff", time.Second),
//
// )
func GetLogger() *zap.Logger {
	return logger
}

// GetSugared
// printf 风格模式 例
// sugar.Infow("failed to fetch URL",
//
//	// Structured context as loosely typed key-value pairs.
//	"url", url,
//	"attempt", 3,
//	"backoff", time.Second,
//
// )
func GetSugared() *zap.SugaredLogger {
	return logger.Sugar()
}
