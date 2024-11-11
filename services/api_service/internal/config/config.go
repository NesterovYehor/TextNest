package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/NesterovYehor/TextNest/pkg/http"
	key_manager "github.com/NesterovYehor/TextNest/services/api_service/internal/grpc_client/key_manager_client"
	upload_service "github.com/NesterovYehor/TextNest/services/api_service/internal/grpc_client/upload_service_client"
)

var (
	version   string
	buildTime string
)

type Config struct {
	Env           string
	Http          *httpserver.Config
	KeyManager    key_manager.KeyManagerServiceClient
	UploadService upload_service.UploadServiceClient
	Grpc          struct {
		UploadAddr string
		KGSAddr    string
	}
}

// InitConfig initializes the configuration, including gRPC clients
func InitConfig() *Config {
	cfg := &Config{}
	port := ":8989"
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.Grpc.UploadAddr, "uploadAddr", "localhost:3489", "Upload service address")
	flag.StringVar(&cfg.Grpc.KGSAddr, "kgsAddr", "localhost:5555", "Key Manager service address")

	cfg.Http = httpserver.NewConfig(port)

	displayVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()

	if *displayVersion {
		fmt.Printf("Service Version:\t%s\n", version)
		fmt.Printf("Build Time:\t%s\n", buildTime)
		os.Exit(0)
	}

	// Initialize gRPC clients
	InitializeKeyManagerClient(cfg)
	InitializeUploadClient(cfg)

	return cfg
}

