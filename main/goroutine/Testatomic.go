package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

//var i1 int = 100
//var lock1 sync.Mutex //通过加锁方式保证线程同步安全
//
//func add2() {
//	lock1.Lock()
//	i1++
//	lock1.Unlock()
//}
//func sub2() {
//	lock1.Lock()
//	i1--
//	lock1.Unlock()
//}

//还可以通过使用atomic原子变量来实现线程同步安全
var i1 int32 = 100

//atomic内部使用compare and swap 实现
func add2() {
	atomic.AddInt32(&i1, 1)
}
func sub2() {
	atomic.AddInt32(&i1, -1)
}

//load载入，即为读
func atomicload() {
	var m int32 = 99
	t := atomic.LoadInt32(&m) //read
	fmt.Printf("t:%v\n", t)
}

//存储，即为写
func atomicstore() {
	var m int32 = 99
	atomic.StoreInt32(&m, 1) //write
	fmt.Printf("m:%v\n", m)
}

//cas，比较替换
func atomiccas() {
	var m int32 = 99
	b := atomic.CompareAndSwapInt32(&m, 99, 100) //write
	fmt.Printf("b:%v\n", b)
	fmt.Printf("m..:%v\n", m)
}
func main() {
	for i := 0; i < 100; i++ {
		go add2()
		go sub2()
	}
	time.Sleep(time.Second)
	fmt.Printf("i:%v\n", i1)
	atomicload()
	atomicstore()
	atomiccas()
}
