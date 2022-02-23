package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	//反引号
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email"`
}

func test1() {
	p := Person{
		Name:  "Tom",
		Age:   21,
		Email: "t@gmail.com",
	}
	//m, _ := xml.Marshal(p) //xml转化为byte切片，打印时强转为string
	m, _ := xml.MarshalIndent(p, "", "  ") //格式化xml，第二个参数为前缀，第三个参数为缩进
	fmt.Printf("%v\n", string(m))
}

//xml字符串转对象
func test2() {
	s := `
<person>
  <name>Tom</name>
  <age>21</age>
  <email>t@gmail.com</email>
</person>
`
	b := []byte(s)
	var p Person
	xml.Unmarshal(b, &p)
	fmt.Printf("p:%v\n", p)
}

//xml文件解析成对象
func test3() {
	//f, _ := ioutil.ReadFile("a.xml") //解析文件可以用ioutil的方法，读取到字节缓冲区
	//var p Person
	//xml.Unmarshal(f, &p) //使用xml的字节切片转对象
	//fmt.Printf("p3:%v\n", p)
	f1, _ := os.Open("a.xml")
	defer f1.Close()
	var p Person
	d := xml.NewDecoder(f1)
	d.Decode(&p)
	fmt.Printf("p3:%v\n", p)
}

//对象写入xml文件
func test4() {
	p := Person{
		Name:  "sun",
		Age:   20,
		Email: "sun@gmail.com",
	}
	f, _ := os.OpenFile("a.xml", os.O_RDWR|os.O_APPEND, 0777)
	e := xml.NewEncoder(f)
	e.Encode(p)
}
func main() {
	//test2()
	test3()
	//test4()
}
