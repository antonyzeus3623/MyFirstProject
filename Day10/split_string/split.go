package splitstring

import "strings"

//Split切割字符串
//example
// abc, b => [a c] //字符串“abc”，按照b分割

func Split(str string, sep string) []string { //用sep去分割str
	// str="babcbdef"   sep="b"
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)

	}
	ret = append(ret, str)
	return ret
}
