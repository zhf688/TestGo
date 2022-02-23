package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//定义链接客户端
var client *mongo.Client

//初始化链接
func initDb() {
	//创建options
	uri := options.Client().ApplyURI("mongodb://192.168.115.128:27017")
	//使用uri创建链接
	client, err := mongo.Connect(context.TODO(), uri)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("pingerr:%v\n", err)
	} else {
		fmt.Printf("mongodb连接成功！！！\n")
	}
}

func main() {
	initDb()
}
