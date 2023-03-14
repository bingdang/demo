package main

import (
	"fmt"
	"reflect"
)

/*
反射相关操作：


*/

type P struct {
	Id   int
	Name string
	Age  int
}

type Dev struct {
	P
	Duty string
}

func (p P) Hello() {
	fmt.Println("p method Hello")
}

// 传入interface
func Pony(o interface{}) {
	//获取类型
	pt := reflect.TypeOf(o)
	fmt.Println("pony 类型:", pt, pt.Name())
	/*
	 p1调用
	   pony 类型: main.P P
	 dev调用
	   pony 类型: main.Dev Dev
	*/

	//获取值
	pv := reflect.ValueOf(o)
	fmt.Println("pony 值:", pv)
	/*
	  p1调用
	    pony 值: {1 fei 18}
	  dev调用
	    pony 值: {{2 chi 23} devops}
	*/

	//获取结构体字段个数，pt.NumField()
	for i := 0; i < pt.NumField(); i++ {
		//获取结构体内每个字段
		f := pt.Field(i)

		//获取结构体内字段值信息，interface(): 获取字段对应的值
		fv := pv.Field(i).Interface()

		fmt.Println(f.Name, f.Type, fv)
		/*
			p1调用
				Id int 1
				Name string fei
				Age int 18

			dev调用
			    P main.P {2 chi 23}
				Duty string devops
		*/
	}

	//获取方法
	for i := 0; i < pv.NumMethod(); i++ {
		m := pt.Method(i)
		fmt.Println(m.Name, m.Type)
	}
	/*
		  p1调用
			Hello func(main.p)
		  dev调用
		    Hello func(main.Dev)
	*/

}

// 是否为匿名字段
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
	p1 := P{1, "fei", 18}
	Pony(p1)

	//判断是否为匿名字段
	// 1.假数据
	dev1 := Dev{P{2, "chi", 23}, "devops"}
	Pony2(dev1)
	Pony(dev1)

}
