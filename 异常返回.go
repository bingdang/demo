package main

import (
	"errors"
	"fmt"
)

func CircleCalculation(radius float32) (surface float32, err error) {
	if radius <= 0 {
    // 构建个异常对象
		err = errors.New("半径不能为负")
		return
	}
	surface = 3.14 * radius * radius
	return
}

func main() {
	surface, err := CircleCalculation(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(surface)
	}
}

/*
felix@MacBook-Pro project02 % go run main.go
78.5
felix@MacBook-Pro project02 % go run main.go
半径不能为负
felix@MacBook-Pro project02 % echo $?
0
*/
