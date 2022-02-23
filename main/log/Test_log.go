package main

import (
	"fmt"
	"log"
	"os"
)

//print打印及printf格式化打印
func test1() {
	log.Print("logprint")
	log.Printf("printf:%d\n", 100)
	name := "tom"
	age := 20
	log.Println(name, " ", age)
}

//
func test2() {
	log.Print("begin")
	defer log.Print("panic end...")
	log.Panic("panic") //panic会打印日志并抛出异常，结束程序执行,但defer会执行
	log.Print("end...")
}

//fatal
func test3() {
	log.Print("begin")
	defer log.Print("fatal end...")
	log.Fatal("fatal") //fatal打印日志并退出程序，defer及后边的程序均不执行
	log.Print("end...")
}
func init() {
	//初始化设置log日志输出格式，Lshortfile表示文件名，Llongfile表示文件名和文件路径
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//设置日志输出前缀
	log.SetPrefix("Testlog: ")
	//设置日志输出到文件
	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("日志文件输出错误")
		log.Fatal("")
	}
	log.SetOutput(f)
}
func main() {
	i := log.Flags()
	fmt.Printf("flags:%v\n", i)
	log.Print("my log")
}
