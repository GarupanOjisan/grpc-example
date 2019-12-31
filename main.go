package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/garupanojisan/grpc-sample/protobuf"
	"github.com/garupanojisan/grpc-sample/server/user"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(s, &user.Server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
