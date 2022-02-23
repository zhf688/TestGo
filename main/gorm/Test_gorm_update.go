package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db2 *gorm.DB

//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}

func initdb2() {
	//连接mysql，创建mysql的dns连接串
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"
	//gorm打开数据库连接，返回db对象
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	db2 = d
	fmt.Printf("db2:%v\n", db2)
}

//更新操作可查找官方文档
func update2() {
	db2.Model(&Product{}).Where("Code=?", "1006").Update("price", 606)
}
func main() {
	initdb2()
	update2()
}
