package main

import "go.uber.org/zap"

//标识符：变量名 函数名 类型名 方法名
//Go语言中如果标识符首字母是大写的，就表示对外部包可见（暴露的、公有的）

// Person 这是一个人的结构体
type Person struct {
	name string
	age  int8
}

//newPerson构造函数 //Go语言的构造函数约定成俗：以new开头
//当结构体比较大的时候，尽量使用结构体指针，减小程序的内存开销
func newPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//方法是作用于特定类型的函数
//接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
//Dream Person做梦的方法
func (p Person) Dream() {
	zap.S().Debugf("%s的梦想是好好学习Go语言！", p.name)
}

type Person2 struct {
	name   string
	age    int8
	dreams []string
}

func (p *Person2) SetDreams(dreams []string) {
	p.dreams = make([]string, len(dreams))
	copy(p.dreams, dreams)
}

//使用值接收者：传拷贝进去
func (p Person) setAge2(newAge int8) {
	p.age = newAge
}

//使用指针接收者：传内存地址进去
func (p *Person) setAge(newAge int8) {
	p.age = newAge
}

func main() {
	InitLogger()
	p := newPerson("小王子", 18)
	p.Dream()
	zap.S().Debug(p.age) // 18
	p.setAge2(20)
	zap.S().Debug(p.age) // 18
	p.setAge(20)
	zap.S().Debug(p.age) // 20

	p1 := Person2{
		name: "小王子",
		age:  20,
	}
	data := []string{"吃饭", "睡觉", "打豆豆"}
	p1.SetDreams(data)
	zap.S().Debug(p1.dreams)
	data[1] = "不睡觉"
	p1.SetDreams(data)
	zap.S().Debug(p1.dreams)

}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().Panicf("日志初始化错误,error:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
