package logger

import "go.uber.org/zap"

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true
	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化错误, err: ", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
