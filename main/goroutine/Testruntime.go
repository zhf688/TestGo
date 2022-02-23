package main

import (
	"fmt"
	"runtime" //定义一些协程管理相关的api
	"time"
)

func showruntime() {
	for i := 0; i < 5; i++ {
		fmt.Printf("i:%v,java\n", i)
		//time.Sleep(time.Millisecond*100)
		if i >= 3 {
			runtime.Goexit() //直接退出协程
		}
	}
}
func showruntime1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("i1...:%v,python\n", i)
		//time.Sleep(time.Millisecond*100)
	}
}
func main() {
	fmt.Printf("runtime:cpunumber:%v\n", runtime.NumCPU())
	go showruntime()
	go showruntime1()
	for i := 0; i < 2; i++ {
		runtime.Gosched() //让出cpu，让其他go协程先执行
		fmt.Printf("golang\n")
	}
	time.Sleep(time.Second * 2)
}
