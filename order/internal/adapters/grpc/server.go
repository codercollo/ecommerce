package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/codercollo/ecommerce/golang/order"
	"github.com/codercollo/ecommerce/order/config"
	"github.com/codercollo/ecommerce/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Adapter connects gRPC network transport requests to the internal API port/
type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

// NewAdapter constructs a new gRPC server adapter instance.
func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

// Run boots up the TCP listener and starts serving gRPC requests.
func (a Adapter) Run() {
	var err error

	// Initialize the TCP listener on configured application port.
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	// Register order service implementation with gRPC server.
	//
	// Enable gRPC Server reflection in development for schema inspection
	order.RegisterOrderServer(grpcServer, a)
	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	// Start accepting incoming gRPC client connections
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port")
	}
}
