package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/IBM/sarama"
	"github.com/NesterovYehor/TextNest/pkg/kafka"
	jsonlog "github.com/NesterovYehor/TextNest/pkg/logger"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/config"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/handlers"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/repository"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/services"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Initialize configuration
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	log := jsonlog.New(logFile, slog.LevelInfo)

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.PrintFatal(ctx, err, nil)
	}
	db, err := openDB("")
	if err != nil {
		log.PrintFatal(ctx, err, nil)
	}
	defer db.Close()

	// Initialize and start the expiration service
	expirationService := services.NewExpirationService(db)
	log.PrintInfo(ctx, "Expiration Service Is Started", nil)
	expirationService.Start(cfg, ctx, log)
}

// openDB initializes a new database connection and checks for errors
func openDB(dsn string) (*sql.DB, error) {
	// Open the database connection
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Connected to the database successfully")

	return db, nil
}

func runKafkaConsumer(cfg *config.Config, ctx context.Context, metadataRepo *repository.MetadataRepository, storagerepo *repository.StorageRepository, log *jsonlog.Logger) error {
	srv := services.NewPasteService(*metadataRepo, *storagerepo)
	handlers := map[string]kafka.MessageHandler{
		"delete-expired-paste-topic": func(msg *sarama.ConsumerMessage) error {
			return handlers.HandleDeleteExpiredPaste(ctx, msg, srv, log, cfg.BucketName)
		},
	}

	consumer, err := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.GroupID, cfg.Kafka.Topics, handlers, ctx)
	if err != nil {
		log.PrintError(ctx, fmt.Errorf("Failed to create a new Kafka consumer:", err), nil)
		return err
	}

	if err := consumer.Start(); err != nil {
		log.PrintError(ctx, fmt.Errorf("Kafka consumer stopped with error:", err), nil)
		consumer.Close()
		return err
	}

	return nil
}
