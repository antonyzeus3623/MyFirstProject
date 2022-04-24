package main

import "go.uber.org/zap"

//阶乘
func factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

//菲波那切数
func fabonacci(n int) int {
	if n < 2 {
		return n
	}
	return fabonacci(n-1) + fabonacci(n-2)
}

func main() {
	InitLogger()
	ret := factorial(6)
	zap.S().Debug(ret)

	for i := 0; i < 10; i++ {
		fabonacci(i)
		zap.S().Debugf("斐波那契数列: 第%d个元素，值为%d  ", i, fabonacci(i))
	}
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr", "test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().Panicf("日志初始化失败, error:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
