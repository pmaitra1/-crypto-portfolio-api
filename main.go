package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    // Import auth package
    // If `auth.go` is in the same folder as `main.go`, use the package name `main`
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    InitDB() // Initialize database connection

    r := gin.Default()

    // Define routes
    r.POST("/register", RegisterUser)
    r.POST("/login", LoginUser)

    // Protected routes that require authentication
    r.GET("/portfolio/:id", JWTMiddleware(), GetPortfolio)
    r.POST("/portfolio", JWTMiddleware(), AddAsset)
    r.PUT("/portfolio/:id", JWTMiddleware(), UpdateAsset)
    r.DELETE("/portfolio/:id", JWTMiddleware(), DeleteAsset)

    // Run the server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}

