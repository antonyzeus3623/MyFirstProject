package main

import (
	"MyFirstProject/Day5/calc"
	"os"

	"fmt"
	"strconv"
)

func main() {
	for {
		var a, b, c string
		var result float64
		fmt.Println("请输入你要进行的运算，按“q”退出：")
		fmt.Scanf("%s\n", &b)
		fmt.Println("请输入第1个数字：")
		fmt.Scanf("%s\n", &a)
		fmt.Println("请输入第2个数字：")
		fmt.Scanf("%s\n", &c)
		aa, _ := strconv.ParseFloat(a, 64)
		cc, _ := strconv.ParseFloat(c, 64)
		switch b {
		case "+":
			result = calc.Add(aa, cc)
		case "-":
			result = calc.Sub(aa, cc)
		case "*":
			result = calc.Mul(aa, cc)
		case "/":
			result = calc.Div(aa, cc)
		case "q":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
		fmt.Println(result)
	}

}
