package main

import (
	"fmt"
	"time"
)

func c1add(c chan string) {
	<-time.NewTimer(time.Second * 3).C
	c <- "c1add"
}
func c2add(c chan string) {
	<-time.NewTimer(time.Second * 2).C
	c <- "c2add"
}

func main() {
	//创建两个管道
	c1 := make(chan string)
	c2 := make(chan string)

	//管道需要放到协程中，否则报错
	go c1add(c1)
	go c2add(c2)

	select {
	case str := <-c1:
		fmt.Println(str)
	case str2 := <-c2: //监控到管道2首先存入数据
		fmt.Println(str2) //执行这一行
	}
	/*
		c2add
	*/

}
