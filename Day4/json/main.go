package main

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

//结构体与JSON序列化
type Student struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

type Class struct {
	Title    string `json:"title"`
	Students []*Student
}

//定义一个结构体
type Monster struct {
	Name  string
	Age   int
	Sal   float64
	Skill string
}

func testClass() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			Id:     i,
		}
		c.Students = append(c.Students, stu)
	}

	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		zap.S().DPanicf("Json序列化失败,err：", err.Error())
		return
	}
	zap.S().Debugf("json:%s", data)

	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"Id":0,"Gender":"男","Name":"stu00"},{"Id":1,"Gender":"男","Name":"stu01"},{"Id":2,"Gender":"男","Name":"stu02"},{"Id":3,"Gender":"男","Name":"stu03"},{"Id":4,"Gender":"男","Name":"stu04"},{"Id":5,"Gender":"男","Name":"stu05"},{"Id":6,"Gender":"男","Name":"stu06"},{"Id":7,"Gender":"男","Name":"stu07"},{"Id":8,"Gender":"男","Name":"stu08"},{"Id":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		zap.S().DPanicf("Json反序列化失败,error:", err.Error())
		return
	}
	zap.S().Debugf("c1:%#v", c1)
}

//将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "小林"
	a["age"] = 20
	a["address"] = "中国"
	//将map序列化
	data, err := json.Marshal(a)
	if err != nil {
		zap.S().DPanicf("json marshal failed, err:%s", err.Error())
	}
	zap.S().Debugf("map序列化后的结果为：%s", data)
}

//对slice进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Jack"
	m1["age"] = 21
	m1["address"] = "American"

	// var m2 map[string]interface{
	m2 := make(map[string]interface{})
	m2["name"] = "Tom"
	m2["age"] = 25
	m2["address"] = [2]string{"成都", "四川"}
	slice = append(slice, m1, m2)
	//将slice序列化
	data, err := json.Marshal(slice)
	if err != nil {
		zap.S().DPanicf("json marshal failed, err:%s", err.Error())
	}
	zap.S().Debugf("slice序列化后的结果为：%s", data)
}

//对基本数据类型进行序列化
func testFloat64() {
	var num float64 = 123.45
	//将float64序列化
	data, err := json.Marshal(num)
	if err != nil {
		zap.S().DPanicf("json marshal failed, err：%s", err.Error())
	}
	zap.S().Debugf("float64序列化后的结果为：%s", data)
}

// 将json字符串反序列化成struct
func testUnmarshalStruct() {
	//说明:str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := `{"Name":"牛魔王","Age":500,"Sal":10.2,"Skill":"牛魔拳"}`
	//定义一个Monster实例
	var monster Monster
	//将monster反序列化
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		zap.S().DPanicf("struct反序列化失败,err:%s", err.Error())
	}
	zap.S().Debugf("struct反序列化的结果为：%v", monster)

}

// 将json字符串反序列化成map
func testUnmarshalMap() {
	str := `{"Name":"牛魔王","Age":500,"Sal":10.2,"Skill":"牛魔拳"}`
	var a map[string]interface{}
	// 反序列化
	// 注意:反序列化map,不需要make,因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		zap.S().DPanicf("map反序列化失败,err:%s", err.Error())
	}
	zap.S().Debugf("map反序列化的结果为：%v", a)
}

// 将json字符串反序列化成slice
func testUnmarshalSlice() {
	str := `[{"address":"北京","age":"7","name":"jack"},` + `{"address":["墨西哥","夏威夷"],"age":"20","name":"tom"}]`
	var s []map[string]interface{}
	//反序列化
	//注意：不需要make,因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &s)
	if err != nil {
		zap.S().DPanicf("slice反序列化失败,err:%s", err.Error())
	}
	zap.S().Debugf("slice反序列化的结果为：%v", s)
}

func main() {
	InitLogger()
	testClass()
	testMap()
	testSlice()
	testFloat64()
	testUnmarshalStruct()
	testUnmarshalMap()
	testUnmarshalSlice()
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr", "test.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化失败，err:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
