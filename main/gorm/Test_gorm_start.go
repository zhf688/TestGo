package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//数据库对应的结构体
//type Product struct {
//	gorm.Model
//	Code  string
//	Price uint
//}

func initdb() {
	//连接mysql，创建mysql的dns连接串
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"
	//gorm打开数据库连接，返回db对象
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	db = d
}

//创建表
func createtable() {
	//迁移schema，创建结构体绑定的表
	db.AutoMigrate(&Product{})
}

//插入数据
func insert() {
	p := Product{Code: "1002", Price: 200}
	create := db.Create(&p)
	fmt.Printf("insert:%v\n", create.RowsAffected)
}

//选择字段插入数据
func insert2() {
	p := Product{Code: "1004", Price: 400}
	db.Select("Code", "CreatedAt").Create(&p)
}

//批量插入数据
func insertmany() {
	var ps = []Product{{Code: "1005", Price: 500}, {Code: "1006", Price: 600}}
	db.Create(&ps)
}

//查找数据
func find() {
	var p Product
	//根据id查询
	db.First(&p, 1)
	fmt.Printf("findone:%v\n", p)
	//根据条件查询
	db.First(&p, "code = ?", "1002")
	fmt.Printf("findbycode:%v\n", p)
}

//更新数据
func update() {
	var p Product
	db.First(&p, 1)
	//更新一个字段
	tx := db.Model(&p).Update("code", "abc")
	fmt.Printf("update-one:%v\n", tx.RowsAffected)
	//更新多个字段
	tx = db.Model(&p).Updates(Product{Code: "aaa", Price: 80})
	fmt.Printf("update-two:%v\n", tx.RowsAffected)
	//更新多个字段，另一种写法
	tx = db.Model(&p).Updates(map[string]interface{}{"Price": 233, "Code": "bbb"})
	fmt.Printf("update-two!!:%v\n", tx.RowsAffected)
}

//删除记录,此删除操作是软删除，只会更新数据库删除时间字段，并不会删除数据
func delete() {
	var p Product
	db.First(&p, 1)
	db.Delete(&p)
}
func main() {
	initdb()
	//createtable()
	//insert()
	find()
	//update()
	//delete()
	//insert2()
	//insertmany()
}
