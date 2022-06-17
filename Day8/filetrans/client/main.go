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
1.提示用户输入文件名。接收文件名path（含访问路径）
2.使用os.Stat()获取文件属性，得到纯文件名（去除访问路径）
3.主动连接server端，结束时关闭连接
4.给server端发送文件名 conn.Write()
5.读取server端回复的确认数据 conn.Read()
6.判断是否为“ok”。如果是，封装函数SendFile()发送文件内容。传参path和conn
7.只读Open文件, 结束时Close文件
8.循环读文件，读到EOF终止文件读取
9.将读到的内容原封不动Write给server端
*/
func SendFile(path string, conn net.Conn) {
	// 以只读方式打开文件
	f, err := os.Open(path)
	if err != nil {
		zap.S().DPanicf("os.Open err: ", err)
		return
	}
	defer f.Close() // 发送结束关闭文件

	// 循环读取文件，原封不动地写给server端
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf) // 读取文件内容到切片缓冲区
		if err != nil {
			if err == io.EOF {
				zap.S().DPanic("文件发送完毕 ")
			} else {
				zap.S().DPanicf("f.Read err: ", err)
			}
			return
		}
		conn.Write(buf[:n]) //原封不动写给server端
	}
}

func main() {
	logger.InitLogger()
	// 提示输入文件名
	fmt.Println("请输入需要传输的文件: ")
	var path string
	fmt.Scan(&path)

	// 获取文件名 fileInfo.Name()
	fileInfo, err := os.Stat(path)
	if err != nil {
		zap.S().DPanicf("os.Stat err: ", err)
		return
	}

	// 主动连接server端
	conn, err := net.Dial("tcp", "127.0.0.1:8005") // 127.0.0.1:8005 10.1.6.242:20000
	if err != nil {
		zap.S().DPanicf("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	// 给server端先发送文件名
	_, err = conn.Write([]byte(fileInfo.Name()))
	if err != nil {
		zap.S().DPanicf("conn.Write err: ", err)
		return
	}

	// 读取server端回复的确认数据--ok
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		zap.S().DPanicf("conn.Read err: ", err)
		return
	}

	// 判断如果是ok，则发送文件内容
	if "ok" == string(buf[:n]) {
		SendFile(path, conn) // 封装函数读文件，发送给server端，需要path、conn
	}
}
