package main

import (
	"github.com/hashicorp/go-hclog"

	"github.com/nessus/europa/infra/database"
	"github.com/nessus/europa/internal/server"
)

const (
	port string = ":50051"
)

func main() {
	logger := hclog.Default()

	db := database.InitSQLite(logger)

	server.StartgRPCServer(port, db, logger)
}
