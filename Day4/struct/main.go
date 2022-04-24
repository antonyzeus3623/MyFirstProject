package main

import (
	"unsafe"

	"go.uber.org/zap"
)

type person struct {
	Name       string
	Age, Score int
}
type number struct {
	a int8
	b int8
	c int8
	d int8
}
type Part1 struct {
	d int
	e int8
	f int16
	g int32
	h int64
	i bool
	j string
	k byte
	l rune
}

type student struct {
	name string
	age  int
}

func main() {
	InitLogger()
	//结构体实例化
	var p1 person
	p1.Name = "小莉"
	p1.Age = 18
	p1.Score = 100
	zap.S().Debug(p1)
	var p2 = new(person)
	zap.S().Debugf("p2的类型为:%T", p2)
	zap.S().Debugf("p2=%v", p2)  //%v 按默认格式输出:p2=&{ 0 0}
	zap.S().Debugf("p2=%+v", p2) //在%v的基础上额外输出字段名:p2=&{Name: Age:0 Score:0}
	zap.S().Debugf("p2=%#v", p2) //在%+v的基础上额外输出类型名:p2=&main.person{Name:"", Age:0, Score:0}
	//Go语言支持对结构体指针直接使用.来直接访问结构体的成员变量字段
	p2.Name = "小王子"
	p2.Age = 20
	p2.Score = 100
	zap.S().Debug(*p2)
	//使用`&`对结构体进行取地址操作相当于对该结构体类型进行了一次`new`实例化操作
	p3 := &person{}
	p3.Name = "小林"
	p3.Age = 22
	p3.Score = 99
	zap.S().Debug(*p3)
	//使用键值对初始化
	p4 := person{
		Name:  "小张",
		Age:   21,
		Score: 98,
	}
	zap.S().Debug(p4)
	//结构体占用一块连续的内存
	zap.S().Debugf("p4.Name的内存地址为：%p", &p4.Name)
	zap.S().Debugf("p4.Agee的内存地址为：%p", &p4.Age)
	zap.S().Debugf("p4.Score的内存地址为：%p", &p4.Score)
	n := number{
		1, 2, 3, 4,
	}
	zap.S().Debugf("n.a的内存地址为：%p", &n.a)
	zap.S().Debugf("n.b的内存地址为：%p", &n.b)
	zap.S().Debugf("n.c的内存地址为：%p", &n.c)
	zap.S().Debugf("n.d的内存地址为：%p", &n.d)
	zap.S().Debugf("int size: %d", unsafe.Sizeof(int(0)))
	zap.S().Debugf("int8 size: %d", unsafe.Sizeof(int8(0)))
	zap.S().Debugf("int16 size: %d", unsafe.Sizeof(int16(0)))
	zap.S().Debugf("int32 size: %d", unsafe.Sizeof(int32(0)))
	zap.S().Debugf("int64 size: %d", unsafe.Sizeof(int64(0)))
	zap.S().Debugf("bool size: %d", unsafe.Sizeof(bool(true)))
	zap.S().Debugf("string size: %d", unsafe.Sizeof(string("")))
	zap.S().Debugf("byte size: %d", unsafe.Sizeof(byte(0)))
	zap.S().Debugf("rune size: %d", unsafe.Sizeof(rune(0)))
	part := Part1{}
	zap.S().Debugf("bool size: %d,align:%d", unsafe.Sizeof(part), unsafe.Alignof(part))
	zap.S().Debugf("")

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "古力娜扎", age: 17},
		{name: "刘亦菲", age: 18},
	}
	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		zap.S().Debug(k, "=>", v.name)
	}
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr", "test.log"} //输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
	config.DisableCaller = true                         //禁止使用调用函数的文件名和行号来注释日志。默认进行注释日志
	config.DisableStacktrace = true                     //是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪

	_logger, err := config.Build()
	if err != nil {
		zap.S().Panicf("日志初始化失败, error:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
