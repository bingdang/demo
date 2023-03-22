package main

import (
	"fmt"
	"time"
)

// select监控管道是否打满
func main() {
	ch1 := make(chan string, 3)
	go func(ch1 chan string) {
		for {
			select {
			case ch1 <- "hello":
				fmt.Println("W:Hello")
			default:
				fmt.Println("管道满了") //当管道满时执行这句话
			}
			<-time.NewTimer(time.Millisecond * 500).C
			//写入数据延迟500ms
		}
	}(ch1)

	for v := range ch1 {
		fmt.Println("R", v)
		<-time.NewTimer(time.Second * 1).C
		//读取数据延迟1s
	}

}

/*
felix@MacBook-Pro 0319 % go run main.go
W:Hello
R hello
W:Hello
W:Hello
R hello
W:Hello
R hello
W:Hello
W:Hello
R hello
W:Hello
管道满了
管道满了
R hello
W:Hello
管道满了
R hello
W:Hello
R hello
管道满了
W:Hello
R hello
W:Hello
管道满了
R hello
W:Hello
管道满了
R hello
W:Hello
管道满了
R hello
W:Hello
管道满了
R hello
W:Hello
管道满了
R hello
W:Hello
^Csignal: interrupt
*/
