package main

import "fmt"

//整型
func main() {
	//十进制数
	var a int = 10
	fmt.Printf("a:%d\n", a)
	fmt.Printf("a:%b\n", a) //占位符：%b表示二进制数
	fmt.Printf("a:%c\n", a)
	fmt.Printf("a:%i\n", a)

	//八进制数
	var b int = 077
	fmt.Printf("b:%o\n", b)

	//十六进制数
	var c int = 0xff
	fmt.Printf("c:%x\n", c)
	fmt.Printf("c:%X\n", c)
}
