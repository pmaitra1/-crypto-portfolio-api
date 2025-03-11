package main

import (
	"gorm.io/gorm"
)

// PortfolioItem represents an asset in the portfolio
type PortfolioItem struct {
	gorm.Model         // Automatically adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
	Price      float64 `json:"price"`
	UserID     uint    `json:"user_id"`
}

func (PortfolioItem) TableName() string {
	return "portfolio" // Set the table name explicitly to "portfolio"
}
