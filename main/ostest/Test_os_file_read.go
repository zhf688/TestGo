package main

import (
	"fmt"
	"io"
	"os"
)

//打开关闭文件
func openclose() {
	open, err := os.Open("a1.txt") //调用的是openFile方法
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("filename:%v\n", open.Name())
		open.Close()
	}
	file, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755) //os.O_RDWR|os.O_CREATE,读写，如果没有文件则创建
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("file1name:%v\n", file.Name())
		file.Close()
	}
}

//创建文件和创建临时文件
func createFileTmp() {
	f, _ := os.Create("a2.txt") //创建文件，等价于os.OpenFile("a2.txt",os.O_CREATE|os.O_RDWR|os.O_TRUNC,0666)
	fmt.Printf("f:%v\n", f.Name())
	f2, _ := os.CreateTemp("", "tmp") //第一个参数表示目录，为空则创建到默认目录，第二个参数为临时文件前缀
	fmt.Printf("f2:%v\n", f2.Name())
}

//循环读取文件
func readOps() {
	f3, _ := os.Open("a.txt")
	for {
		buf := make([]byte, 3) //创建buf缓冲区
		n, err := f3.Read(buf) //读取文件,n为int类型,表示每次读取多少个字节
		if err == io.EOF {     //读取到文件结尾停止
			break
		}
		fmt.Printf("n:%v\n", n)
		fmt.Printf("read:%v\n", string(buf))
	}
	f3.Close()
	f4, _ := os.Open("a.txt")
	//从某个偏移量开始读取，读取内容多少由缓冲区大小决定
	buf := make([]byte, 4)
	at, _ := f4.ReadAt(buf, 3) //从第三个字节开始读取,读4个字节
	fmt.Printf("at:%v\n", at)
	fmt.Printf("f4:%v\n", string(buf))
	f4.Close()
}

//读取目录
func readDir() {
	dir, _ := os.ReadDir("main/a") //读取目录及目录下的文件
	for _, d := range dir {
		fmt.Printf("isDir:%v\n", d.IsDir())
		fmt.Printf("dname:%v\n", d.Name())
	}
}

//定位
func seek() {
	fi, _ := os.Open("a.txt")
	fi.Seek(3, 0) //seek方法，定位从索引为3的位置读取
	buf := make([]byte, 5)
	n, _ := fi.Read(buf)
	fmt.Printf("n:%v\n", n)
	fmt.Printf("buf:%v\n", string(buf))
}
func main() {
	//openclose()
	//createFileTmp()
	//readOps()
	//readDir()
	seek()
}
