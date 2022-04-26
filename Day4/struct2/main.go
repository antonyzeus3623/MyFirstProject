package main

import "go.uber.org/zap"

//结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	zap.S().Debugf("%s会动！", a.name)
}

type Dog struct {
	feet int
	*Animal
}

func (d *Dog) wang() {
	zap.S().Debugf("%s会叫！", d.name)
}

func main() {
	InitLogger()
	d1 := &Dog{
		feet: 8,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d1.wang()
	d1.move()
	// zap.S().Debugf("%#v", d1)

}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化失败，err:", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
