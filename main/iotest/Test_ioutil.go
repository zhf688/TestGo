package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readall() {
	//r := strings.NewReader("hello zhang")

	r, _ := os.Open("a.txt") //除了可以使用newreader，也可以打开文件读取
	defer r.Close()
	all, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("string(all):%v\n", string(all))
	}
}

//读取目录及文件列表
func testreaddir() {
	dir, _ := ioutil.ReadDir(".")
	for _, v := range dir {
		fmt.Printf("v.name:%v\n", v.Name())
	}
}

//读取文件内容
func testReadFile() {
	file, _ := ioutil.ReadFile("a.txt")
	fmt.Printf("file:%v\n", string(file))
}

func main() {
	readall()
	testreaddir()
	testReadFile()
}
