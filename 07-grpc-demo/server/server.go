package main

import (
	"context"
	"fmt"
	"grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type appService struct {
	proto.UnimplementedAppServiceServer
}

func (s *appService) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	fmt.Printf("Add Operation: X = %d, Y = %d\n", x, y)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	s := &appService{}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, s)
	grpcServer.Serve(listener)
}
