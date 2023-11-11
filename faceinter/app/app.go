package app

import (
	"faceinter/data"
	"fmt"
)

type App interface {
	Sum(int, int) int
}

type app struct {
	mydesql data.Sql //依赖interface类型
}

func NewAppIns(mydesql data.Sql) App { //不依赖具体实现将mydesql字段当作依赖传递进来
	return &app{
		mydesql: mydesql,
	}
}

func (app1 *app) Sum(a, b int) int {
	err := app1.mydesql.Open()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	return a + b
}
