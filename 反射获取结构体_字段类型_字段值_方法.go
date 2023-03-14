package main

import (
	"fmt"
	"reflect"
)

/*
反射相关概念
类型：
Go 程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。
*例如：使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。

种类：
种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：
type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)
*type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

反射操作：
reflect 包提供了 reflect.TypeOf 和 reflect.ValueOf 两个函数来获取任意对象的 Value 和 Type。

reflect.TypeOf(x) //返回表示x这个对象 类型名称的字符串
reflect.TypeOf(a).Kind //返回reflect.Kind 种类的常量，例如需要比较种类时需要.Kind进行获取
reflect.xxxx  //表示具体的值，例如 reflect.String
reflect.ValueOf(x) //返回表示x这个对象 的值
	方法名			说明
	Interface() interface {}	将值以 interface{} 类型返回，可以通过类型断言转换为指定类型
	Int() int64	将值以 int 类型返回，所有有符号整型均可以此方式返回
	Uint() uint64	将值以 uint 类型返回，所有无符号整型均可以此方式返回
	Float() float64	将值以双精度（float64）类型返回，所有浮点数（float32、float64）均可以此方式返回
	Bool() bool	将值以 bool 类型返回
	Bytes() []bytes	将值以字节数组 []bytes 类型返回
	String() string	将值以字符串类型返回
	例如: reflect.ValueOf(x).String() 已字符串形式返回x中的值
reflect.ValueOf(x).Kind //返回表示x这个变量的种类
reflect.ValueOf(x).Elem //获取指针地址指向的值
reflect.ValueOf(x).Elem.SetFloat(3.1415) //修改值为3.1415
reflect.ValueOf(x).Elem.Kind //返回表示x这个指针变量存的数据的种类

结构体相关操作：
reflect.TypeOf(o)
reflect.TypeOf(o).NumField() //获取结构体字段个数
reflect.TypeOf(o).Field(i) //获取结构体每个字段 配合for循环
reflect.ValueOf(o).Field(i).Interface() //获取结构体每个字段值 配合for循环
reflect.ValueOf(o).Field(i).Kind //获取结构体内某个字段的种类

reflect.TypeOf(o).NumMethod() //获取结构体方法个数
reflect.TypeOf(o).Method(i) //获取结构体绑定的每个方法 配合for循环

reflect.TypeOf(p).Field(0).Anonymous //返回bool 判断结构体内字段是否时匿名字段
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
