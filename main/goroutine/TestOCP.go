package main

import "fmt"

//OCP设计原则，对扩展是开放的，对修改是关闭的
//定义通用接口
type Pet interface {
	eat()
	sleep()
}

//定义类型
type Dog struct {
}

//类型实现接口方法
func (dog Dog) eat() {
	fmt.Printf("dog eat...\n")
}
func (dog Dog) sleep() {
	fmt.Printf("dog sleep...\n")
}

type Cat struct {
}

func (cat Cat) eat() {
	fmt.Printf("cat eat...\n")
}
func (cat Cat) sleep() {
	fmt.Printf("cat sleep...\n")
}

type Person struct {
}

//传入通用接口类型，调用接口中的方法
func (person Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main7() {
	dog := Dog{}
	cat := Cat{}
	person := Person{}
	//传入实现了接口的类型，表示向上类型转换
	person.care(dog)
	person.care(cat)
}
