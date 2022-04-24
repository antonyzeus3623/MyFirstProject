package main

import "go.uber.org/zap"

//任意类型添加方法
//在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。
type MyInt int //MyInt 将int定义为自定义MyInt类型

func (m MyInt) sayHello() {
	zap.S().Debug("Hello,我是一个int")
}

func main() {
	InitLogger()
	var m MyInt
	m.sayHello()
	m = 100
	zap.S().Debugf("%T  %#v", m, m)
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化失败，error:%s", err.Error)
	}
	zap.ReplaceGlobals(_logger)
}
