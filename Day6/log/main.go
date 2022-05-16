package main

import (
	"fmt"
	"log"
	"os"
)

func logDemo() {
	log.Println("1. 这是一条普通的日志")
	v := "普通的"
	log.Printf("2. 这是一条%s的日志", v)
	log.Fatalln("3. 这是一条会触发Fatal的日志")
	log.Panicln("4. 这是一条会触发Panic的日志")
}

func loggerDemo() {
	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败：", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Llongfile | log.Lmicroseconds)
	log.Println("这是一条普通的日志") // 2022/05/13 13:56:49.703599 D:/MyFirstProject/Day6/log/main.go:15: 这是一条普通的日志

	// 配置日志前缀
	log.SetPrefix("[日志前缀]")
	log.Println("这是一条带有前缀的日志")

}

func InitLogger() {
	logFile, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败：", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Llongfile | log.Lmicroseconds)
}

func main() {
	InitLogger()
	// logDemo()
	// loggerDemo()

	logger := log.New(os.Stdout, "<new>", log.Llongfile|log.Ldate|log.Lmicroseconds)
	logger.Println("这是自定义的logger记录的日志。")
}
