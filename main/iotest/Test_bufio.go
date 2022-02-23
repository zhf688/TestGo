package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//bufio的读操作
func testBufIo() {
	//r := strings.NewReader("===========")//创建新的reader，也可以打开文件
	r, _ := os.Open("a.txt")
	defer r.Close()
	r1 := bufio.NewReader(r)            //创建带缓冲区的reader
	readString, _ := r1.ReadString('|') //reader调用readstring方法
	fmt.Printf("r1:%v\n", readString)
}

//bufio的写操作
func testBufWrite() {
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0777)
	defer f.Close() //关闭文件
	w := bufio.NewWriter(f)
	w.WriteString("|bufwriter")
	w.Flush() //缓冲区写入后需要重新刷新
}

//Scanner扫描操作
func testScan() {
	r := strings.NewReader("hello 世界！")
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanRunes) //将字符分割
	for s.Scan() {
		fmt.Printf("scanner:%v\n", s.Text())
	}
}
func main() {
	//testBufIo()
	//testBufWrite()
	testScan()
}
