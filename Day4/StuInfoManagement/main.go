package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// 定义一个Student结构体
type Student struct {
	Id    int
	Name  string
	Age   int8
	Score float32
}

var (
	stuList []Student
	stuNum  = 0 // 用于生成id
)

func getStuList() { // 返回学生列表， 用于查询
	// return stuList
	if len(stuList) == 0 {
		zap.S().Debug("学生列表为空")
	} else {
		for _, k := range stuList {
			zap.S().Debugf("学号：%d，姓名：%s，年龄：%d，分数：%f", k.Id, k.Name, k.Age, k.Score)
		}
	}
}

func newStu(id int, name string, age int8, score float32) *Student { //构造函数
	return &Student{
		Id:    id,
		Name:  name,
		Age:   age,
		Score: score,
	}
}

// 添加学生
func addStu() {
	var id = stuNum
	stuNum += 1
	var name string
	var age int8
	var score float32
	// fmt.Println("请输入学生姓名：")
	zap.S().Debug("请输入学生姓名：")
	_, err := fmt.Scan(&name)
	zap.S().Debug("请输入学生年龄：")
	_, err = fmt.Scan(&age)
	zap.S().Debug("请输入学生分数：")
	_, err = fmt.Scan(&score)
	if err != nil {
		zap.S().DPanicf("您的输入信息有误，err:%s", err.Error())
	}
	newStu(id, name, age, score)
	stuList = append(stuList, *newStu(id, name, age, score))
}

// 输入学生id进行删除
func deleteStu() {
	var i int
	zap.S().Debug("请输入需要删除的学生id：")
	_, err := fmt.Scan(&i)
	if err != nil {
		zap.S().DPanicf("您的输入信息有误，err:%s", err.Error())
	}
	for _, v := range stuList {
		if i == v.Id {
			stuList = append(stuList[:i], stuList[i+1:]...)
		}
	}
}

// 输入学生id，对其进行信息编辑、修改
func editStuInfo() {
	var i int
	zap.S().Debug("请输入需要修改信息的学生id：")
	_, err := fmt.Scan(&i)
	if err != nil {
		zap.S().DPanicf("您的输入信息有误, err:%s", err.Error())
	}
	for index, v := range stuList {
		if i == v.Id {
			zap.S().Debug("请输入学生姓名：")
			_, err := fmt.Scan(&stuList[index].Name) // 直接修改对应的stuList[i]的内部变量值
			zap.S().Debug("请输入学生年龄：")
			_, err = fmt.Scan(&stuList[index].Age)
			zap.S().Debug("请输入学生分数：")
			_, err = fmt.Scan(&stuList[index].Score)
			if err != nil {
				zap.S().DPanicf("您的输入信息有误，err:%s", err.Error())
			}
		}
	}
}

func main() {
	InitLogger()
	for {
		var do int8
		zap.S().Debug("请选择您要执行的操作:1.添加学生信息 2.查看学生列表 3.修改学生信息 4.删除学生信息 5.退出")
		_, err := fmt.Scan(&do)
		if err != nil {
			zap.S().DPanicf("您的输入信息有误，err:%s", err.Error())
		}
		switch do {
		case 1:
			addStu()
		case 2:
			getStuList()
		case 3:
			editStuInfo()
		case 4:
			deleteStu()
		case 5:
			os.Exit(1)
		default:
			zap.S().DPanicf("您的输入信息有误，err:%s", err.Error())
		}
	}
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "student.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化失败，err:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
