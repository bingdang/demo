package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	m1 := make(map[int]int, 10)
	var wl sync.Mutex

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			wl.Lock()
			m1[x] = x
			wl.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(m1)

}

/*
不加锁
fatal error: concurrent map writes

加锁
felix@MacBook-Pro 0319 % go run main.go
map[0:0 1:1 2:2 3:3 4:4 5:5]
*/
