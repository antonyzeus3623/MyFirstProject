package main

import (
	"fmt"
	"reflect"
)

// 结构体反射
type Student struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

// 反射
func main() {
	var a float64 = 3.14
	reflectType(a)

	var b int16 = 20
	reflectType(b)

	var c int64 = 100
	fmt.Printf("修改之前值为：%d\n", c)

	reflectSetValue(&c)
	fmt.Printf("修改之后值为：%d\n", c)

	// 结构体反射
	stu1 := Student{
		Name:  "小王子",
		Score: 100,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind()) // student struct
	// 通过for循环遍历结构体的所有字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	// 通过字段名获取指定结构体字段信息
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}
	printMethod(stu1)

}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

// 通过反射设置变量的值
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s Student) Study() string {
	msg := "好好学习，天天向上"
	fmt.Println(msg)
	return msg
}

func (s Student) Sleep() string {
	msg := "好好睡觉，快快长大"
	fmt.Println(msg)
	return msg
}

// 遍历打印s包含的方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
