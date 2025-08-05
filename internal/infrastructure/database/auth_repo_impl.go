package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/narongkiad/cafe-backend/internal/auth/domain"

	// Replace with your actual database driver
	_ "github.com/lib/pq"
)

// authRepositoryImpl is the concrete implementation of AuthRepository
type authRepositoryImpl struct {
    db *sql.DB
}

// NewAuthRepository creates a new instance of AuthRepository
func NewAuthRepository(db *sql.DB) *authRepositoryImpl {
    return &authRepositoryImpl{db: db}
}

func (r *authRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
    // You would typically use a library like UUID to generate a unique ID
    query := `
        INSERT INTO users (id, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id, email, created_at, updated_at
    `
    // Generate a UUID for the user
    user.ID = "some-generated-uuid" // replace with real UUID generation

    err := r.db.QueryRowContext(ctx, query, user.ID, user.Email, user.Password).
        Scan(&user.ID, &user.Email, &user.CreatedAt, &user.UpdatedAt)
    
    if err != nil {
        // Handle unique constraint errors (e.g., email already exists)
        return nil, fmt.Errorf("failed to create user: %w", err)
    }

    return user, nil
}

func (r *authRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
    query := `
        SELECT id, email, password, created_at, updated_at
        FROM users
        WHERE email = $1
    `
    user := &domain.User{}
    err := r.db.QueryRowContext(ctx, query, email).
        Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return user, nil
}

// NewDBConnection is a placeholder function to get a database connection
func NewDBConnection() (*sql.DB, error) {
    // In a real application, you would read connection string from config
    // and handle connection pooling.
    connStr := "user=youruser dbname=yourdb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}