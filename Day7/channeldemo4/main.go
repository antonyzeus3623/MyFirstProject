package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel练习
// 1. 启动一个goroutine，生成100个数发送到ch1中
// 2. 启动一个goroutine，从ch1中取值，计算其平方放到ch2中
// 3. 在main中，从ch2取值并打印出来

func makeNum() <-chan int { // 返回一个只接收通道
	ch1 := make(chan int, 1)
	go func() {
		for i := 0; i < 100; i++ {
			rand.Seed(time.Now().UnixNano()) // 添加随机数种子
			num := rand.Intn(100)
			ch1 <- num
		}
		close(ch1)
	}()

	return ch1
}

func getNum(ch <-chan int) chan int {
	ch2 := make(chan int, 1)
	go func() {
		for i := 0; i < 100; i++ {
			num := <-ch
			ch2 <- num * num
		}
		close(ch2)
	}()

	return ch2
}

func main() {
	ch1 := makeNum()
	res := getNum(ch1)
	for v := range res {
		fmt.Printf("%#v  ", v)
	}

}
