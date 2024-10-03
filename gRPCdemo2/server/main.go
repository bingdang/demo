package main

import (
	"google.golang.org/grpc"
	"net"
	"test/pb"
	"test/serverImpl"
)

func main() {
	//new一个grpc服务
	gser := grpc.NewServer()
	//注册服务
	pb.RegisterMessageSenderServer(gser, serverImpl.MessageSenderServerImpl{})
	//启动gRPC
	tcplis, err := net.Listen("tcp", ":8001")
	if err != nil {
		panic("tcp启动失败:" + err.Error())
	}
	err = gser.Serve(tcplis)
	if err != nil {
		panic("grpc启动失败:" + err.Error())
	}
}
