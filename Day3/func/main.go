package main

import (
	"errors"
	"fmt"
	"log"

	"go.uber.org/zap"
)

//可变参数  (本质上，函数的可变参数是通过切片来实现的。)
func intSum1(x ...int) int {
	fmt.Println(x) //x是一个slice
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}

//固定参数搭配可变参数
func intSum2(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

//多返回值
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//函数作为参数
func intSum3(x, y int) int {
	return x + y

}

func calc2(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//函数作为返回值

func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return intSum3, nil
	case "-":
		return intSub2, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

func intSub2(x, y int) int {
	return x - y
}

func main() {
	InitLogger()

	ret1 := intSum1()
	ret2 := intSum1(10)
	ret3 := intSum1(10, 20)
	ret4 := intSum1(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4)
	zap.S().Info(ret1, ret2, ret3, ret4)
	ret5 := intSum1(10)
	ret6 := intSum1(10, 10)
	ret7 := intSum1(10, 10, 20)
	ret8 := intSum1(10, 10, 20, 30)
	// fmt.Println(ret5, ret6, ret7, ret8)
	zap.S().Info(ret5, ret6, ret7, ret8)
	sum, sub := calc(3, 2)
	// fmt.Println(sum, sub)
	zap.S().Info(sum, sub)
	//函数作为参数
	ret9 := calc2(10, 20, intSum3)
	// fmt.Println(ret9)
	zap.S().Info(ret9)
	// fmt.Printf("the type of intSum3:%T\n", intSum3)
	zap.S().Infof("the type of intSum3:%T", intSum3)

	ret10, err := do("*")
	if err != nil {
		// fmt.Println("输入有误")
		zap.S().Errorf("输入有误, Error: %s", err.Error())
	}
	// fmt.Printf("the type of ret10:%T\n", ret10)
	zap.S().Infof("the type of ret10:%T", ret10)

}

func InitLogger() {
	// config := zap.NewProductionConfig() //创建一个Logger日志记录器（NewProduction()方法）
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr", "log/test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		log.Panicf("日志初始化失败, Error: %s", err.Error())
	}

	zap.ReplaceGlobals(_logger)
}
