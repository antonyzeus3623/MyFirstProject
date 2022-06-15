package main

import (
	"fmt"
	"sync"
)

// sync.Mutex
// 多个goroutine资源竞态，使用sync.Mutex提供的两个方法：sync.lock和sync.unlock互斥锁
// 保证同一时间只有一个goroutine可以访问共享资源

var (
	x  int64
	wg sync.WaitGroup // 等待组
	m  sync.Mutex
)

// add 对全局变量x执行5000次加1操作
func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x += 1
		m.Unlock() // 修改后释放
	}
}

func main() {
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)

}
