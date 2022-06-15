package main

// 使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。
// 1)开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
// 2)开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// 3)主 goroutine 从resultChan取出结果并打印到终端输出

import (
	"fmt"
	"math/rand"
	"time"
)

var jobChan = make(chan int64)
var resultChan = make(chan map[int64]int64)

//generateRandomNumbers 用于持续循环生成int64的随机数，并传到jobChan channel
func generateRandomNumbers() {
	for {
		x := rand.Int63()
		jobChan <- x
		time.Sleep(time.Second)
	}
}

//sumRandomNumbers 从通道jobChan中读取数据，然后计算各个位数之和存入resultChan channel中
func sumRandomNumbers() {
	for {
		RandomNumber := <-jobChan
		value := RandomNumber
		var sum int64 = 0
		for value > 0 {
			sum = sum + value%10
			value = value / 10
		}
		resultChan <- map[int64]int64{RandomNumber: sum}
	}
}

func main() {
	//1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	go generateRandomNumbers()
	//2. 开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	for i := 0; i < 24; i++ {
		go sumRandomNumbers()
	}
	//3. 主goroutine从resultChan取出结果并打印到终端输出
	for i := range resultChan {
		for RandomNumber, sum := range i {
			fmt.Printf("Random Number:%d \t Sum:%d \n", RandomNumber, sum)
		}
	}
}
