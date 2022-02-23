package main

import "fmt"

//闭包写法，变量x在函数调用期间一直有效
func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

func main4() {
	f := add() //函数f就是一个闭包函数，变量x在f调用过程中一直有效
	r := f(10)
	fmt.Printf("r:%v\n", r)
	r = f(20)
	fmt.Printf("r:%v\n", r)

	add, sub := cal(100) //复杂闭包函数，将两个函数作为返回值进行计算，base变量在函数调用期间一直有效，重新调用时，变量重置
	r1 := add(1)
	fmt.Printf("r1:%v\n", r1)
	r1 = sub(2)
	fmt.Printf("r1:%v\n", r1)

}
