package main

import (
	"flag"
	"fmt"
)

//flag获取命令行参数
func main() {
	//创建一个标志位参数
	name := flag.String("name", "王杰", "请输入姓名")
	age := flag.Int("age", 18, "请输入真是年龄")
	//使用
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
}
