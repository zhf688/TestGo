package main

import "fmt"

//append追加
func testapppend() {
	s := []int{1, 2, 3}
	i := append(s, 100)
	fmt.Printf("i%v\n", i)
	s1 := []int{4, 5, 6}
	i = append(s, s1...) //把一个切片追加到另一个切片中时,后边要加上"..."
	fmt.Printf("i:%v\n", i)
}

//len()长度
func testlen() {
	s := "hello"
	i := len(s)
	fmt.Printf("i:%v\n", i)
}

//panic抛出异常
func testpanic() {
	defer fmt.Printf("defer结束,panic后可执行\n")
	panic("panic\n")
	fmt.Printf("panic后边执行")
}
func main() {
	//testlen()
	testpanic()
	fmt.Printf("主函数执行\n")
}
