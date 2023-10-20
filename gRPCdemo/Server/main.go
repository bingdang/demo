package main

import (
	"context"
	"fmt"
	"gRPCdemo/pb"
	"google.golang.org/grpc"
	"net"
)

/*
1.
go get -u google.golang.org/grpc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.3

~ ❯ protoc --version                                                                                                              17:39:36
libprotoc 3.20.3

~ ❯ protoc                                                                                                                        17:39:40
protoc              protoc-gen-go       protoc-gen-go-grpc

*/

// 定义服务
type server struct {
	pb.UnimplementedGreeterServer
}

// 定义方法
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Answer: "Hello " + in.Name}, nil
}

func main() {
	// 向gRPC注册服务
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	s := grpc.NewServer()                  //创建gRPC服务器
	pb.RegisterGreeterServer(s, &server{}) //在gRPC服务端注册服务

	//启动服务
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

}
