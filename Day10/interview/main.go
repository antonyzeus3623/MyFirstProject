package main

import "fmt"

// 定义x一次走1步，y一次走2步，让xy从起点出发去走，如果某一时刻他们能在同一节点相遇，那就说明这个链表有闭环。
// type a struct {
// 	val  int
// 	next *a
// }

func f(n int) int {
	// var x int
	// var y int
	// for i := 1; i <= n; i++ {

	// }
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f(n-1) + f(n-2)

}
func main() {
	fmt.Println(f(10))
}
