package main

import (
	"fmt"
	"sync"
)

// 启动多个goroutine
// 10个goroutine是并发执行的，而goroutine的调度是随机的。

// 声明全局等待组变量
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // 告知当前goroutine完成
	fmt.Println("你好", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 登记1个goroutine
		go hello(i)
	}
	wg.Wait() // 阻塞等待当前goroutine完成
}

// func main() {
// 	for i := 0; i < 5; i++ {
// 		go func() {
// 			fmt.Println(i)
// 		}()
// 	}
// }
