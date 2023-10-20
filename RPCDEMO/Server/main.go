package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type RpcMe struct {
}

func (RpcMe *RpcMe) Add(args [2]int, reply *int) error {
	*reply = args[0] + args[1]
	fmt.Println("count Add", args[0], args[1])
	return nil
}

func main() {
	//创建服务实例
	rpcMe1 := &RpcMe{}
	//注册rpc服务
	err := rpc.Register(rpcMe1)
	if err != nil {
		fmt.Println(err)
		return
	}

	//重命名rpc服务
	err = rpc.RegisterName("RPCDemo", rpcMe1)
	if err != nil {
		fmt.Println(err)
		return
	}

	//监听
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Service Start", os.Getpid())

	for {
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(accept)

	}
}
