package main

import "fmt"

func main() {
	// &: 取地址
	// *: 根据地址取值
	a := 10
	b := &a
	fmt.Printf("a:%d  ptr:%p\n", a, &a) //a:10  ptr:0xc000012088
	fmt.Printf("b:%p  type:%T\n", b, b) //b:0xc000012088  type:*int  *int：表示int类型的指针
	fmt.Println(&b)                     //0xc000006028
	c := *b
	fmt.Printf("c:%d  type:%T\n", c, c) //c:10  type:int

	//new函数为指针变量初始化（分配内存空间）
	var d *int
	d = new(int)
	*d = 10
	fmt.Printf("d:%d", *d)

	//make函数用于slice、map以及channel的初始化
	var e map[string]int
	e = make(map[string]int, 10)
	e["古力娜扎"] = 18
	fmt.Println(e)
}
