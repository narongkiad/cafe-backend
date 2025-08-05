package handler

import (
    "your-module-name/internal/auth/usecase"
    "github.com/gofiber/fiber/v2"
)

// AuthHandler holds the dependencies for the auth handler
type AuthHandler struct {
    authUsecase usecase.AuthUsecase
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(uc usecase.AuthUsecase) *AuthHandler {
    return &AuthHandler{
        authUsecase: uc,
    }
}

// RegisterRequest is the request body for the register endpoint
type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}

// Register handles the user registration
func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var req RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    // TODO: Add validation using a library like go-playground/validator

    user, err := h.authUsecase.Register(c.Context(), req.Email, req.Password)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "User created successfully",
        "user":    user,
    })
}

// LoginRequest is the request body for the login endpoint
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

// Login handles user login and returns a JWT token
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
    }

    // TODO: Add validation

    token, err := h.authUsecase.Login(c.Context(), req.Email, req.Password)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{"token": token})
}