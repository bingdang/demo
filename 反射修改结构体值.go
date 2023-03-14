package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//修改结构体内单个字段
func modifyName(o interface{}) {
	//获取传入对象的指针所指向的值
	v := reflect.ValueOf(o).Elem()

	//获取单个字段
	fieldbyname := v.FieldByName("Name")
	fieldbyname.SetString("gebideng")
}

func main() {
	user1 := User{1, "feichi", 18}
	modifyName(&user1)
	fmt.Println(user1)
	//{1 gebideng 18}
}
