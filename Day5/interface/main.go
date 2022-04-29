package main

import "fmt"

// Mover 接口
type Mover interface {
	Move()
}

// Sayer 接口
type Sayer interface {
	Say()
}

// Dog 定义狗结构体类型
type Dog struct {
	Name string
}

// Cat 定义猫结构体类型
type Cat struct {
	Name string
}

// Car 定义汽车结构体类型
type Car struct {
	Brand string
}

// 实现Mover接口
func (d Dog) Move() {
	fmt.Printf("名字叫%s的小狗会动！\n", d.Name)
}

// 实现Sayer接口
func (d Dog) Say() {
	fmt.Printf("名字叫%s的小狗会唱歌！\n", d.Name)
}

// 实现Mover接口
func (c *Cat) Move() {
	fmt.Printf("名字叫%s的小猫会动！\n", c.Name)
}

//实现Mover接口
func (c Car) Move() {
	fmt.Printf("%s的速度70迈！\n", c.Brand)
}

// 接口实现示例1
func testInterface1() {
	var x Mover

	var d1 = Dog{
		Name: "小黄",
	}
	x = d1
	x.Move()

	var d2 = &Dog{
		Name: "毛毛",
	}
	x = d2
	x.Move()

	var c1 = &Cat{}
	x = c1
	x.Move()
}

// 同一类型实现不同接口
func testInterface2() {
	var d = Dog{Name: "小布"}
	var m Mover = d
	var s Sayer = d
	m.Move() // 对Sayer类型调用Say方法
	s.Say()  // 对Mover类型调用Move方法
}

// 多类型实现同一接口
func testInterface3() {
	var obj Mover
	obj = Dog{Name: "草莓"}
	obj.Move()

	obj = Car{Brand: "奔驰"}
	obj.Move()
}

// 空接口
type Any interface{}

// 实现空接口
func testInterface4() {
	var x Any
	x = "你好"
	fmt.Printf("Type:%T  Value:%v\n", x, x)
	x = 18
	fmt.Printf("Type:%T  Value:%v\n", x, x)
	x = true
	fmt.Printf("Type:%T  Value:%v\n", x, x)
	x = Dog{}
	fmt.Printf("Type:%T  Value:%v\n", x, x)
}

func main() {
	testInterface1()
	testInterface2()
	testInterface3()
	testInterface4()
}
