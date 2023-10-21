package main

import (
	"context"
	"flag"
	"gRPCdemo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "feichi"
)

var (
	addr = flag.String("addr", "127.0.0.1:8080", "the address to connect to")
	name = flag.String("name", defaultName, "Nameg to greet")
)

func main() {
	flag.Parse()
	//链接到server端
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//创建grpc客户端
	client := pb.NewGreeterClient(conn)

	//上下文超时，1秒;	Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//调用SayHello方法
	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("cloud not greet: %v", err)
	}

	log.Printf("Greeting: %s,%v", r.GetAnswer(), r.GetTs().AsTime().Format("2006-01-02 15:04:05"))
}
