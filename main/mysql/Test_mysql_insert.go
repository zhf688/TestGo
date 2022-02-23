package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //驱动必须导入，否则链接数据库报错
)

var db1 *sql.DB

func initDb() (err error) {
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"

	//不能使用:=，定义了全局变量db，要赋值之后在main中使用
	db1, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db1.Ping()
	if err != nil {
		return err
	}
	return nil
}

//插入数据
func insert() {
	sqlstr := "insert into myuser(username,age,email) values(?,?,?)"
	r, err := db1.Exec(sqlstr, "zhang", 20, "zhang@qq.com")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		id, _ := r.LastInsertId() //获取最后的id
		fmt.Printf("id:%v\n", id)
	}
}
func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("数据库链接失败\n")
		fmt.Printf("err:%v\n", err)
	}
	insert()
}
