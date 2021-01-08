package main

import (
	"context"
	"log"
	"net"

	"github.com/nessus/europa/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Authenticate(ctx context.Context, request *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	u := &pb.User{
		Id:       "1",
		Email:    "teste@teste.com",
		Name:     "teste",
		Password: "",
	}

	res := &pb.AuthenticationResponse{
		User:  u,
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
