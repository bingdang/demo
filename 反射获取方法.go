package main

import (
	"fmt"
	"reflect"
)

type p struct {
	Id   int
	Name string
	Age  int
}

func (p p) Hello() {
	fmt.Println("Hello")
}

//传入interface
func Pony(o interface{}) {
	//获取类型
	pt := reflect.TypeOf(o)
	fmt.Println("类型:", pt, pt.Name()) //类型: main.p p

	//获取值
	pv := reflect.ValueOf(o)
	fmt.Println("值:", pv) //值: {1 fei 18}

	//获取结构体字段个数，pt.NumField()
	for i := 0; i < pt.NumField(); i++ {
		//获取结构体内每个字段
		f := pt.Field(i)

		//获取结构体内字段值信息，interface(): 获取字段对应的值
		fv := pv.Field(i).Interface()

		fmt.Println(f.Name, f.Type, fv)
		/*
			Id int 1
			Name string fei
			Age int 18
		*/
	}

	//获取方法
	for i := 0; i < pv.NumMethod(); i++ {
		m := pt.Method(i)
		fmt.Println(m.Name, m.Type)
	}
	/*
		Hello func(main.p)
	*/

}

func main() {
	p1 := p{1, "fei", 18}
	Pony(p1)

}
