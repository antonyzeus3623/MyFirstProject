package main

import "fmt"

// channel声明和初始化
func channelDemo() {
	var ch1 chan int
	var ch2 chan bool
	var ch3 chan string
	ch4 := make(chan int)
	ch5 := make(chan bool, 1)
	fmt.Println(ch1, ch2, ch3, ch4, ch5)
}

// 多返回值模式
func f2(ch chan int) {
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Printf("value:%#v  ok:%#v\n", value, ok)
	}
}

func main() {
	channelDemo()
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}
