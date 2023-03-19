package main

import "fmt"

// 单向channel
// 生产者
func producer(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i
	}
}

// 消费者
func consumer(in <-chan int) {
	for data := range in {
		fmt.Println(data)
	}
}

func main() {
	var c1 = make(chan int)
	go producer(c1)
	consumer(c1)
}

/*
0
1
2
3
4
*/
