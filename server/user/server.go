package user

import (
	"context"
	pb "github.com/garupanojisan/grpc-sample/proto"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func (u *Server) Get(ctx context.Context, req *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	resp := &pb.UserGetResponse{
		User: &pb.User{
			Id:    req.GetId(),
			Name:  "example user",
			Email: "example@example.com",
		},
	}
	return resp, nil
}
