package main

import "fmt"

type person struct {
	name string
}

func test1() {

	p1 := person{
		name: "tom",
	}
	p2 := &person{
		name: "tom",
	}
	fmt.Printf("p1 %T\n", p1)
	fmt.Printf("p2 %T\n", p2)
}

func showperson1(per person) {
	per.name = "tom...."
}
func showperson2(per *person) {
	//(*per).name自动解引用
	per.name = "tom..."
}

//方法，名称前(per person)是receiver，表示属于那个类型的方法
func (per person) showperson3() {
	per.name = "tom123...."
}

func (per *person) showperson4() {
	//(*per).name自动解引用
	per.name = "tom123..."
}

func main5() {
	//test1()
	p1 := person{
		name: "tom",
	}
	//函数调用
	showperson1(p1)
	fmt.Printf("p1%v\n", p1)
	p2 := &person{
		name: "tom",
	}
	showperson2(p2)
	fmt.Printf("p1%v\n", *p2)

	//方法调用
	p1.showperson3()
	fmt.Printf("p1..%v\n", p1)
	p2.showperson4()
	fmt.Printf("p2..%v\n", *p2)
}
