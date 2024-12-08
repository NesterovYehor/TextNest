package config

import (
	"context"
	"fmt"
	"os"

	jsonlog "github.com/NesterovYehor/TextNest/pkg/logger"
	"gopkg.in/yaml.v3"
)

const DefaultConfigFile = "config.yaml"

type Config struct {
    Grpc struct {
        Port string `yaml:"port"`
        Host string `yaml:"host"`
    } `yaml:"grpc"`
    DBURL      string `yaml:"db"`
    BucketName string `yaml:"s3.bucket_name"`
    S3Region   string `yaml:"region"`
}

// LoadConfig loads the configuration from a YAML file.
func LoadConfig(log *jsonlog.Logger, ctx context.Context) (*Config, error) {
	file := os.Getenv("CONFIG_FILE")
	if file == "" {
		file = DefaultConfigFile
	}

	data, err := os.Open(file)
	if err != nil {
		log.PrintFatal(ctx, fmt.Errorf("failed to read configuration file: %w", err), nil)
		return nil, err
	}
	defer data.Close()

	var cfg Config
	decoder := yaml.NewDecoder(data)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.PrintFatal(ctx, fmt.Errorf("failed to parse configuration file: %w", err), nil)
		return nil, err
	}

	// Validate required fields
	if cfg.Grpc == nil || cfg.Grpc.Port == "" {
		log.PrintFatal(ctx, fmt.Errorf("gRPC configuration is incomplete"), nil)
	}
	if cfg.DBURL == "" {
		log.PrintFatal(ctx, fmt.Errorf("database URL is not set"), nil)
	}
	if cfg.BucketName == "" || cfg.S3Region == "" {
		log.PrintError(ctx, fmt.Errorf("S3 configuration is incomplete, some features may be unavailable"), nil)
	}

	return &cfg, nil
}
