package main

import (
	"fmt"
	"time"
)

//Ticker，间隔定时执行
func main() {
	ticker := time.NewTicker(time.Second) //每间隔1秒执行一次
	count := 0
	for _ = range ticker.C {
		fmt.Printf("ticker...\n")
		count++
		if count >= 5 {
			ticker.Stop() //满足条件，停止定时器，跳出循环
			break
		}
	}
	time.Sleep(time.Second)
	ticker1 := time.NewTicker(time.Second) //每间隔1秒执行一次，不重新定义而使用ticker的话，会造成死锁
	chanint1 := make(chan int)
	go func() { //通道写数据，需要另外开辟一个协程
		for _ = range ticker1.C {
			//每间隔1秒往通道写数据
			select {
			case chanint1 <- 1:
			case chanint1 <- 2:
			case chanint1 <- 3:
			}
		}
	}()
	defer close(chanint1)
	sum := 0
	for v := range chanint1 {
		fmt.Printf("接收:%v\n", v)
		sum += v
		if sum >= 10 {
			break
		}
	}

}
