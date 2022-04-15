package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"unicode"
)

func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["王涛"] = 99
	scoreMap["李瑶"] = 95
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["王涛"])
	fmt.Printf("type of scoreMap: %T\n", scoreMap)

	//map在声明时填充元素
	userInfo := map[string]string{
		"username": "王杰",
		"password": "12345",
	}
	fmt.Println(userInfo)

	//判断某个键是否存在
	scoreMap2 := map[string]int{
		"ada":    95,
		"antony": 98,
	}
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap2["ada"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//map的遍历：使用for range遍历map
	scoreMap3 := map[string]int{
		"小红": 95,
		"小花": 98,
		"小白": 100,
	}
	for k, v := range scoreMap3 {
		fmt.Println(k, v)
	}
	//使用delete()函数删除键值对
	delete(scoreMap3, "小花")
	for k, v := range scoreMap3 {
		fmt.Println(k, v)
	}

	//按照指定顺序遍历map
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子
	//随机生成一组scoreMap4数据
	var scoreMap4 = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap4[key] = value
	}
	//取出map中的所有key存入切片keys
	keys := make([]string, 0, 200)
	for key := range scoreMap4 {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap4[key])
	}

	//map和slice组合
	//元素为map类型的切片
	var s1 = make([]map[string]string, 4, 10) //s1切片已初始化，但map还未初始化
	s1[0] = make(map[string]string, 8)        //对map进行初始化
	s1[0]["四川"] = "成都"
	s1[0]["浙江"] = "杭州"
	s1[0]["安徽"] = "合肥"
	fmt.Println(s1[0])

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30, 40, 50}
	fmt.Println(m1)

	/*统计一个字符串中每个单词出现的次数，主要思路是：

	  1.将字符串中的单词分割并存入slice
	  2.从slice中取出单词并存入map，存入过程中完成计数
	  3.打印结果
	  编写过程中遇到的主要问题：字符串如果是完整的句子，就会包含空格和标点符号，直接调用strings包中的Split函数没有办法进行很好的切分。	*/
	ss := "how are you? I am fine,thank you."
	//1.将字符串中的单词分割并存入slice（下面的方法可以识别出空格和标点符号等多种分隔符的字符串）
	var ss1 = make([]string, 0, 10) //初始化时，长度要为0
	var word string                 //定义一个变量word来接收单词
	word = ""
	for _, w := range ss { //遍历字符串
		if !unicode.IsLetter(w) { //遇到非字母字符时，说明一个单词结束，将单词存入slice中，并且重置word变量
			if word != "" { //由于可能存在两个或多个非字母字符相邻，所以在录入slice时要先进行一次判断
				ss1 = append(ss1, word)
			}
			word = ""
			continue
		} else {
			word = fmt.Sprintf("%s%c", word, w) // 字母元素拼接到word后面，组成单词
		}
	}
	//2.从slice中取出单词并存入map，存入过程中完成计数
	wordsMap := make(map[string]int, 10) //初始化一个map用来统计和存储结果
	for _, key := range ss1 {
		wordsMap[key]++
	}
	//3.打印结果
	for k, v := range wordsMap {
		fmt.Printf("word:%v  number of times:%d\n", k, v)
	}

	/* 2.观察下面代码，写出最终的打印结果。 */
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
