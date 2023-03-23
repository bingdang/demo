package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁。写锁类似与互斥锁，读锁在读数据时不锁定并发
var x int
var wr sync.WaitGroup
var rw sync.RWMutex

func add() {
	rw.Lock()
	fmt.Println("wlocak")
	x += 1
	<-time.NewTimer(time.Second * 2).C
	rw.Unlock()
	fmt.Println("wunlocak")
	wr.Done()
}

func read() {
	rw.RLock()
	fmt.Printf("rlocak:%d\n", x)
	<-time.NewTimer(time.Second * 2).C
	rw.RUnlock()
	fmt.Println("readunlocak")
	wr.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wr.Add(1)
		go add()
	}

	for i := 0; i < 10; i++ {
		wr.Add(1)
		go read()
	}

	wr.Wait()

}

/*
felix@MacBook-Pro 0319 % go run main.go
wlocak
wunlocak
rlocak:1
rlocak:1
rlocak:1
rlocak:1
rlocak:1
rlocak:1
rlocak:1
rlocak:1
rlocak:1
rlocak:1
readunlocak
readunlocak
readunlocak
readunlocak
readunlocak
readunlocak
readunlocak
readunlocak
readunlocak
readunlocak
wlocak
wunlocak
wlocak
wunlocak
wlocak
wunlocak
wlocak
wunlocak
wlocak
^Csignal: interrupt
*/
