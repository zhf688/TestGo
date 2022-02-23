package main

import (
	"fmt"
	"sync"
	"time"
)

var i int = 100
var wg sync.WaitGroup
var lock sync.Mutex //互斥锁

func add1() {
	defer wg.Done() //结束减一
	lock.Lock()     //加锁实现线程同步，执行期间其他线程无法进入
	i += 1
	fmt.Printf("i++:%v\n", i)
	time.Sleep(time.Millisecond * 10) //设置等待时长，模拟线程不同步
	lock.Unlock()                     //执行完成后要解锁
}
func sub1() {
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	defer wg.Done()
	i -= 1
	fmt.Printf("i--:%v\n", i)
	lock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1) //启动加一
		go add1()
		wg.Add(1)
		go sub1()
	}
	wg.Wait() //等待所有goroutine结束
	fmt.Printf("i:%v\n", i)
}
