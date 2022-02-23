package main

import (
	"fmt"
	"strconv"
)

//golang中没有继承的概念，但是可以通过结构体嵌套实现继承的功能

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() string {
	return "eat..."
}
func (a Animal) sleep() string {
	return "sleep..."
}

//同一个包下，不能有相同的类型结构体（TestOCP中有Dog和Cat）
type Dog1 struct {
	a     Animal //变量a可以省略，初始化时直接写Animal类型,顺序初始化变量名都省略
	color string
}
type Cat1 struct {
	a      Animal
	weight int
}

func main8() {
	dog := Dog1{
		a: Animal{
			name: "huahua",
			age:  2,
		},
		color: "red",
	}
	cat := Cat1{
		a: Animal{
			name: "xiangxiang",
			age:  3,
		},
		weight: 30,
	}
	fmt.Println(dog.a.name + " " + dog.color + " " + dog.a.eat())
	//go语言中数字转字符串方法：strconv.Itoa(a int)
	//strconv.Atoi()：字符串转整型 返回2个值，第一个是值，第二个是错误，下面没有做处理
	fmt.Println(cat.a.name + " " + strconv.Itoa(cat.weight) + " " + cat.a.sleep())
}
