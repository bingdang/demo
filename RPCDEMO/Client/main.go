package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//拨号
	dial, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer dial.Close()

	//构造参数
	var (
		args   = [2]int{8, 9}
		result = 0
	)

	//远程调用
	err = dial.Call("RPCDemo.Add", args, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}
