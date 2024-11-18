package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UserGRPCPort         string
	ContentGRPCPort      string
	AdminGRPCPort        string
	NotificationGRPCPort string
	AdminUsername        string
	AdminPassword        string
	JWTSecretKey         string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return Config{
		AdminGRPCPort:        os.Getenv("ADMIN_GRPC_PORT"),
		UserGRPCPort:         os.Getenv("USER_GRPC_PORT"),
		ContentGRPCPort:      os.Getenv("CONTENT_GRPC_PORT"),
		NotificationGRPCPort: os.Getenv("NOTIFICATION_GRPC_PORT"),
		AdminUsername: os.Getenv("ADMIN_USERNAME"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
		JWTSecretKey:         os.Getenv("JWT_SECRET"),
	}
}
