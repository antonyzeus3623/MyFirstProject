package main

import "fmt"

var a float32 = 3.14

const b = 5

//init初始化函数
func init() {
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	sayHi()
}

func sayHi() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("你好，世界！")
}
