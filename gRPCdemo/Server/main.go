package main

import (
	"context"
	"errors"
	"fmt"
	"gRPCdemo/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net"
	"time"
)

/*
把大象放进冰箱分几步？

    把冰箱门打开。
    把大象放进去。
    把冰箱门带上。

gRPC开发同样分三步：
    编写.proto文件，生成指定语言源代码
    编写服务端代码
    编写客户端代码

1.
go get -u google.golang.org/grpc
https://github.com/protocolbuffers/protobuf/releases/tag/v3.20.3

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 //安装go语言插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 //安装grpc插件

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

	//业务逻辑
	//time.Sleep(time.Millisecond * 50)
	time.Sleep(time.Second * 3)
	fmt.Println("有人调SayHello啦")
	select {
	case <-ctx.Done():
		return nil, errors.New("time out")
	default:
		return &pb.HelloReply{
			Answer: "Hello " + in.Name,
			Ts:     timestamppb.Now(),
		}, nil
	}
}

func (s *server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	fmt.Println("有人调Add啦")

	select {
	case <-ctx.Done():
		return nil, errors.New("time out")
	default:
		return &pb.AddReply{
			S:  in.GetC1() + in.GetC2(),
			Ts: timestamppb.Now(),
		}, nil
	}
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
	fmt.Println("START")
	err = s.Serve(listen)
	if err != nil {
		fmt.Printf("failed to start: %v", err)
		return
	}

}
