package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

func InitLogger() *zap.SugaredLogger {
	loggerMgr := initZapLog()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr.Sugar()
	return logger
}

var Log = InitLogger()

/*Example Log

Log.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", http.StatusOK,
		"backoff", time.Second,
	)
*/
