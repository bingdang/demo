package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
bufio包实现了带缓冲区的读写，是对文件读写的封装
bufio缓冲写数据
	os.O_WRONLY：只写
	os.O_CREATE：创建文件
	os.O_RDONLY：只读
	os.O_RDWR：读写
	os.O_TRUNC：清空
	os.O_APPEND：追加
bufio读数据

*/

func bufiow() {
	// 参数2：打开模式，所有模式都在文档
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	feichifile, err := os.OpenFile("./feichi.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer feichifile.Close()
	// 获取writer对象
	writer := bufio.NewWriter(feichifile)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString("你好\n")
	}
	//刷新缓冲区，强制写出
	writer.Flush()
}

func bufior() {
	feichifile, err := os.OpenFile("./feichi.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer feichifile.Close()
	reader := bufio.NewReader(feichifile)
	for {
		Line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(Line))
	}
}

func ioutilw() {
	err := ioutil.WriteFile("./feichi.txt", []byte("cakepanit.com"), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ioutilr() {
	filecontent, err := ioutil.ReadFile("./feichi.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(filecontent))
}

func main() {
	//bufio实现写
	bufiow()
	//bufio实现读
	bufior()

	//ioutil实现写
	ioutilw()
	//ioutil实现读
	ioutilr()
}

/*
bufio实现:
felix@MacBook-Pro project02 % go run main.go
你好
你好
你好
你好
你好

ioutil实现:
cakepanit.com
*/
