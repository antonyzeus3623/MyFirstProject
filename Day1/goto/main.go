package main

import "fmt"

//goto语句通过标签进行代码间的无条件跳转

func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if j == 'C' {
				goto breakTag //跳到指定的标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
breakTag: //label标签
	fmt.Println("结束for循环")
}
