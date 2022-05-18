package main

import (
	"fmt"
	"sync"
)

var b chan int
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println("不带缓冲区的通道")
	b = make(chan int) // 不带缓冲区的通道初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道中取到了", x)
	}()

	b <- 10
	fmt.Println("将10发送至通道成功")
	wg.Wait()
}

func bufChannel() {
	fmt.Println("带缓冲区的通道")
	b = make(chan int, 2) // 带缓冲区的通道初始化
	b <- 10
	fmt.Println("将10发送至通道成功")
	x := <-b
	fmt.Println("从通道中取到了", x)
}

func main() {
	fmt.Println(b)
	noBufChannel()
	bufChannel()

}
