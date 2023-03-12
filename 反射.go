package main

import (
	"fmt"
	"reflect"
)

/*
在编程中，反射指的是在运行的过程中看到自己。
- 在实际的编程过程中我们知道，创建的这个变量或者对象是什么类型或者是什么样子的，同时很容易能对它进行操作。
- 在运行过程中，程序没有我们的眼睛，它并不知道这个东西是怎么样的，这个时候就需要运用到反射。通过反射我可以知道自己长什么样子。
*/

//获取interface类型信息
func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	k := t.Kind()
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64") //a is float64
	case reflect.String:
		fmt.Println("a is string")
	}
}

//获取interface值信息
func reflect_value(a interface{}) {
	vl := reflect.ValueOf(a)
	vt := vl.Kind() //返回种类
	switch vt {
	case reflect.Float64:
		fmt.Println("a is ", vl.Float()) //a is  3.14
	case reflect.String:
		fmt.Println("a is ", vl.String())
	}
}

//反射方式修改值
func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	kind := v.Kind()
	fmt.Printf("%v\n", kind) //ptr
	switch kind {
	case reflect.String:
		fmt.Println("a is String")
	case reflect.Ptr:
		//Elme()获取指针地址指向的值
		v.Elem().SetFloat(3.1415)
		fmt.Println(v.Elem().Float(), v.Elem().Kind()) //3.1415 float64

	}
}

func main() {
	var a float64 = 3.14
	reflect_type(a)
	reflect_value(a)
	reflect_set_value(&a) //通过反射方式修改值，必须传递指针
	fmt.Println(a)        //3.1415
}
