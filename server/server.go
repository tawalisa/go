package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grcp/mygrcp"
	"log"
	"net"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (s Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

func (s Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

//func (Server) SayHello(c context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//	log.Printf("Received: %v", in.GetName())
//	return &pb.HelloResponse{Message: "Hello " + in.GetName()}, nil
//}
//func (Server) mustEmbedUnimplementedUserServiceServer() {}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
