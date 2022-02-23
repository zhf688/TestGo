package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //驱动必须导入，否则链接数据库报错
)

var db2 *sql.DB

type User struct {
	Id    int
	Name  string
	Age   int
	Email string
}

func initDb2() (err error) {
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"

	//不能使用:=，定义了全局变量db，要赋值之后在main中使用
	db2, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db2.Ping()
	if err != nil {
		return err
	}
	return nil
}

//查询一行，使用db.QueryRow方法
func queryonerow() {
	s := "select * from myuser where id = ?"
	row := db2.QueryRow(s, 2)
	var u User
	err := row.Scan(&u.Id, &u.Name, &u.Age, &u.Email)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("u:%v\n", u)
	}
}

//查询多行,db.Query()
func queryRows() {
	s := "select * from myuser"
	rows, err := db2.Query(s)
	defer rows.Close()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		for rows.Next() {
			var u User
			err := rows.Scan(&u.Id, &u.Name, &u.Age, &u.Email)
			if err != nil {
				fmt.Printf("for-err:%v\n", err)
			} else {
				fmt.Printf("u:%v\n", u)
			}
		}
	}
}

func main() {
	err := initDb2()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("连接成功\n")
	}
	queryonerow()
	queryRows()
}
