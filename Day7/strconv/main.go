package main

import (
	"fmt"
	"strconv"
)

// string与int类型转换
func stringInt() {
	// Atoi()函数用于将字符串类型的整数转换为int类型
	s1 := "200"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Printf("字符串类型的整数转换为int类型失败，err：%v", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}

	i2 := 400
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)
}

// Parse转换字符串为给定类型
func parsetoother() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Printf("convert failed, err:%v", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", b, b)
	}

	f, err := strconv.ParseFloat("10", 64)
	if err != nil {
		fmt.Printf("convert failed, err:%v", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", f, f)
	}

	i, err := strconv.ParseInt("-2", 10, 64)
	if err != nil {
		fmt.Printf("convert failed, err:%v", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", i, i)
	}

	u, err := strconv.ParseUint("2", 10, 64)
	if err != nil {
		fmt.Printf("convert failed, err:%v", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", u, u)
	}
}

// 将给定类型数据格式化为string类型
func formatDemo() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatInt(-100, 10)
	s3 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s4 := strconv.FormatUint(20, 10)
	fmt.Printf("s1: type:%T value:%#v\n", s1, s1)
	fmt.Printf("s2: type:%T value:%#v\n", s2, s2)
	fmt.Printf("s3: type:%T value:%#v\n", s3, s3)
	fmt.Printf("s4: type:%T value:%#v\n", s4, s4)
}

func main() {
	stringInt()
	parsetoother()
	formatDemo()
}
