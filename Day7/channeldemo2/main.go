package main

import "fmt"

// 只接收通道与只发送通道
// <- chan int // 只接收通道，只能接收不能发送
// chan <- int // 只发送通道，只能发送不能接收

// 单向通道示例

// Producer 返回一个通道
// 并持续将符合条件的数据发送至返回的通道中
// 数据发送完成后会将返回的通道关闭
func Producer() <-chan int { //Producer 返回一个只接收通道
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// Consumer 从通道中接收数据进行计算
func Consumer(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

func main() {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res)
}
