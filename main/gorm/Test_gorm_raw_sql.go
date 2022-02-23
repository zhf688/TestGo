package main

//原生sql,使用db.Raw()函数
import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbr *gorm.DB

//数据库对应的结构体
type Prod struct {
	Code  string
	Price uint
}

//数据库对应的结构体
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func initdbr() {
	//连接mysql，创建mysql的dns连接串
	dsn := "root:123456@tcp(192.168.115.128:3306)/mysql?charset=utf8mb4&parseTime=true"
	//gorm打开数据库连接，返回db对象
	d, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("failed to connect database")
	}
	dbr = d
}

//原生sql查询
func rawSelect() {
	var pro Prod
	dbr.Raw("select * from products where code = ?", "1006").Scan(&pro)
	fmt.Printf("pro:%v\n", pro)
	var p Prod
	dbr.Raw("select code,price from products where code = ?", "1005").Scan(&p)
	fmt.Printf("pro-1:%v\n", p)
}

//原生sql执行更新删除操作
func rawExec() {
	dbr.Exec("update products set price = ? where code = ?", 555, "1005")
}

//命名参数
func testRaw() {
	var p Product
	dbr.Where("price = @proprice", sql.Named("proprice", 300)).Find(&p)
	fmt.Printf("test-raw:%v\n", p)
}

//row和rows
func testRow() {
	var price int
	row := dbr.Table("products").Where("code=?", "1005").Select("price").Row()
	row.Scan(&price)
	fmt.Printf("price:%v\n", price)

	rows, _ := dbr.Model(&Product{}).Where("price > ?", 200).Select("code").Rows()
	for rows.Next() {
		var code string
		rows.Scan(&code)
		fmt.Printf("code:%v\n", code)
	}
}
func main() {
	initdbr()
	//rawSelect()
	//rawExec()
	//testRaw()
	testRow()
}
