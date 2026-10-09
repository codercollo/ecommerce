package main

import (
	"log"

	"github.com/codercollo/ecommerce/order/config"
	"github.com/codercollo/ecommerce/order/internal/adapters/db"
	"github.com/codercollo/ecommerce/order/internal/adapters/grpc"
	"github.com/codercollo/ecommerce/order/internal/application/api"
)

// main boots up the Order microservice application and its adapters.
func main() {
	// Initialize outbound database storage adapter.
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)

	// Initialize inbound gRPC transport adapter and start server
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
