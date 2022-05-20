package main

import (
	"fmt"
	"time"
)

// select多路复用 打印10以内的奇数
func selectdemo1() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

// select demo2
func selectdemo2() {
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(5 * time.Second):
		return
	}
}

func main() {
	// selectdemo1()
	selectdemo2()
}
