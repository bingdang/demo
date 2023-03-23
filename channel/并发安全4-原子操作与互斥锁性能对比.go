package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作与互斥锁操作性能对比
var x int64
var wr sync.WaitGroup

// 定义互斥锁
var rw sync.Mutex

func addMutex() {
	defer wr.Done()
	for i := 0; i < 10000000; i++ {
		rw.Lock()
		x += 1
		rw.Unlock()
	}
}

func addAtomic() {
	defer wr.Done()
	for i := 0; i < 10000000; i++ {
		atomic.AddInt64(&x, 1)
	}
}

// 互斥锁时间		Time56141658000
// 原子操作时间	Time19930707000
func main() {
	startTime := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		wr.Add(1)
		//go addMutex()
		go addAtomic()
	}
	wr.Wait()
	endTime := time.Now().UnixNano()
	fmt.Printf("Time%d\n", endTime-startTime)
}
