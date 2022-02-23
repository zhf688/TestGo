package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client2 *mongo.Client

func initDB() {
	//创建options
	uri := options.Client().ApplyURI("mongodb://192.168.115.128:27017")
	//连接mongodb
	connect, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		log.Fatal(err)
	}
	err1 := connect.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("connect-err%v\n", err1)
	} else {
		fmt.Printf("mongodb连接成功！\n")
	}
	client2 = connect
}

//查找文档记录
func find() {
	initDB()
	ctx := context.TODO()
	defer client2.Disconnect(ctx)                                   //关闭数据库连接
	collection := client2.Database("mymongo").Collection("student") //获取集合
	//查找文档，返回类型为游标cursor，bson.D{}是过滤条件，空则代表无过滤，可以多个条件
	cursor, err := collection.Find(ctx, bson.D{{"name", "zhanghf"}})
	defer cursor.Close(ctx) //线程结束后关闭游标
	if err != nil {
		log.Fatal(err)
	} else {
		//游标遍历
		for cursor.Next(ctx) {
			var result bson.D
			cursor.Decode(&result)
			fmt.Printf("result:%v\n", result)
			fmt.Printf("result-map%v\n", result.Map())
		}
	}
}
func main() {
	find()
}
