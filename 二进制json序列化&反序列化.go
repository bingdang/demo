package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"math/rand"
)

/*
二进制json
	go get -u -v github.com/vmihailenco/msgpack 下载第三方包

作用，二进制json对于计算机处理时 效率比传统josn效率高，但是人类不可读
*/

type Human struct {
	Name string
	Age  int
	Sex  string
}

//二进制json序列化
func Wfile(filename string) (err error) {
	var Hlist []*Human
	//生成假数据
	for i := 0; i <= 5; i++ {
		h := &Human{
			Name: fmt.Sprintf("name%d", i),
			Age:  rand.Intn(100),
			Sex:  "男",
		}
		Hlist = append(Hlist, h)
	}

	marshal, err := msgpack.Marshal(Hlist)
	if err != nil {
		fmt.Println(err)
		return
	}

	//二进制写出
	err = ioutil.WriteFile(filename, marshal, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//二进制json反序列化
func Rfile(filename string) (err error) {
	var Hlist []*Human
	//读取文件拿到二进制json
	marshal, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	//将二进制json写入Human类型的切片
	err = msgpack.Unmarshal(marshal, &Hlist)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(marshal))

	//遍历切片
	for _, v := range Hlist {
		fmt.Println(v)
	}
	return
}

func main() {
	err := Wfile("./feichi.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Rfile("./feichi.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
felix@MacBook-Pro 2023Project % go run main.go
���Name�name0�Age�Q�Sex�男��Name�name1�Age�W�Sex�男��Name�name2�Age�/�Sex�男��Name�name3�Age�;�Sex�男��Name�name4�Age�Q�Sex�男��Name�name5�Age��Sex�男
&{name0 81 男}
&{name1 87 男}
&{name2 47 男}
&{name3 59 男}
&{name4 81 男}
&{name5 18 男}
*/
