package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup等待组

func main() {
	/*
		ch1 := make(chan string, 2)
		go func(ch1 chan string) {
			ch1 <- "123"
		}(ch1)
		go func(ch2 chan string) {
			ch2 <- "456"
		}(ch1)

		count := 2
		for v := range ch1 {
			fmt.Println("R", v)
			count--
			if count == 0 {
				close(ch1)
			}
		}
	*/
	//利用等待组实现上面的效果
	ch2 := make(chan string, 2)

	//声明等待组
	var wr sync.WaitGroup
	wr.Add(2) //添加计数,有几组协程添加几个
	go func(ch2 chan string) {
		fmt.Println("协程1")
		ch2 <- "协程1"
		wr.Done() //操作结束上面的计数减1
	}(ch2)
	go func(ch2 chan string) {
		fmt.Println("协程2")
		ch2 <- "协程2"
		wr.Done()
	}(ch2)

	wr.Wait() //等待所有操作结束
}
