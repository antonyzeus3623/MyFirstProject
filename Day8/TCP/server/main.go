package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server

// 处理函数
func processConn(conn net.Conn) {
	// 3.与客户端通信
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Printf("read from client failed, err:%v\n", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() {
	// 1.本地端口启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:20000") // 本地电脑测试 127.0.0.1:20000  李登祥电脑测试 10.1.6.242:20000
	if err != nil {
		fmt.Printf("start tcp server on 127.0.0.1:20000 failed, err:%v \n", err)
		return
	}
	// 2.等待别人来跟我建立连接
	for {
		conn, err := listener.Accept() // 建立连接
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		go processConn(conn) // 启动一个goroutine处理连接
	}
}
