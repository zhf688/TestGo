package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("timer...%v\n", time.Now())
	//新建timer变量，类型为Timer，Timer结构体中，有C通道
	timer := time.NewTimer(time.Second * 2)
	t1 := <-timer.C //利用C通道可实现时间等待
	fmt.Printf("t1:%v\n", t1)

	fmt.Printf("timer1...begin..%v\n", time.Now())
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C //不用变量，也可实现等待
	fmt.Printf("timer1...end..%v\n", time.Now())

	time.Sleep(time.Second * 2)
	fmt.Printf("time.sleep...%v\n", time.Now())

	<-time.After(time.Second * 2) //使用time.After函数，实现时间等待，返回channel类型
	fmt.Printf("time...after..2s%v\n", time.Now())

	//停止定时器的使用
	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Printf("timer2...%v\n", time.Now())
	}()
	s := timer2.Stop() //停止定时器后，原定义的定时器不再执行
	if s {
		fmt.Printf("timer2...stop,%v\n", time.Now())
	}

	//重置定时器的使用
	fmt.Printf("timereset...before  %v\n", time.Now())
	timer3 := time.NewTimer(time.Second * 5)
	timer3.Reset(time.Second * 2)
	<-timer3.C
	fmt.Printf("timereset...end  %v\n", time.Now())
}
