package main

import (
	"io"
	"log"
	"os"
	"strings"
)

//io操作，其他例子官方说明
func testCopy() {
	reader := strings.NewReader("hello")
	//copy到标准输出控制台，也可以打开文件，copy到文件中
	_, err := io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	testCopy()
}
