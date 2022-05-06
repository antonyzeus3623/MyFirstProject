package main

import (
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

// 编写代码利用反射实现一个ini文件的解析器程序。
// 根据ini文件的key，输出key对应的value

type MysqlConfig struct {
	Ip       string `ini:"ip"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
}

type Config struct {
	MysqlConfig `ini:"mysqlId"`
}

func loadIni(filename string, data interface{}) (err error) {
	// 1. 参数校验，参数必须是指针类型和结构体类型
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct {
		zap.S().Debug("type error should be struct ptr")
		return err
	}
	// 2. 打开文件
	file, err := ioutil.ReadFile(filename) // ReadFile读取文件中的所有数据，返回读取的数据和遇到的错误
	if err != nil {
		zap.S().DPanicf("open file error: %v", err.Error())
		return err
	}
	// 3. 一行一行读取文件，将各行内容保存在lineSlice切片中
	lineSlice := strings.Split(string(file), "\r\n") // \r\n：表示回车+换行
	// 定义section切片（小节切片）
	var section string
	var structName string
	// 3.1 遍历slice
	for idx, line := range lineSlice {
		// 3.2 去除首尾空格
		line = strings.TrimSpace(line)
		// 3.3 如果是空行，则继续
		if len(line) == 0 {
			continue
		}
		// 3.1 如果是注释，则跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") { // HasPrefix是否以xx为前缀开头,ini文件注释以“;”或“#”开头
			continue
		}
		if strings.HasPrefix(line, "[") {
			// 3.2 如果开头是[，而且是以]结尾的，且[]不为空，则被认定是section（小节）
			if line[0] == '[' && line[len(line)-1] == ']' && len(line) > 2 {
				// 将section 加入到section切片变量里
				section = line[1 : len(line)-1]
			}
			//根据字符串section去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ { // Elem()方法:主要用来获取指针类型（只能使用在数组、chan、map、指针、切片几个类型上）的类型对象
				field := t.Elem().Field(i)
				if section == field.Tag.Get("ini") {
					structName = field.Name
				}
			}
		} else {
			// 4. 拆分键值对
			// 4.1 以等号分割，左边为key，右边为value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") { // Index：子串sep在字符串s中第一次出现的位置，不存在则返回-1。
				zap.S().Debugf("line:%d syntax error", idx+1)
				return
			}
			// 获取到line的key和value
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])     // user
			value := strings.TrimSpace(line[index+1:]) // mysql

			// 4.2 将ini文件中的section和结构体的名字对应
			v := reflect.ValueOf(data)                 // interface的值,此时为空
			sValue := v.Elem().FieldByName(structName) // 结构体的值 { 10.0.0.1 mysql abc123 0}
			sType := sValue.Type()                     // 结构体的类型 main.MysqlConfig

			// 判断结构体的字段是否是struct类型
			if sType.Kind() != reflect.Struct {
				zap.S().DPanic("should be struct")
				return err
			}
			// 根据structName 去data里面把嵌套的结构体取出
			// 声明字段值
			var fname string                  // 结构体的字段名
			var fieldType reflect.StructField // 结构体字段类型
			for i := 0; i < sValue.NumField(); i++ {
				fieldName := sType.Field(i)
				fieldType = fieldName
				//  遍历结构体的每一个字段，判断这个tag是不是等于key
				if fieldName.Tag.Get("ini") == key {
					// 找到对应的字段
					fname = fieldName.Name
					break
				}
			}
			// 4.3如果key == tag，and value的类型等于结构体定义的类型，则赋值
			fieldObj := sValue.FieldByName(fname)
			switch fieldType.Type.Kind() { // 判断ini文件中value的种类，字串/数字/布尔
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				fieldObj.SetInt(valueInt)

			case reflect.Bool:
				var valueBool bool
				valueBool, err := strconv.ParseBool(value)
				if err != nil {
					return err
				}
				fieldObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				fieldObj.SetFloat(valueFloat)
			}
		}
	}
	return
}
func main() {
	InitLogger()

	var config Config
	err := loadIni("D:/MyFirstProject/Day5/parseini/config.ini", &config)
	if err != nil {
		zap.S().DPanic(err.Error())
	}
	zap.S().Debugf("ip = %v, password = %v, port = %v, user = %v",
		config.MysqlConfig.Ip,
		config.MysqlConfig.Password,
		config.MysqlConfig.Port,
		config.MysqlConfig.User,
	)
}

func InitLogger() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "parse_ini.log"}
	config.DisableCaller = true
	config.DisableStacktrace = true

	_logger, err := config.Build()
	if err != nil {
		zap.S().DPanicf("日志初始化失败，err:%s", err.Error())
	}
	zap.ReplaceGlobals(_logger)
}
