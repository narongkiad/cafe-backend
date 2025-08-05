package middleware

import (
    "strings"
    "github.com/narongkiad/cafe-backend/pkg/jwt"

    "github.com/gofiber/fiber/v2"
)

// AuthMiddleware is a Fiber middleware to check for a valid JWT token
func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Authorization header is required",
            })
        }

        // The token is typically in the format "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Authorization header format must be 'Bearer <token>'",
            })
        }

        tokenString := parts[1]
        claims, err := jwt.ParseToken(tokenString)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "error": "Invalid or expired token",
            })
        }

        // Store the user ID in the Fiber context for subsequent handlers
        c.Locals("userID", claims.UserID)

        // Continue to the next middleware or handler
        return c.Next()
    }
}