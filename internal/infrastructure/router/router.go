package router

import (
    "github.com/gofiber/fiber/v2"
    "your-module-name/internal/auth/handler"
)

// SetupAuthRoutes sets up the authentication routes
func SetupAuthRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
    auth := app.Group("/api/v1/auth")

    auth.Post("/register", authHandler.Register)
    auth.Post("/login", authHandler.Login)

    // Example of a protected route using a hypothetical middleware
    // protected := app.Group("/api/v1/protected").Use(middleware.AuthMiddleware())
    // protected.Get("/profile", handler.GetProfile)
}