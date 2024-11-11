package key_manager

import (
	"context"

	"github.com/NesterovYehor/TextNest/services/key_generation_service/internal/keymanager"
	"github.com/redis/go-redis/v9"
)

// KeyManagerServer now includes a redis.Client
type KeyManagerServer struct {
	UnimplementedKeyManagerServiceServer
	RedisClient *redis.Client
}

// NewKeyManagerServer creates a new KeyManagerServer with a Redis client
func NewKeyManagerServer(redisClient *redis.Client) *KeyManagerServer {
	return &KeyManagerServer{
		RedisClient: redisClient,
	}
}

// GetKey now uses the Redis client passed in the server struct
func (s *KeyManagerServer) GetKey(ctx context.Context, req *GetKeyRequest) (*GetKeyResponse, error) {
	key, err := keymanager.GetKey(s.RedisClient) // Pass Redis client to GetKey
	if err != nil {
		return &GetKeyResponse{Error: err.Error()}, nil
	}
	return &GetKeyResponse{Key: key}, nil
}

// ReallocateKey now uses the Redis client passed in the server struct
func (s *KeyManagerServer) ReallocateKey(ctx context.Context, req *ReallocateKeyRequest) (*ReallocateKeyResponse, error) {
	err := keymanager.ReallocateKey(req.GetKey(), s.RedisClient) // Pass Redis client to ReallocateKey
	if err != nil {
		return &ReallocateKeyResponse{Error: err.Error()}, nil
	}
	return &ReallocateKeyResponse{Message: "Key reallocated successfully"}, nil
}