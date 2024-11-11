package main

import (
	"log"
	"net"

	// "github.com/liju-github/EcommerceAdminService/proto/user"
	// "github.com/liju-github/EcommerceAdminService/services"
	// util "github.com/liju-github/EcommerceAdminService/utils"
	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	util.SetJWTSecretKey(cfg.JWTSecretKey)

	// Initialize database connection
	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close(dbConn)

	AdminService := service.NewAdminService(AdminRepo)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	Admin.RegisterAdminServiceServer(grpcServer, AdminService)

	log.Println("Admin Service is running on gRPC port: " + cfg.GRPCPort)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server startup failed: %v", err)
	}
}
