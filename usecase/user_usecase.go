package usecase

import (
	"context"
	"fmt"

	"redis/domain"
)

// UserUsecase contains business logic for user operations
type UserUsecase struct {
	userRepo domain.UserRepository
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(userRepo domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

// GetUser retrieves a user by ID
func (u *UserUsecase) GetUser(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

// CreateUser creates a new user
func (u *UserUsecase) CreateUser(ctx context.Context, user *domain.User) error {
	if user.ID == "" || user.Name == "" || user.Email == "" {
		return fmt.Errorf("user ID, name and email are required")
	}

	return u.userRepo.Save(ctx, user)
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("user ID is required")
	}

	return u.userRepo.Delete(ctx, id)
}

