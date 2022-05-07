package main

import (
	"bufio"
	"fmt"
	"os"
)

// 获取用户输入，带空格
func userBufio() {
	fmt.Println("请输入：")
	var str string
	reader := bufio.NewReader(os.Stdin) // 获取标准输入  NewReader的参数就是接口类型
	str, _ = reader.ReadString('\n')    // ReadString('\n')：读到回车表示输入结束，会返回内容str和io.EOF。
	fmt.Printf("你输入的内容是：%s\n", str)
}
func main() {
	userBufio()
}
