package main

import (
	"fmt"
	"os"
)

//文件及目录操作

//创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("f:%v\n", f)
}

//创建目录
func crearDir() {
	err := os.Mkdir("a", os.ModePerm) //os.ModePerm 表示创建目录权限为777
	if err != nil {
		fmt.Printf("err:%v\n", err) //err:mkdir a: 当文件已存在时，无法创建该文件
	}
	//创建多层级目录
	os.MkdirAll("a/b/c", os.ModePerm)
}

//删除目录或文件
func remove() {
	err := os.Remove("a.txt") //删除文件
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	err1 := os.RemoveAll("a/b") //删除b目录及目录下的文件
	if err1 != nil {
		fmt.Printf("err1:%v\n", err1)
	}
}

//获取工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Printf("pwd:%v\n", dir)

	//切换目录
	os.Chdir("G:/")
	dir, err = os.Getwd()
	fmt.Printf("pwd:%v\n", dir)
}

//写文件
func writeFile() {
	os.WriteFile("a.txt", []byte("hello world|hello go"), os.ModePerm)
}

//读文件
func readFile() {
	data, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Printf("err:%v\n")
	}
	os.Stdout.Write(data)
}

func main() {

	//createFile()
	//crearDir()
	//remove()
	//getWd()
	writeFile()
	readFile()
}
