package main

import "fmt"

//数组的长度是数组类型的一部分
func main() {

	//数组的定义和初始化
	var a1 = [3]int{1, 2, 3}
	var a2 = [4]int{3, 3, 4}
	var a3 = [...]string{"北京", "广州", "成都", "上海"} //让编译器根据初始值的个数自行推断数组的长度
	a4 := [...]int{1: 2, 3: 5, 5: 1}             //使用指定索引值的方式来初始化数组
	fmt.Println(a1, a2, a3, a4)
	fmt.Printf("type of a1:%T,\ntype of a2:%T,\ntype of a3:%T,\ntype of a4:%T\n", a1, a2, a3, a4)

	//数组的遍历
	//方法1：for循环遍历数组(根据索引遍历)
	for i := 0; i < len(a3); i++ {
		fmt.Println(a3[i])
	}
	//方法2：for range遍历
	for index, value := range a3 {
		fmt.Println(index, value)
	}
	// for _, v := range a3 {
	// 	fmt.Println(v)
	// }

	//多维数组
	citys := [...][2]string{
		{"四川", "成都"},
		{"浙江", "杭州"},
		{"安徽", "合肥"},
	}
	for _, v1 := range citys {
		for _, v2 := range v1 {
			fmt.Printf("%s\t\n", v2) //\t：表示制表符，以表格的形式输出
		}
	}

	//数组练习题
	//1:求数组[1, 3, 5, 7, 8]所有元素的和
	var b1 = [...]int{1, 3, 5, 7, 8}
	var sum int
	for _, v := range b1 {
		sum += v
	}
	fmt.Printf("数组b1之和为:%v\n", sum)

	//2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)
	b2 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(b2); i++ {
		for j := i + 1; j < len(b2); j++ {
			if b2[i]+b2[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}

}
