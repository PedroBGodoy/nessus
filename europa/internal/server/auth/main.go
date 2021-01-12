package auth

import (
	"context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewServer is
func NewServer(ctx context.Context, port string) (*UserService, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("unable to listen on port %v", err)
	}
	server := grpc.NewServer()
	reflection.Register(server)

	userService := &UserService{
		ctx: ctx,
	}
	RegisterUserServiceServer(server, userService)

	if err = server.Serve(lis); err != nil {
		log.Fatalf("unable to listen on port %v", err)
	}

	return userService, nil
}
