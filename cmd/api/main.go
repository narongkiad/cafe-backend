package main

import (
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
    "your-module-name/internal/auth/handler"
    "your-module-name/internal/auth/usecase"
    "your-module-name/internal/infrastructure/database/postgres"
    "your-module-name/internal/infrastructure/router"
)

func main() {
    app := fiber.New()

    // 1. Setup Database Connection (mocked for simplicity)
    dbConn, err := postgres.NewDBConnection() // Replace with your actual DB connection logic
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    // 2. Instantiate Repositories
    authRepo := postgres.NewAuthRepository(dbConn)

    // 3. Instantiate Usecases
    timeout := time.Duration(10) * time.Second
    authUsecase := usecase.NewAuthUsecase(authRepo, timeout)

    // 4. Instantiate Handlers
    authHandler := handler.NewAuthHandler(authUsecase)

    // 5. Setup Router
    router.SetupAuthRoutes(app, authHandler)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Fatal(app.Listen(":" + port))
}