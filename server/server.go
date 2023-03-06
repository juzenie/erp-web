package main

import (
	"context"
	"erp-web/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) GetOrderList(ctx context.Context, request *proto.OrderRequest) (*proto.OrderListResponse,
	error) {
	fmt.Println("v记录未找到")
	return nil, status.Errorf(codes.NotFound, "记录未找到")
	//return &proto.HelloReply{
	//	Message: "hello "+request.Name,
	//}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterOrderServer(g, &Server{})
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
