package main

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "log" 
    "net/http"
    "time"
)

// User represents a user in the system
type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
}

func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
            c.Abort()
            return
        }

        tokenString = tokenString[7:]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, nil
            }
            return []byte("your_secret_key"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || claims["user_id"] == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
            c.Abort()
            return
        }

        // Debug log: Verify if user_id is extracted properly
        log.Println("User ID extracted from token:", claims["user_id"])

        c.Set("user_id", claims["user_id"])
        c.Next()
    }
}


// RegisterUser handles user registration
func RegisterUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid data format"})
        return
    }

    if user.Username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Username is required"})
        return
    }
    if user.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Password is required"})
        return
    }

    // Hash the password before storing it
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }

    user.Password = string(hashedPassword)

    // Save the user to the database
    result := DB.Create(&user)
    if result.Error != nil {
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(200, gin.H{"message": "User registered successfully"})
}


// LoginUser handles user login and generates JWT
func LoginUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": "Invalid data format"})
        return
    }

    // Find the user in the database
    var dbUser User
    result := DB.Where("username = ?", user.Username).First(&dbUser)
    if result.Error != nil {
        c.JSON(401, gin.H{"error": "Invalid username or password"})
        return
    }

    // Check if the password is correct
    err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
    if err != nil {
        c.JSON(401, gin.H{"error": "Invalid username or password"})
        return
    }

    // Generate JWT token, now including user_id
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": dbUser.Username,
        "user_id": dbUser.ID,  // Add the user_id to the token
        "exp":      time.Now().Add(time.Minute * 30).Unix(),
    })

    tokenString, err := token.SignedString([]byte("your_secret_key"))
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{
        "token":   tokenString,
        "user_id": dbUser.ID,  // Include user ID in the response
    })
}
