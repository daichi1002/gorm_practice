package server

import (
	pb "test/post"
	"test/server/service"

	"google.golang.org/grpc"
)

func NewServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterPostServiceServer(server, &service.PostsService{})

	return server
}