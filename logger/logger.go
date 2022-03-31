package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Mlogger *zap.Logger

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//path:=c.Request.URL

	}
}

func init() {
	logConfig := zap.NewProductionEncoderConfig()
	logConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logConfig.TimeKey = "Time"

	logConfig.LevelKey = "Level"
	logConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	logConfig.MessageKey = "message"
	logConfig.NameKey = "LoggerName"

	logConfig.EncodeCaller = zapcore.ShortCallerEncoder

	encoder := zapcore.NewConsoleEncoder(logConfig)

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("file open or create failed ,err is %s", err)
		return
	}

	fileCore := zapcore.NewCore(encoder, file, zap.WarnLevel)
	stdoutCore := zapcore.NewCore(encoder, os.Stdout, zap.DebugLevel)
	multiCore := zapcore.NewTee(fileCore, stdoutCore)

	Mlogger = zap.New(multiCore)
	Mlogger = Mlogger.Named("file and stdOut")
	//Mlogger = Mlogger.WithOptions(zap.AddCaller(), zap.OnFatal(zapcore.WriteThenNoop), zap.AddStacktrace(zap.ErrorLevel))
	zap.ReplaceGlobals(Mlogger)
}
