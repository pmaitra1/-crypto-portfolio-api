package main

import (
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables from .env file (if running locally)
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get the PORT from environment variables (Heroku sets this automatically)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"  // Default to 8080 if PORT is not set
    }

    // Initialize database connection
    InitDB()

    // Create a new Gin router
    r := gin.Default()

    // Define routes
    r.POST("/register", RegisterUser)
    r.POST("/login", LoginUser)

    // Protected routes that require JWT authentication
    r.GET("/portfolio/:id", JWTMiddleware(), GetPortfolio)
    r.POST("/portfolio", JWTMiddleware(), AddAsset)
    r.PUT("/portfolio/:id", JWTMiddleware(), UpdateAsset)
    r.DELETE("/portfolio/:id", JWTMiddleware(), DeleteAsset)

    // Start the Gin server, binding to the correct port
    log.Printf("Starting server on port %s", port)
    if err := r.Run(":" + port); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}
