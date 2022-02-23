package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//结构体属性名首字母大写
type Person struct {
	Name    string
	Age     int
	Email   string
	Perents []string
}

//对象转json
func test1() {
	p := Person{
		Name:  "zhang",
		Age:   20,
		Email: "zhang@gmail.com",
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("p:%v\n", string(b))
}

//json转对象
func test2() {
	//多字符串引用，需要使用"`"符号
	b := []byte(`{"Name":"zhang","Age":20,"Email":"zhang@gmail.com"}`)
	var p Person
	json.Unmarshal(b, &p)
	fmt.Printf("p:%v\n", p)
}

//嵌套json数据转对象
func test3() {
	b := []byte(`{"Name":"zhang","Age":20,"Email":"zhang@gmail.com","perents":["big zhang","keke"]}`)
	var p map[string]interface{}
	json.Unmarshal(b, &p)
	fmt.Printf("p:%v\n", p)

}
func test4() {
	p := Person{
		Name:    "zhang",
		Age:     20,
		Email:   "zhang@gmail.com",
		Perents: []string{"big zhang", "keke"},
	}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("p:%v\n", string(b))
}

//打开json文件读取数据,使用json库中的decoder读取文件流数据
func test5() {
	f, _ := os.Open("a.json")
	defer f.Close()
	d := json.NewDecoder(f)
	var m map[string]interface{}
	d.Decode(&m)
	fmt.Printf("m%v\n", m)
}

//把json对象信息写入文件,使用json.NewEncoder()
func test6() {
	p := Person{
		Name:    "sun",
		Age:     20,
		Email:   "sun@gmail.com",
		Perents: []string{"big sun", "keke"},
	}
	f, _ := os.OpenFile("a.json", os.O_RDWR|os.O_APPEND, 0777)
	e := json.NewEncoder(f)
	e.Encode(p)
}
func main() {
	//test2()
	//test4()
	//test5()
	test6()
}
