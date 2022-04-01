package main

import "fmt"

func main() {
	figer := 3
	switch figer {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小指")
	default:
		fmt.Println("无效的输入")
	}

	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println("输入有误")
	}

	age := 18
	switch {
	case age < 20:
		fmt.Println("好好学习吧！")
		fallthrough
	case age >= 20 && age <= 30:
		fmt.Println("好好谈恋爱！")
	case age > 30 && age < 50:
		fmt.Println("好好工作吧")
	case age >= 50:
		fmt.Println("好好享受生活吧！")
	default:
		fmt.Println("此人已不存在")
	}

}
