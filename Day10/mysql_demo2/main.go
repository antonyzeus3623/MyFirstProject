package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入mysql驱动包 init()
)

//Go连接MySQL示例
var db *sql.DB //从数据库的连接池拿取一个连接

func initDB() (err error) {
	//数据库信息
	//用户名:密码@tcp(ip:端口)/数据库名
	dsn := "root:123456@tcp(127.0.0.1:3306)/sql_test"
	//连接数据库
	db, err = sql.Open("mysql", dsn) //不会校验用户名和密码是否正确
	if err != nil {                  //当dsn格式不正确的时候会报错
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		return
	}
	// defer db.Close()
	return
}

type user struct {
	id   int
	name string
	age  int
}

//查询单个记录
func queryRowDemo(id int) {
	var u1 user
	//1. 写查询单条记录的SQL语句
	sqlStr := `select id, name, age from user where id=?;` // 其中?表示占位符
	//2. 执行并拿到结果
	//从连接池里拿一个连接去数据库查询单条数据，必须对objRow对象调用Scan方法，因为该方法会释放数据库连接,如果不释放会被一直占用，导致程序卡住
	db.QueryRow(sqlStr, 1002).Scan(&u1.id, &u1.name, &u1.age)
	//3. 打印结果
	fmt.Printf("u1:%#v\n", u1)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功！")
	queryRowDemo(1001)
}
