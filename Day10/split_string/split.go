package splitstring

import (
	"strings"
)

//Split切割字符串
//example
// abc, b => [a c] //字符串“abc”，按照b分割

func Split(str string, sep string) []string { //用sep去分割str
	// str="babcbdef"   sep="b"
	var ret = make([]string, 0, strings.Count(str, sep)+1) //申请内存
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)

	}
	ret = append(ret, str)
	return ret
}

// Fib 是一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
