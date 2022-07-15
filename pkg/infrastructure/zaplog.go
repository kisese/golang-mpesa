package infrastructure

import (
	"fmt"
	"github.com/kisese/golang_mpesa/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerMgr, _ := config.Build()

	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync()
	logger := loggerMgr.Sugar()

	fmt.Println("Structured Logger established")

	Log = logger
}

/*Example Log

Log.Infow("failed to fetch URL",
		"url", "http://example.com",
		"attempt", http.StatusOK,
		"backoff", time.Second,
	)
*/
