package main

import (
	"MyFirstProject/Day8/filetrans/logger"
	"fmt"
	"io"
	"net"
	"os"

	"go.uber.org/zap"
)

/*
1.创建监听listener，程序结束时关闭。
2.阻塞等待客户端连接，程序结束时关闭conn。
3.读取客户端发送文件名。保存fileName。
4.回发“ok”给客户端做应答
5.封装函数 RecvFile接收客户端发送的文件内容。传参fileName 和conn
6.按文件名Create文件，结束时Close
7.循环Read客户端发送的文件内容，当读到EOF说明文件读取完毕
8.将读到的内容原封不动Write到创建的文件中
*/
func RecvFile(fileName string, conn net.Conn) {
	// 创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		zap.S().DPanicf("Create err: ", err)
		return
	}
	defer f.Close()

	// 接收client端发送的文件内容，原封不动写入文件中
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				zap.S().DPanic("文件接收完毕")
			} else {
				zap.S().DPanicf("Read err: ", err)

			}
			return
		}
		f.Write(buf[:n]) // 写入文件，读多少写多少
	}
}

func main() {
	logger.InitLogger()
	fmt.Println("等待文件传输中...")
	fmt.Println("注意：文件传输完成将自动关闭本程序")
	// 创建监听
	listener, err := net.Listen("tcp", "127.0.0.1:8005") // 127.0.0.1:8005 10.1.6.242:20000
	if err != nil {
		zap.S().DPanicf("Listen err: ", err)
		return
	}
	defer listener.Close()

	// 阻塞等待client端连接
	conn, err := listener.Accept()
	if err != nil {
		zap.S().DPanicf("Accept err: ", err)
		return
	}
	defer conn.Close()

	//读取client端发送的文件名
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		zap.S().DPanicf("Read err: ", err)
		return
	}
	fileName := string(buf[:n]) // 保存文件名

	// 回复ok给client端
	conn.Write([]byte("ok"))

	// 接收文件内容
	RecvFile(fileName, conn) //封装函数接收文件内容，传fileName和conn
}
