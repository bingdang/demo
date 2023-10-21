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
	c1   = flag.Int64("c1", 0, "第一个加数")
	c2   = flag.Int64("c2", 0, "第二个加数")
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
		//log.Fatalf("cloud not greet: %v", err)
	}

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	radd, err := client.Add(ctx2, &pb.AddRequest{
		C1: *c1,
		C2: *c2,
	})
	if err != nil {
		log.Fatalf("cloud not add: %v", err)
	}

	log.Printf("Greeting: %s,%v", r.GetAnswer(), r.GetTs().AsTime().Format("2006-01-02 15:04:05"))
	log.Printf("Greeting: %v,%v", radd.GetS(), radd.GetTs().AsTime().Format("2006-01-02 15:04:05"))
}
