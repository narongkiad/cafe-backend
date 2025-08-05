package repository

import (
    "context"
    "github.com/narongkiad/cafe-backend/internal/auth/domain"
)

// AuthRepository defines the interface for data access
type AuthRepository interface {
    CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
    GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}