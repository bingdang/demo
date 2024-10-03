package serverImpl

import (
	"context"
	"log"
	"test/pb"
)

// 定义结构体: MessageSenderServerImpl 结构体嵌入了 pb.UnimplementedMessageSenderServer。这是 gRPC 自动生成的类型，用于提供默认的未实现方法，便于后续实现。
type MessageSenderServerImpl struct {
	*pb.UnimplementedMessageSenderServer
}

func (MessageSenderServerImpl) Send(ctx context.Context, req *pb.MessageReq) (*pb.MessageReply, error) {
	log.Printf("Receive Message:%s", req.GetSaySomething())
	return &pb.MessageReply{
		ResponseSomething: req.GetSaySomething() + "是笨蛋",
	}, nil
}
