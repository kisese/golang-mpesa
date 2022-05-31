package infrastructure

import (
	"fmt"
	"github.com/kisese/golang_mpesa/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func InitLogger() {

	outputPaths := []string{"stdout", utils.GetPath() + "/pkg/storage/log/debug.log"}
	fmt.Printf("pwd %v  :", utils.GetPath()+"/pkg/storage/log/debug.log")
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = outputPaths
	config.ErrorOutputPaths = outputPaths
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
