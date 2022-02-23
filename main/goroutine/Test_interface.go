package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{} //因为实现Speak接口的是*Student类型，所以必须使用指针类型赋值变量，不能使用Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
