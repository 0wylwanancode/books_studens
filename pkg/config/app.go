package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
var db *sql.DB

//链接数据库
func init() {
	//链接mysql的语句
	dsn := "root:wylwanan123..@tcp(127.0.0.1:3306)/social?charset=utf8"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Open() 打开数据库失败")
		log.Fatal(err)
	}
	if err = d.Ping(); err != nil {
		fmt.Println("Ping() 测试数据库失败")
		log.Fatal(err)
	}
	fmt.Println("链接数据库成功")
	db = d
}
func GetDB() *sql.DB {
	return db
}
