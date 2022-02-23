package main

import (
	"fmt"
	"time"
)

var c = make(chan int)

func main() {
	//开启一个协程，执行匿名函数，写入数据到channel
	go func() {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c) //如果未关闭通道，遍历次数超过通道长度，会造成死锁：fatal error: all goroutines are asleep - deadlock!
	}()
	time.Sleep(time.Millisecond * 1000)
	//遍历方式一：
	//for i := 0; i < 3; i++ {
	//	r := <-c //读取通道内容
	//	fmt.Printf("r:%v\n",r)
	//}

	//遍历方式2
	for v := range c {
		fmt.Printf("v:%v\n", v)
	}
}
