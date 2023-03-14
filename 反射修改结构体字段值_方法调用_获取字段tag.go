package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string `json:"name1" db:"name2"`
	Age  int
}

//1、修改结构体内单个字段
func modifyName(o interface{}) {
	//获取传入对象的指针所指向的值
	v := reflect.ValueOf(o).Elem()

	//获取单个字段
	fieldbyname := v.FieldByName("Name")
	if fieldbyname.Kind() == reflect.String {
		fieldbyname.SetString("gebideng")
	}
}

//2、反射的方法调用
func (u User) Hello(name string) {
	fmt.Println("Hello:", name)
}

func main() {
	user1 := User{1, "feichi", 18}
	//修改结构体内的字段值
	modifyName(&user1)
	fmt.Println(user1)
	//{1 gebideng 18}

	//方法调用
	//1.获取方法
	v := reflect.ValueOf(user1)
	m1 := v.MethodByName("Hello")
	//2.调用方法
	//构造一些 方法的参数，并传入
	args := []reflect.Value{reflect.ValueOf("feichi")}
	//无参数方法 var args2 = []reflect.Value
	m1.Call(args)
	//Hello: feichi

	//获取结构体tag
	k := reflect.TypeOf(&user1).Elem()
	kt1 := k.Field(1).Tag.Get("json")
	kt2 := k.Field(1).Tag.Get("db")
	fmt.Println(kt1, kt2)
	//name1 name2
}
