package main

import (
	"fmt"
	"reflect"
)

// 反射
func main() {
	var a float64 = 3.14
	fmt.Println("type", reflect.TypeOf(a))
	fmt.Println("value", reflect.ValueOf(a))

}
