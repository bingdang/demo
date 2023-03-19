package main

import (
	"fmt"
	"time"
)

func noBufferChannel() {
	c := make(chan int)
	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i <= 3; i++ {
			c <- i
			fmt.Printf("子协程运行中[%d],len:%d,cap:%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Printf("协程运行中%d\n", num)
	}

	fmt.Println("noBufferChannel函数结束")
}

/*
noBufferChannel()执行
协程运行中0
子协程运行中[0],len:0,cap:0
子协程运行中[1],len:0,cap:0
协程运行中1
协程运行中2
子协程运行中[2],len:0,cap:0
子协程运行中[3],len:0,cap:0
子协程结束
协程运行中3
noBufferChannel函数结束
*/

func haveBufferChannel() {
	c := make(chan int, 4)
	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i <= 3; i++ {
			c <- i
			fmt.Printf("子协程运行中[%d],len:%d,cap:%d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Printf("协程运行中%d\n", num)
	}

	fmt.Println("haveBufferChannel函数结束")
}

/*
haveBufferChannel()执行
子协程运行中[0],len:1,cap:4
子协程运行中[1],len:2,cap:4
子协程运行中[2],len:3,cap:4
子协程运行中[3],len:4,cap:4
子协程结束
协程运行中0
协程运行中1
协程运行中2
协程运行中3
haveBufferChannel函数结束
*/

func channelNoClose() {
	c := make(chan int)
	defer fmt.Println("channelNoClose关闭")
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("数据存入", i)
		}
	}()

	for {
		if v, ok := <-c; ok {
			fmt.Println("数据读取", v)
		} else {
			break
		}
	}
}

/*
数据存入 0
数据读取 0
数据读取 1
数据存入 1
数据存入 2
数据读取 2
数据读取 3
数据存入 3
数据存入 4
数据读取 4
fatal error: all goroutines are asleep - deadlock!
*/

func channelClose() {
	c := make(chan int)
	defer fmt.Println("channelClose关闭")
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("数据存入", i)
		}
		close(c)
	}()

	for {
		if v, ok := <-c; ok {
			fmt.Println("数据读取", v)
		} else {
			break
		}
	}
}

/*
数据存入 0
数据读取 0
数据读取 1
数据存入 1
数据存入 2
数据读取 2
数据读取 3
数据存入 3
数据存入 4
数据读取 4
channelClose关闭
*/

func main() {
	noBufferChannel()
	haveBufferChannel()
	channelClose()
	channelNoClose()
}
