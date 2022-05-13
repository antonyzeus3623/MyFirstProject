package main

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"./time.log", "stdout"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化失败, err:%v", err.Error())
		return
	}
	zap.ReplaceGlobals(_logger)
}

func timeDemo() {
	now := time.Now() // 获取当前时间
	zap.S().Debugf("current time:%v", now)

	Year := now.Year()
	Month := now.Month()
	Day := now.Day()
	Hour := now.Hour()
	Minute := now.Minute()
	Second := now.Second()
	zap.S().Debugf("Year:%v Month:%v Day:%v Hour:%v Minute:%v Second:%v", Year, Month, Day, Hour, Minute, Second)
}

// 使用time.Unix() 函数将时间戳转换为时间格式
func timeUnixDemo(timestamp int64) {
	// 时间戳
	// Unix Time是自1970年1月1日 00:00:00 UTC 至当前时间经过的总秒数
	now := time.Now() // 获取当前时间
	zap.S().Debug(now.Unix())
	zap.S().Debug(now.UnixNano())

	timeObj := time.Unix(timestamp, 0) // 将时间戳转换为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	zap.S().Debugf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

}

// Add 求之后的时间
func addTime() {
	now := time.Now()
	later := now.Add(time.Hour)
	zap.S().Debug(later)
}

// 定时器
func tickDemo() {
	ticker := time.Tick(time.Second) // 设置一个1秒的定时器
	for i := range ticker {
		zap.S().Debug(i) // 每秒都会执行的任务
	}
}

// 时间格式化
func timeFormat() {
	now := time.Now()
	// 格式化的模板为 2006-01-02 15:04:05

	// 24小时制
	zap.S().Debug(now.Format("2006-01-02 15:04:05 Mon Jan"))
	zap.S().Debug(now.Format("2006-01-02 15:04:05"))
	// 12小时制
	zap.S().Debug(now.Format("2006-01-02 03:04:05 PM Mon Jan"))

	// 小数点后写0，因为有3个0所以格式化输出的结果也保留3位小数
	zap.S().Debug(now.Format("2006/01/02 15:04:05.000")) // 2022/02/27 00:10:42.960
	// 小数点后写9，会省略末尾可能出现的0
	zap.S().Debug(now.Format("2006/01/02 15:04:05.999")) // 2022/02/27 00:10:42.96

	// 只格式化时分秒部分
	zap.S().Debug(now.Format("15:04:05"))
	// 只格式化日期部分
	zap.S().Debug(now.Format("2006.01.02"))
}

// 解析字符串格式的时间
func parseTime() {

	startTime := time.Now().UnixNano()
	// time.Parse在解析时不需要额外指定时区信息
	// 在没有时区指示符的情况下，time.Parse 返回UTC时间
	timeObj, err := time.Parse("2006/01/02 15:04:05", "2022/02/14 13:14:20")
	if err != nil {
		zap.S().Debug(err.Error())
		return
	}
	zap.S().Debug(timeObj)

	// 在有时区指示符的情况下，time.Parse 返回对应时区的时间表示
	// RFC3339     = "2006-01-02T15:04:05Z07:00"
	timeObj, err = time.Parse(time.RFC3339, "2022-10-05T11:25:20+08:00")
	if err != nil {
		zap.S().Debug(err.Error())
		return
	}
	zap.S().Debug(timeObj)

	// time.ParseInLocation函数需要在解析时额外指定时区信息
	now := time.Now()
	zap.S().Debug(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		zap.S().Debug(err.Error())
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err = time.ParseInLocation("2006/01/02 15:04:05", "2022/02/15 13:14:20", loc)
	if err != nil {
		zap.S().Debug(err.Error())
		return
	}
	zap.S().Debug(timeObj)
	zap.S().Debug(timeObj.Sub(now)) // 求时间差

	endTime := time.Now().UnixNano()
	second := float64((endTime - startTime) / 1e9)
	Milliseconds := float64((endTime - startTime) / 1e6)
	nanoSeconds := float64(endTime - startTime)
	zap.S().Debug(second)
	zap.S().Debug(Milliseconds)
	zap.S().Debug(nanoSeconds)
}

// 获取当前时间，格式化输出为2017/06/19 20:30:05格式
func exercise1() {
	now := time.Now()
	zap.S().Debug(now.Format("2017/06/19 20:30:05"))
}

func main() {
	InitLogger()
	// timeDemo()

	// timeUnixDemo(1652405501)
	// addTime()
	// tickDemo()
	// timeFormat()
	parseTime()
	// exercise1()
}
