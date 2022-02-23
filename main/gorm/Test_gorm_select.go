package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db1 *gorm.DB

//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}

func initdb1() {
	//连接mysql，创建mysql的dns连接串
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"
	//gorm打开数据库连接，返回db对象
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	db1 = d
	fmt.Printf("db1:%v\n", db1)
}

//First、Last、Take、Find、Where
func select1() {
	var pr Product
	// 获取第一条记录（主键升序）
	//db1.First(&pr)
	//// SELECT * FROM users ORDER BY id LIMIT 1;
	//fmt.Printf("pr:%v\n",pr)
	//
	//// 获取一条记录，没有指定排序字段
	//db1.Take(&pr)
	//// SELECT * FROM users LIMIT 1;
	//fmt.Printf("pr:%v\n",pr)

	// 获取最后一条记录（主键降序）
	//db1.Last(&pr)
	//// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	//fmt.Printf("pr:%v\n", pr)

	result := []map[string]interface{}{}
	db1.Model(&Product{}).First(&result)
	fmt.Printf("result:%v\n", result)
	//result再次使用take获取时，有两条记录
	db1.Table("products").Take(&result)
	fmt.Printf("result:%v\n", result)

	//db1.Where(&Product{Code: "1005"}).Find(&pr)
	//fmt.Printf("pr-where:%v\n", pr)

	db1.Find(&pr, "Code=?", "1006")
	fmt.Printf("pr-find:%v\n", pr)
}
func main() {
	initdb1()
	select1()
}
