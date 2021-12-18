package main

import (
	"MyFirstProject/snow"
	"fmt"
)

func getInput() {
	var x int
	var y int
	var euma string
	fmt.Print("输入x: ")
	fmt.Scanf("%d\n", &x)

	fmt.Print("输入+ - * / 符号: ")
	fmt.Scanf("%s\n", &euma)

	fmt.Print("输入y: ")
	fmt.Scanf("%d\n", &y)

	data := snow.NewGetNum(x, y)
	switch euma {
	case "+":
		// fmt.Println("你输入的是", euma)
		fmt.Println("结果为：", data.Add(data))
	case "-":
		fmt.Println(data.Minus(data))
	case "*":
		data.Take(data)
	case "/":
		data.Division(data)
	}

}
func main() {
	for {
		getInput()
	}
}
