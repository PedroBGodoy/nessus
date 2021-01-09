package main

import (
	"context"
	"log"
	"net"

	"github.com/nessus/europa/internal/user"
	"github.com/nessus/europa/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Authenticate(ctx context.Context, request *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	token := request.GetToken()

	u, err := user.Authenticate(token)
	if err != nil {
		res := &pb.AuthenticationResponse{
			User:  nil,
			Error: err.Error(),
		}

		return res, nil
	}

	uPb := &pb.User{
		Id:    u.UserID,
		Name:  u.Name,
		Email: u.Email,
	}

	res := &pb.AuthenticationResponse{
		User:  uPb,
		Error: "",
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to list %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
