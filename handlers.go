package main

import (
    "strings"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetPortfolio fetches all the portfolio items
func GetPortfolio(c *gin.Context) {
    // Get the asset ID from the URL
    id := c.Param("id")

    // Get the user_id from the JWT token
    userID := c.MustGet("user_id").(float64)

    // Declare a variable to hold the portfolio item
    var portfolio PortfolioItem

    // Query the database for the asset by its ID
    result := DB.First(&portfolio, id)
    if result.Error != nil {
        // If an error occurs (e.g., asset not found)
        log.Println("Error fetching portfolio item:", result.Error)
        c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
        return
    }

    // Check if the asset belongs to the authenticated user
    if portfolio.UserID != uint(userID) {
        c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this asset"})
        return
    }

    // Return the asset
    c.JSON(http.StatusOK, portfolio)
}



// AddAsset adds a new asset to the portfolio
func AddAsset(c *gin.Context) {
    userID := c.MustGet("user_id").(float64) // Extract user_id from the context

    var newAsset PortfolioItem
    // Bind the incoming JSON payload to the newAsset variable
    if err := c.ShouldBindJSON(&newAsset); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
        return
    }
    // Check if the user_id provided in the asset is the same as the user_id from the JWT token
    if newAsset.UserID != uint(userID) {
        log.Println("Error: User ID in the payload does not match the authenticated user")
        c.JSON(http.StatusForbidden, gin.H{"error": "User ID in the payload does not match the authenticated user"})
        return
    }

    newAsset.Name = strings.ToLower(newAsset.Name)
    // Validate that name and amount are not empty
    if newAsset.Name == "" {
        log.Println("Error: Name is required")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
        return
    }
    if newAsset.Amount <= 0 {
        log.Println("Error: Amount must be greater than 0")
        c.JSON(http.StatusBadRequest, gin.H{"error": "Amount must be a positive number"})
        return
    }

    // Get the current price of the asset from CoinGecko
    price, err := GetCurrentPrice(newAsset.Name)
    if err != nil {
        log.Println("Error fetching asset price:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get asset price"})
        return
    }

    newAsset.Price = price
    newAsset.UserID = uint(userID) // Set the user_id for the asset

    // Save the asset in the database
    result := DB.Create(&newAsset)
    if result.Error != nil {
        log.Println("Error saving asset:", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save asset"})
        return
    }

    // Return the added asset
    c.JSON(http.StatusOK, newAsset)
}



// DeleteAsset deletes an asset from the portfolio by ID
func DeleteAsset(c *gin.Context) {
    id := c.Param("id")
    if _, err := strconv.Atoi(id); err != nil {
        log.Println("Error: Invalid asset ID format:", id)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID format"})
        return
    }

    userID := c.MustGet("user_id").(float64) // Get the user_id from the JWT token

    var asset PortfolioItem
    result := DB.First(&asset, id)
    if result.Error != nil {
        log.Println("Error fetching asset:", result.Error)
        c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
        return
    }

    // Check if the asset belongs to the authenticated user
    if asset.UserID != uint(userID) {
        c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to delete this asset"})
        return
    }

    result = DB.Delete(&asset)
    if result.Error != nil {
        log.Println("Error deleting asset:", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete asset"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Asset deleted"})
}

// UpdateAsset updates an existing asset's details
// Update an existing asset in the portfolio
func UpdateAsset(c *gin.Context) {
    id := c.Param("id")

    // Validate that id is a valid number
    if _, err := strconv.Atoi(id); err != nil {
        log.Println("Error: Invalid asset ID format:", id)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid asset ID format"})
        return
    }

    userID := c.MustGet("user_id").(float64) // Get the user_id from the JWT token

    var updatedAsset PortfolioItem
    if err := c.ShouldBindJSON(&updatedAsset); err != nil {
        log.Println("Error binding JSON:", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
        return
    }

    // Find the asset by ID
    var asset PortfolioItem
    result := DB.First(&asset, id)
    if result.Error != nil {
        log.Println("Error fetching asset:", result.Error)
        c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
        return
    }

    // Check if the asset belongs to the authenticated user
    if asset.UserID != uint(userID) {
        c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this asset"})
        return
    }

    // Only allow the user to update price and amount
    if updatedAsset.Price != 0 {
        asset.Price = updatedAsset.Price
    }
    if updatedAsset.Amount != 0 {
        asset.Amount = updatedAsset.Amount
    }

    // Don't allow the user to change the asset's name
    if updatedAsset.Name != "" && updatedAsset.Name != asset.Name {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot change the name of the asset"})
        return
    }

    // Save the updated asset back to the database
    result = DB.Save(&asset)
    if result.Error != nil {
        log.Println("Error updating asset:", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update asset"})
        return
    }

    c.JSON(http.StatusOK, asset)
}
