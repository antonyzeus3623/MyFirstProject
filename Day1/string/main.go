package main

import (
	"fmt"
	"math"
)

func main() {

	//字符串修改
	s1 := "白萝卜"
	s2 := []rune(s1) //把字符串强制转换成一个rune切片
	s2[0] = '红'
	fmt.Println(string(s2)) //把rune切片强制转换成字符串
	s3 := "big"
	s4 := []byte(s3)
	s4[0] = 'p'
	fmt.Println(string(s4))
	s5 := "H"
	s6 := byte('H')
	fmt.Printf("s5:%T,s6:%T\n", s5, s6)

	//类型转换
	var a, b = 3, 4
	//因为math.Sqrt()接收的参数类型是float64，所以需要强制类型转换
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Printf("三角形的斜边长为:%d\n", c)

}
