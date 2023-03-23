package main

import (
	"fmt"
	"sync"
)

//互斥锁

var x int
var wr sync.WaitGroup

// 声明锁
var mux sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		mux.Lock()
		x += 1
		mux.Unlock()
		//在操作x变量时，先锁住x让x别乱动。操作完再放开x锁。注意这里由于锁损耗性能，所以尽量按最小粒度去加锁
	}
	wr.Done()
}

func main() {
	wr.Add(2)
	go add()
	go add()
	wr.Wait()
	fmt.Println(x)

}

/*
加锁前：
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
6985
felix@MacBook-Pro 0319 % go run main.go
7531

加锁后:
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000
felix@MacBook-Pro 0319 % go run main.go
10000


*/
