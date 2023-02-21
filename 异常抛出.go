package main

import "fmt"

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		// 自己抛
		panic("半径不能为负")
	}
	return 3.14 * radius * radius
}

func test04() {
	// 延时执行匿名函数
	// 延时到何时？（1）程序正常结束   （2）发生异常时
	defer func() {
		// recover() 复活 恢复
		// 会返回程序为什么挂了
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(getCircleArea(-5))
	fmt.Println("这里有没有执行")
}

func main() {
	test04()
}

/*
felix@MacBook-Pro project02 % go run main.go
半径不能为负
*/
