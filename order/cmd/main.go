package main

import (
	"github.com/nmdra/Go-gRPC-Sample-Project/order/config"
	"github.com/nmdra/Go-gRPC-Sample-Project/order/internal/adapters/db"
	"github.com/nmdra/Go-gRPC-Sample-Project/order/internal/adapters/grpc"
	"github.com/nmdra/Go-gRPC-Sample-Project/order/internal/application/core/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
