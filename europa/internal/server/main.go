package server

import (
	"net"
	"os"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"github.com/hashicorp/go-hclog"
	"github.com/nessus/europa/internal/server/protos"
)

// StartgRPCServer is
func StartgRPCServer(port string, db *gorm.DB, logger hclog.Logger) {
	server := grpc.NewServer()
	reflection.Register(server)

	userService := NewUserService(db, logger)
	protos.RegisterUserServiceServer(server, userService)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("unable to start gRPC server", "port", port, "error", err)
		os.Exit(1)
	}

	server.Serve(lis)
}
