package main

import "fmt"

func add(a int, b int) {
	var c int
	c = a + b
	fmt.Println(c)
}
func mutiply(a int, b int) {
	var c int
	c = a*b
	fmt.Println(c)
func main() {
	x := 4
	y := 5
	add(x, y)
	mutiply(x, y)
}
