package main

import "fmt"

// 定义学生结构体
type Student struct {
	name  string
	age   int8
	score float32
}

// 定义老师结构体
type Teacher struct {
	name string
	age  int8
}

// 定义接口
type ClassAtion interface {
	Say()
	AfterClass()
}

// 学生继承接口(接口的2个函数 需要全部实例化才行)
func (s Student) Say() {
	fmt.Printf("my name is %s\n", s.name)
}

func (s Student) AfterClass() {
	fmt.Println("go home after class")
}

// 老师继承接口
func (t Teacher) Say() {
	fmt.Printf("my name is %s\n", t.name)
}

func (t Teacher) AfterClass() {
	fmt.Println("go to office after class")
}

func testInterface() {

	// 定义学生和老师
	stu := Student{
		name:  "明明",
		age:   18,
		score: 99.5,
	}
	teacher := Teacher{
		name: "小莉老师",
		age:  25,
	}

	// 定义一个接口
	var x ClassAtion
	// 接口指向学生
	x = stu
	// 使用接口里面的函数
	x.Say()
	x.AfterClass()

	// 改变接口指向老师并使用
	x = teacher
	x.Say()
	x.AfterClass()

	c, ok := x.(Teacher)
	if ok {
		fmt.Println("类型断言成功")
		fmt.Println(c)
	} else {
		fmt.Println("类型断言失败")
	}

}

func main() {
	testInterface()
}
