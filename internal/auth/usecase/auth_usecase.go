package usecase

import (
    "context"
    "fmt"
    "time"

    "golang.org/x/crypto/bcrypt"
    "your-module-name/internal/auth/domain"
    "your-module-name/internal/auth/repository"
    "your-module-name/pkg/jwt"
)

// AuthUsecase defines the interface for auth usecases
type AuthUsecase interface {
    Register(ctx context.Context, email, password string) (*domain.User, error)
    Login(ctx context.Context, email, password string) (string, error)
}

type authUsecase struct {
    authRepo       repository.AuthRepository
    contextTimeout time.Duration
}

// NewAuthUsecase creates a new instance of AuthUsecase
func NewAuthUsecase(repo repository.AuthRepository, timeout time.Duration) AuthUsecase {
    return &authUsecase{
        authRepo:       repo,
        contextTimeout: timeout,
    }
}

func (au *authUsecase) Register(ctx context.Context, email, password string) (*domain.User, error) {
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, fmt.Errorf("failed to hash password: %w", err)
    }

    user := &domain.User{
        Email:    email,
        Password: string(hashedPassword),
    }

    // Save user to the database
    createdUser, err := au.authRepo.CreateUser(ctx, user)
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    return createdUser, nil
}

func (au *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
    // Find the user by email
    user, err := au.authRepo.GetUserByEmail(ctx, email)
    if err != nil {
        return "", fmt.Errorf("invalid credentials: %w", err)
    }

    // Compare the provided password with the stored hashed password
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return "", fmt.Errorf("invalid credentials: %w", err)
    }

    // Generate JWT token
    token, err := jwt.GenerateToken(user.ID)
    if err != nil {
        return "", fmt.Errorf("failed to generate token: %w", err)
    }

    return token, nil
}