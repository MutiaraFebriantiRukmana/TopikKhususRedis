package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
"redis/domain"
)

// UserRepositoryImpl implements the domain.UserRepository using Redis
type UserRepositoryImpl struct {
	client *redis.Client
}

// NewUserRepository creates a new Redis user repository
func NewUserRepository(client *redis.Client) domain.UserRepository {
	return &UserRepositoryImpl{client: client}
}

// GetByID retrieves a user by ID from Redis
func (r *UserRepositoryImpl) GetByID(ctx context.Context, id string) (*domain.User, error) {
	key := fmt.Sprintf("user:%s", id)
	data, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user from redis: %w", err)
	}

	var user domain.User
	if err := json.Unmarshal(data, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &user, nil
}

// Save persists a user to Redis
func (r *UserRepositoryImpl) Save(ctx context.Context, user *domain.User) error {
	key := fmt.Sprintf("user:%s", user.ID)
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	if err := r.client.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to save user to redis: %w", err)
	}

	return nil
}

// Delete removes a user from Redis
func (r *UserRepositoryImpl) Delete(ctx context.Context, id string) error {
	key := fmt.Sprintf("user:%s", id)
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete user from redis: %w", err)
	}
	return nil
}

