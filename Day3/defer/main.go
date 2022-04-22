package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"
)

//defer执行时机
//第一步：返回值 赋值
//defer
//第二步：真正的RET返回
//函数中如果存在defer，那么defer执行的时机在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x //1.返回值赋值 2.defer 3.真的的RET指令
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //1. 返回值赋值x=5  2.执行defer：x++,x=6  3. return x=6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x // 1.返回值y = x = 5 2.defer修改x的值 3.真正的返回 y=5
}

func f4() (x int) {
	defer func(x int) { //把返回值x当做一个参数传进函数中
		x++ //改变的是函数中x的副本
	}(x)
	return 5 //返回值 = x = 5
}

func calc(index string, a, b int) int {
	ret := a + b
	// zap.S().Debug(index, a, b, ret)
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	InitLogger()
	zap.S().Debug(f1())
	zap.S().Debug(f2())
	zap.S().Debug(f3())
	zap.S().Debug(f4())
	zap.S().Info(f1())
	zap.S().Fatal(f1())

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	//分析过程：
	//1.defer calc("AA", x, calc("A", x, y))
	//2.calc("A", x, y) // A 1 2 3
	//3.defer calc("AA", x, 3)
	//4.x = 10
	//5.defer calc("BB", x, calc("B", x, y))
	//6.calc("B", x, y) //B 10 2 12
	//7.defer calc("BB", x, 12)
	//8.y = 20
	//9.calc("BB", x, 12) //BB 10 12 22
	//10.defer calc("AA", x, 3) //AA 10 3 13

}
func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr", "log/test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		log.Panicf("日志初始化失败，Error:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
