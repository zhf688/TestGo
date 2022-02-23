package main

import (
	"fmt"
	"time"
)

var chanint = make(chan int)
var chanstr = make(chan string)

func main() {
	go func() {
		chanint <- 100
		chanstr <- "hello"
	}()
	for {
		select { //select类似于switch语句，用于异步处理io操作，其中的case用于channel的读写操作
		case r := <-chanint:
			fmt.Printf("chanint:%v\n", r)
		case r := <-chanstr:
			fmt.Printf("chanstr:%v\n", r)
		default:
			fmt.Printf("default...\n")
		}
		time.Sleep(time.Second)
	}
}
