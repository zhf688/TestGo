package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

//定义链接客户端
var client1 *mongo.Client

//初始化链接
func initDb1() {
	//创建options
	uri := options.Client().ApplyURI("mongodb://192.168.115.128:27017")
	//使用uri创建链接,此时的client1注意要用定义的全局变量，不能重新初始化赋值
	client1, _ = mongo.Connect(context.TODO(), uri)

	//检查连接
	err := client1.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("pingerr:%v\n", err)
	} else {
		fmt.Printf("mongodb连接成功！！！\n")
	}
}

//插入一条记录
func insert() {
	initDb1()
	s := Student{
		Name: "zhang",
		Age:  20,
	}
	col := client1.Database("mymongo").Collection("student") //获取集合
	one, err := col.InsertOne(context.TODO(), s)             //插入一条记录
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		println(one.InsertedID)
	}
}

//插入多条记录,传入切片
func insertMany(students []interface{}) {
	initDb1()
	collection := client1.Database("mymongo").Collection("student") //获取数据库集合
	many, err := collection.InsertMany(context.TODO(), students)    //插入传入的切片
	if err != nil {
		fmt.Printf("insertmany-err%v\n", err)
	} else {
		fmt.Printf("many.InsertedIDs:%v\n", many.InsertedIDs)
	}
}
func main() {
	//insert()
	s1 := Student{Name: "sun", Age: 19}
	s2 := Student{Name: "li", Age: 18}
	stus := []interface{}{s1, s2}
	insertMany(stus)
}
