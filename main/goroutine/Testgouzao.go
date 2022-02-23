package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

//定义构造函数，传入初始化变量，返回类型对象及错误对象
func Newperson(name string, age int) (*Person1, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age不能小于0")
	}
	return &Person1{name: name, age: age}, nil
}

func main9() {
	//创建对象，赋值变量必须被使用，否则会报错
	person, err := Newperson("", 20)
	if err == nil {
		fmt.Printf("person:%v\n", *person)
	} else {
		fmt.Printf("error:%v\n", err)
	}
}
