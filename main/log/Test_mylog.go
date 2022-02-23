package main

import (
	"log"
	"os"
)

//自定义日志输出
var logger *log.Logger

func init() {
	f, err := os.OpenFile("mylog.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		logger.Printf("日志文件输出错误")
	}
	//自定义输出的文件，前缀，格式等
	logger = log.New(f, "Testlog: ", log.Ldate|log.Ltime|log.Lshortfile)

}
func main() {
	logger.Printf("预下单输出:%v\n", 100)
}
