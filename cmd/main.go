package main

import (
	"log"
	"net"

	"github.com/liju-github/EcommerceAdminService/config"
	"github.com/liju-github/EcommerceAdminService/proto/admin"
	"github.com/liju-github/EcommerceAdminService/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to the User gRPC server
	connUser, err := grpc.NewClient("localhost:"+cfg.UserGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to User gRPC server: %v", err)
	}
	defer connUser.Close()

	// Connect to the Content gRPC server
	connContent, err := grpc.NewClient("localhost:"+cfg.ContentGRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to Content gRPC server: %v", err)
	}
	defer connContent.Close()

	// Initialize the AdminService with the connections
	adminService := services.NewAdminService(connUser, connContent)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":"+cfg.AdminGRPCPort)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	admin.RegisterAdminServiceServer(grpcServer, adminService)

	log.Println("Admin Service is running on gRPC port: " + cfg.AdminGRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server startup failed: %v", err)
	}
}
