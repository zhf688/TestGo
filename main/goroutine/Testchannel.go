package main

import (
	"fmt"
	"math/rand"
	"time"
)

//创建一个通道，通道用于协程间通信
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}

func main11() {
	defer close(values)
	go send()
	fmt.Printf("wait...\n")
	value := <-values
	fmt.Printf("receive:%v\n", value)
	fmt.Printf("end...")
}
