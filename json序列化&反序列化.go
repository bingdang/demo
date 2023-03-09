package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	//`json:""` 为struct tag，序列化时将k首字母转为小心、"-"标识该字段不参与序列化
	Name  string `json:"name"`
	Hobby string `json:"hobby"`
	age   int    `json:"-"`
}

type Mmp struct {
	Age       int    `json:"age"`
	Name      string `json:"name"`
	Hobby     string `json:"hobby"`
	Niubility bool   `json:"niubility"`
}

//json序列化
func jsonEncoded(person interface{}) []byte {
	marshal, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	}
	return marshal
}

//json反序列化
func jsonUnencoded(Mmpjson []byte, mmp interface{}) {
	err := json.Unmarshal(Mmpjson, mmp)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//结构体json序列化
	p1 := Person{
		Name:  "fei",
		Hobby: "女",
		age:   8,
	}
	fmt.Println(string(jsonEncoded(p1)))

	//mapjson序列化
	mmp := make(map[string]interface{})
	mmp["name"] = "Aei"
	mmp["hobby"] = "男"
	mmp["age"] = 18
	mmp["niubility"] = true
	fmt.Println(string(jsonEncoded(mmp)))
	//{"age":18,"hobby":"男","name":"Aei","niubility":true}

	//json反序列化到结构体
	p := []byte(`{"age":18,"hobby":"男","name":"Aei","niubility":true}`)
	var mmp1 Mmp
	jsonUnencoded(p, &mmp1)
	fmt.Println(mmp1)
	//{18 Aei 男 true}

	//json反序列化到interface{}
	var i interface{}
	jsonUnencoded(p, &i)

	//自动转为map
	fmt.Println(i)
	//map[age:18 hobby:男 name:Aei niubility:true]

	//接口的类型断言判断类型 类型断言： value,ok := 元素.(Type)，ok为bool类型
	//map是个无序列表，可以用range遍历
	m := i.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case int:
			fmt.Println(k, "是int类型", vv)
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		case bool:
			fmt.Println(k, "是布尔类型", vv)
		default:
			fmt.Println("其他")
		}
	}
	/*
		age 是float64类型 18
		hobby 是string类型 男
		name 是string类型 Aei
		niubility 是布尔类型 true

	*/
}

/*
输出
 序列化
{"name":"fei","hobby":"女"}
{"age":18,"hobby":"男","name":"Aei","niubility":true}

 反序列化
{18 Aei 男 true}
 类型断言
age 是float64类型 18
hobby 是string类型 男
name 是string类型 Aei
niubility 是布尔类型 true
*/
