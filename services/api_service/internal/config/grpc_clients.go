package config

import (
	"log"

	download_service "github.com/NesterovYehor/TextNest/services/api_service/internal/grpc_client/download_service_client"
	"github.com/NesterovYehor/TextNest/services/api_service/internal/grpc_client/key_manager_client"
	"github.com/NesterovYehor/TextNest/services/api_service/internal/grpc_client/upload_service_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// InitializeKeyManagerClient initializes the Key Manager gRPC client
func InitializeKeyManagerClient(cfg *Config) {
	kgsConn, err := grpc.NewClient(cfg.Grpc.KGSAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // Use Dial instead of NewClient
	if err != nil {
		log.Fatalf("Failed to connect to Key Manager service: %v", err)
	}
	cfg.KeyManager = key_manager.NewKeyManagerServiceClient(kgsConn)

	// No need to close immediately as we need the connection to persist
}

// InitializeUploadClient initializeS the Upload Service gRPC client
func InitializeUploadClient(cfg *Config) {
	uploadConn, err := grpc.NewClient(cfg.Grpc.UploadAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // Use Dial instead of NewClient
	if err != nil {
		log.Fatalf("Failed to connect to Upload service: %v", err)
	}
	cfg.UploadService = upload_service.NewUploadServiceClient(uploadConn)

	// No need to close immediately as we need the connection to persist
}

// InitializeUploadClient initializeS the Upload Service gRPC client
func InitializeDownloadClient(cfg *Config) {
	downloadConn, err := grpc.NewClient(cfg.Grpc.DownloadAddr, grpc.WithTransportCredentials(insecure.NewCredentials())) // Use Dial instead of NewClient
	if err != nil {
		log.Fatalf("Failed to connect to Upload service: %v", err)
	}
	cfg.DownloadService = download_service.NewDownloadServiceClient(downloadConn)

	// No need to close immediately as we need the connection to persist
}
