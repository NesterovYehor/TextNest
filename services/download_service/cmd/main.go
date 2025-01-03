package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/NesterovYehor/TextNest/pkg/grpc"
	jsonlog "github.com/NesterovYehor/TextNest/pkg/logger"
	pb "github.com/NesterovYehor/TextNest/services/download_service/api"
	"github.com/NesterovYehor/TextNest/services/download_service/internal/config"
	"github.com/NesterovYehor/TextNest/services/download_service/internal/coordinators"
)

func main() {
	log, err := setupLogger("app.log")
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		return
	}

	// Setup graceful shutdown on SIGINT or SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Initialize configuration
	cfg, err := config.LoadConfig(log, ctx)
	if err != nil {
		log.PrintError(ctx, err, nil)
	}

	// Initialize gRPC server
	grpcSrv := grpc.NewGrpcServer(cfg.Grpc)

	coord, err := coordinators.NewDownloadCoordinator(ctx, cfg, log)
	if err != nil {
		log.PrintError(ctx, err, nil)
	}
	// Register the UploadService with the gRPC server
	pb.RegisterPasteDownloadServer(grpcSrv.Grpc, coord)

	// Run the gRPC server
	if err := grpcSrv.RunGrpcServer(ctx); err != nil {
		log.PrintFatal(ctx, err, nil)
		return
	}
}

// setupLogger initializes the application logger
func setupLogger(logFilePath string) (*jsonlog.Logger, error) {
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		logFile.Close()
		return nil, err
	}
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	return jsonlog.New(multiWriter, slog.LevelInfo), nil
}
