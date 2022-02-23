package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client

func initMondb() {
	uri := options.Client().ApplyURI("mongodb://192.168.115.128:27017")
	connect, _ := mongo.Connect(context.TODO(), uri)
	cli = connect
}

//修改
func update() {
	initMondb()
	ctx := context.TODO()
	defer cli.Disconnect(ctx)
	collection := cli.Database("mymongo").Collection("student")
	//设置要修改的数据，update变量
	update := bson.D{{"$set", bson.D{{"name", "zhanghf"}, {"age", 27}}}}
	//按条件查找要修改的数据，传入update变量
	many, err := collection.UpdateMany(ctx, bson.D{{"name", "zhang"}}, update)

	if err != nil {
		fmt.Printf("update-err:%v\n", err)
	} else {
		fmt.Printf("update-many:%v\n", many.ModifiedCount)
	}

}

//删除一条数据
func delete() {
	initMondb()
	ctx := context.TODO()
	defer cli.Disconnect(ctx)
	collection := cli.Database("mymongo").Collection("student")

	//删除一条数据
	one, err := collection.DeleteOne(ctx, bson.D{{"name", "li"}})

	if err != nil {
		fmt.Printf("update-err:%v\n", err)
	} else {
		fmt.Printf("delete-one:%v\n", one.DeletedCount)
	}
}
func main() {
	update()
	delete()
}
