package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showmsgwait(i int) {

	defer wp.Done() //goroutine结束就-1
	fmt.Printf("i:%v\n", i)
}
func main12() {

	for i := 0; i < 10; i++ {
		wp.Add(1) //goroutine启动就+1
		go showmsgwait(i)
	}
	wp.Wait() //等待所有的goroutine结束
	fmt.Printf("end...")
}
