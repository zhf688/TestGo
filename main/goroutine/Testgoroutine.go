package main

import (
	"fmt"
	"time"
)

func showmsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg:%v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}
func main10() {
	go showmsg("java")                  //go关键字相当于创建了一个协程，类似于main
	go showmsg("golang")                //在main协程中执行，如果前边全部加上go，程序会没有任何输出
	time.Sleep(time.Millisecond * 2000) //程序等待时间，go并发协程会执行，此操作是在main协程中
}
