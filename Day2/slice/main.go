package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的定义
	var s1 []int    //定义一个存放int类型元素的切片s1
	var s2 []string //定义一个存放int类型元素的切片s2

	//初始化切片
	s1 = []int{2, 4, 6, 8}
	s2 = []string{"赵雅芝", "冯程程", "程淮秀", "白素贞"}
	fmt.Println(s1, s2)

	//长度和容量
	fmt.Printf("len(s1):%d,cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d,cap(s2):%d\n", len(s2), cap(s2))

	//2. 由数组得到切片
	a1 := [...]int{2, 4, 5, 7, 3, 9, 8, 2, 1}
	s3 := a1[0:5] //基于一个数组切割，左包含右不包含（左闭右开）
	s4 := a1[5:]
	s5 := a1[:5]
	s6 := s5[3:7] //对切片再执行切片时，上限边界是切片的容量cap(s5)，而不是长度
	//切片的容量是底层数组从切片的第一个元素到最后一个元素的数量
	fmt.Printf("len(a1):%d,cap(a1):%d\n", len(a1), cap(a1)) //len(a1):9,cap(a1):9
	fmt.Println(s3)
	fmt.Printf("len(s3):%d,cap(s3):%d\n", len(s3), cap(s3)) //len(s3):5,cap(s3):9
	fmt.Println(s4)
	fmt.Printf("len(s4):%d,cap(s4):%d\n", len(s4), cap(s4)) //len(s4):4,cap(s4):4
	fmt.Println(s5)
	fmt.Printf("len(s5):%d,cap(s5):%d\n", len(s5), cap(s5)) //len(s5):5,cap(s5):9
	fmt.Println(s6)
	fmt.Printf("len(s6):%d,cap(s6):%d\n", len(s6), cap(s6)) //len(s6):4,cap(s6):6

	//3.使用make函数动态构造切片
	a2 := make([]int, 2, 10)
	fmt.Println(a2)
	fmt.Printf("len(a2):%d,cap(a2):%d\n", len(a2), cap(a2)) //len(a2):2,cap(a2):10
	if len(a2) == 0 {
		fmt.Println("切片为空")
	} else {
		fmt.Println("切片不为空")
	}
	//切片的赋值拷贝,拷贝前后两个变量共享底层数组，对一个切片的修改会影响另一个切片的内容
	s7 := make([]int, 3)
	s8 := s7
	s8[0] = 100
	fmt.Printf("s7:%v\n", s7) //s7:[100 0 0]
	fmt.Printf("s8:%v\n", s8) //s8:[100 0 0]

	//切片的遍历（与数组一致）
	s9 := []int{1, 3, 5, 7}
	for i, v := range s9 {
		// fmt.Println(i, v)
		fmt.Printf("%v\t %v\t \n", i, v)
	}

	//append()方法为切片添加元素
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	var aSlice []int
	bSlice := []int{3, 4, 5, 6}
	aSlice = append(aSlice, 0, 1, 2)
	fmt.Println(aSlice)                //[0 1 2]
	aSlice = append(aSlice, bSlice...) //添加另一个切片中的元素（后面加…）
	fmt.Println(aSlice)                //[0 1 2 3 4 5 6]

	//使用copy()函数复制切片
	a := []int{1, 2, 3, 4, 5}
	c := make([]int, 5, 8)
	copy(c, a) // 使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(a)
	fmt.Println(c)
	c[0] = 100
	fmt.Println(a)
	fmt.Println(c)

	//从切片中删除元素
	//要从切片a中删除索引为index的元素，操作方法是a = append(a[:index], a[index+1:]...)
	d := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	d = append(d[:2], d[3:]...) //从切片d中删除索引为2的元素
	fmt.Println(d)

	//练习题1
	var e = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		e = append(e, fmt.Sprintf("%v", i))
	}
	fmt.Println(e)

	//练习题2
	var f = [...]int{3, 7, 8, 9, 1}
	var g = []int{}
	g = f[:]
	sort.Ints(g)
	fmt.Println(g)

}
