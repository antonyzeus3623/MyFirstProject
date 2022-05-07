package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件读写操作

// 文件的读取
func readFile() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%s", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 1024)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Printf("file read failed, err:%s", err)
		return
	}
	fmt.Printf("读取了%d字节数据", n)
	fmt.Println(string(tmp[:n]))
}

// 文件的循环读取
func loopRead() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file failed, err:%s", err)
		return
	}
	// 关闭文件
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 1024)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Printf("file read failed, err:%s", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio读取文件
func bufioRead() {
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file filed, err:%s", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //  ReadString('\n')在读到文件最后一行时，会同时返回内容line和io.EOF。
		if err == io.EOF {                   // io.EOF表示由文件结束引起的读取失败
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutilRead
func ioutilRead() {
	content, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

// writeTest
func writeTest() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) //用逻辑或操作"|"实现不同的组合
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "user=mysqlusr\n"
	file.Write([]byte(str))               // 写入字节切片数据
	file.WriteString("password=password") // 直接写入字符串数据
}

// bufio.NewWriter
func bufioWrite() {
	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("尝试写入数据至文件中\n") // 将数据先写入缓存
	}
	writer.Flush() // 将缓存中的内容写入文件

}

// ioutil.WriteFile
func ioutilWrite() {
	str := "尝试字符串写入"
	err := ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func main() {
	// readFile()
	// loopRead()
	// bufioRead()
	// ioutilRead()
	// writeTest()
	bufioWrite()
	// ioutilWrite()
}
