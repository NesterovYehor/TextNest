package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/IBM/sarama"
	"github.com/NesterovYehor/TextNest/pkg/kafka"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/config"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/handlers"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/repository"
	"github.com/NesterovYehor/TextNest/services/cleanup_service/internal/services"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Initialize configuration
	var cfg config.Config
	cfg.Init()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	db, err := openDB("")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize and start the expiration service
	expirationService := services.NewExpirationService(db)
	expirationService.Start(&cfg, ctx)
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

func runKafkaConsumer(cfg *config.Config, ctx context.Context, repo *repository.PasteRepository) error {
	handlers := map[string]kafka.MessageHandler{
		"delete-expired-paste-topic": func(msg *sarama.ConsumerMessage) error {
			return handlers.HandleDeleteExpiredPaste(msg, repo)
		},
	}

	consumer, err := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.GroupID, cfg.Kafka.Topics, handlers, ctx)
	if err != nil {
		log.Println("Failed to create a new Kafka consumer:", err)
		return err
	}

	if err := consumer.Start(); err != nil {
		log.Println("Kafka consumer stopped with error:", err)
		consumer.Close()
		return err
	}

	return nil
}
