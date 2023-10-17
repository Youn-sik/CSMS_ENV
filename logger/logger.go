package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLogger *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	zapLogger, _ = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func PrintErrorLogLevel1(err error) {
	zapLogger.Log(6, err.Error())
}

func PrintErrorLogLevel2(err error) {
	zapLogger.Log(7, err.Error())
}

func PrintErrorLogLevel3(err error) {
	zapLogger.Log(8, err.Error())
}

func PrintErrorLogLevel4(err error) {
	zapLogger.Log(9, err.Error())
}

/*
{"level":"ErrorLogLevel1","timestamp":"2023-01-12T15:52:33.928+0900","caller":"KSCONNECT_ADMIN/server.go:1061","msg":"error level1 test"}
{"level":"ErrorLogLevel2","timestamp":"2023-01-12T15:52:33.928+0900","caller":"KSCONNECT_ADMIN/server.go:1062","msg":"error level2 test"}
{"level":"ErrorLogLevel3","timestamp":"2023-01-12T15:52:33.928+0900","caller":"KSCONNECT_ADMIN/server.go:1063","msg":"error level3 test"}
{"level":"ErrorLogLevel4","timestamp":"2023-01-12T15:52:33.928+0900","caller":"KSCONNECT_ADMIN/server.go:1064","msg":"error level4 test"}
*/
