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

type dev struct {
	p
	duty string
}

func (p p) Hello() {
	fmt.Println("Hello")
}

//传入interface
func Pony(o interface{}) {
	fmt.Println(reflect.ValueOf(o).Field(1).Kind())
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
		fmt.Println(f)
		if reflect.ValueOf(o).Field(i).Kind() != reflect.Struct { //如果结构体内存在匿名字段 另一个结构体则不执行
			//获取结构体内字段值信息，interface(): 获取字段对应的值
			fv := pv.Field(i).Interface()

			fmt.Println(f.Name, f.Type, fv)
			/*
				Id int 1
				Name string fei
				Age int 18
			*/
		}
	}

	//获取方法
	for i := 0; i < pv.NumMethod(); i++ {
		m := pt.Method(i)
		fmt.Println(m.Name, m.Type)
	}
	/*
		Hello func(main.p)
	*/

	//是否为匿名字段

}

func Pony2(p interface{}) {
	pt := reflect.TypeOf(p)
	//Anonymous 匿名
	fmt.Printf("Pony2 %#v\n", pt.Field(0))
	//Pony2 reflect.StructField{Name:"p", PkgPath:"main", Type:(*reflect.rtype)(0x10e71a0), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	fmt.Printf("Pony2 %#v\n", pt.Field(1))
	//Pony2 reflect.StructField{Name:"duty", PkgPath:"main", Type:(*reflect.rtype)(0x10d68a0), Tag:"", Offset:0x20, Index:[]int{1}, Anonymous:false}
	fmt.Println("Pony2", pt.Field(0).Anonymous) //Pony2 true

	fmt.Println("Pony2", reflect.ValueOf(p).Field(0))
	//Pony2 {2 chi 23}

	fmt.Println("Pony2", reflect.ValueOf(p).Field(0).Kind())
	//Pony2 struct

	fmt.Println("Pony2", reflect.ValueOf(p).Field(1).Kind())
	//Pony2 string

}

func main() {
	p1 := p{1, "fei", 18}
	Pony(p1)

	//判断是否为匿名字段
	// 1.假数据
	dev1 := dev{p{2, "chi", 23}, "devops"}
	Pony2(dev1)
	Pony(dev1)

}
