package main

import (
	"fmt"
	"time"
)

func test1() {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	//time.Now().Unix()获取时间戳
	fmt.Printf("timestamp:%v\n", t.Unix())
	timestamp := t.Unix()         //获取时间戳
	t1 := time.Unix(timestamp, 0) //时间戳转化为时间
	fmt.Printf("t1:%v\n", t1)
	//时间计算可以用add,sub,before等函数
}

func testNow() {
	t := time.Now()
	fmt.Printf("t:%v\n", t)
	year := t.Year()
	fmt.Printf("year:%v\n", year)
	//官方给出格式化初始时间为“2006-01-02 15:04:05”
	fmt.Printf(t.Format("2006/01/02 15:04:05") + "\n")
	//字符串时间转化为时间对象
	t, _ = time.ParseInLocation("2006/01/02", t.Format("2006/01/02"), time.Local)
	fmt.Printf("t:%v\n", t)
}
func main() {
	testNow()
	//test1()
}
