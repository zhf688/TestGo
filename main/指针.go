package main

import (
	"fmt"
)

const MAX = 3

func main3() {
	number := [MAX]int{1, 2, 3}
	var pa [MAX]*int
	var pt *int
	for i := 0; i < MAX; i++ {
		pa[i] = &number[i]
	}
	fmt.Printf("指针数组%d\n", &pa)
	for i, x := range pa {
		fmt.Printf("循环遍历索引%d，数组值为：%d,地址为：%d\n", i, *x, x)
	}
	fmt.Printf("空指针数组：%x", pt)
}
