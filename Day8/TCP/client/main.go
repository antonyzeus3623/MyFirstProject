package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TCP client
func main() {
	// 1.与server建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dail 127.0.0.1:20000 failed, err:%v\n", err)
		return
	}
	defer conn.Close() // 关闭连接
	// 2.进行数据收发
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入，ReadString 一次读取一行
		inputInfo := strings.Trim(input, "\r\n") // Trim()去掉元素，\r\n：回车换行
		if strings.ToUpper(inputInfo) == "Q" {   // 如果输入Q就退出
			return
		}
		_, err := conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:]) // 文件接收
		if err != nil {
			fmt.Println("recv failed, err：", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
