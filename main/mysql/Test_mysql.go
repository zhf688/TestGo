package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

////自定义数据库连接函数
var db *sql.DB

func initDb1() (err error) {
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"

	//不能使用:=，定义了全局变量db，要赋值之后在main中使用
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	/*
		db, err := sql.Open("mysql", "root:123456@/mysql-test")
		if err != nil {
			panic(err)
		}
		print(db)
		// See "Important settings" section.
		db.SetConnMaxLifetime(time.Minute * 3) //最大连接时长
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	*/
	err := initDb1()
	if err != nil {
		fmt.Printf("数据库链接失败\n")
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("数据库链接成功")
	}
}
