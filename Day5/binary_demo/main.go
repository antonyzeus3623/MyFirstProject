package main

import "fmt"

// 二进制的用途示例

const (
	eat   int = 4
	sleep int = 2
	play  int = 1
)

// 111
// 左边的1表示吃饭 100
// 中间的1表示睡觉 010
// 右边的1表示娱乐 001

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat | sleep | play)
}
