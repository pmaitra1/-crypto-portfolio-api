package main

import (
    "log"
    "os"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB initializes the PostgreSQL database connection using GORM with pgx
func InitDB() {
    // Get the DATABASE_URL environment variable
    dsn := os.Getenv("DATABASE_URL")
    
    // Check if DATABASE_URL is set and log it for debugging
    if dsn == "" {
        log.Fatalf("DATABASE_URL is not set")
    }
    log.Println("Database URL: ", dsn) // Debugging step: log the DATABASE_URL

    // Open the database connection using GORM's PostgreSQL driver
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info), // Optional: Log SQL queries
    })
    if err != nil {
        log.Fatalf("failed to connect to the database: %v", err)
    }
    log.Println("Successfully connected GORM to the database")

    // Perform the AutoMigrate for PortfolioItem and User models
    if err := DB.AutoMigrate(&User{}, &PortfolioItem{}); err != nil {
        log.Fatalf("failed to migrate database schema: %v", err)
    }
    log.Println("Database schema migrated successfully")
}
