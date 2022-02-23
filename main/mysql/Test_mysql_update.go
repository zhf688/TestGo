package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

////自定义数据库连接函数
var db3 *sql.DB

func initDb3() (err error) {
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"

	//不能使用:=，定义了全局变量db，要赋值之后在main中使用
	db3, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db3.Ping()
	if err != nil {
		return err
	}
	return nil
}

//更新操作db.Exec()
func update() {
	s := "update myuser set username=?,age=? where id=?"
	exec, err := db3.Exec(s, "big tom", 22, 1)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		affected, _ := exec.RowsAffected()
		fmt.Printf("affect:%v\n", affected)
	}
}

//删除操作，db.Exec()
func delete() {
	s := "delete from myuser where id = ?"
	exec, err := db3.Exec(s, 2)
	if err != nil {
		fmt.Printf("delete-err:%v\n", err)
	} else {
		affected, _ := exec.RowsAffected() //影响行数
		fmt.Printf("aff:%v\n", affected)
	}
}

func main() {
	err := initDb3()
	if err != nil {
		fmt.Printf("dberr:%v\n", err)
	}
	//update()
	delete()
}
