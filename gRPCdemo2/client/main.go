package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"test/pb"
)

func main() {
	//此行尝试连接运行在 localhost 上的 gRPC 服务器，端口为 8001，并使用不安全凭据。
	dial, err := grpc.Dial("localhost:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("gRPC连接失败" + err.Error())
	}
	log.Printf("gRPC建联成功")
	defer dial.Close()
	//New一个gRPC Client对象
	msc := pb.NewMessageSenderClient(dial)
	send, err := msc.Send(context.Background(), &pb.MessageReq{
		SaySomething: "虞锋",
	})
	if err != nil {
		panic("gRPC请求失败" + err.Error())
	}
	log.Printf("gRPC reply:%s", send.ResponseSomething)
}
