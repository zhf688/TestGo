package main

import (
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

func main1() {
	fmt.Println("hello")
	var a int = 100
	var b int = 200
	b, a = a, b
	fmt.Println(a, b)
	c, _ := getData()
	_, d := getData()
	fmt.Println(c, d)

	name := "tom"
	age := "20"
	fmt.Printf("名字是：%s\n", name)
	fmt.Printf("a：%d\t,长度：%v\n", a, unsafe.Sizeof(a))
	fmt.Printf("name is %v\n", name)

	var buffer bytes.Buffer
	buffer.WriteString(name)
	buffer.WriteString(age)
	fmt.Printf("buffer %s,类型为：%v\n", buffer, reflect.TypeOf(buffer))
	fmt.Printf("======%#v\n", name)

	//var s2 = make([]int,2)
	var s1 = []int{}
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)
	s2 := s1[:2]
	fmt.Println(s2)
	//切片的删除,使用append函数，append(slice[:index],slice[index+1:]...)
	s3 := append(s1[:2], s1[3:]...)
	fmt.Println(s3)

}

func getData() (int, int) {
	return 100, 200
}
