package controller

import (
	"os"

	"github.com/barmin15/mobile-dating-app/backend/initializer"
	"github.com/barmin15/mobile-dating-app/backend/model"
	"github.com/barmin15/mobile-dating-app/backend/service"
	"github.com/gin-gonic/gin"
)

// RegisterUser registers a new user and generates a JWT token
func RegisterUser(c *gin.Context) {
    var user model.User

    // Bind JSON to the user struct
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Save the user to the database
    if err := initializer.DB.Create(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Generate JWT token
    jwtService := service.NewJWTService(os.Getenv("SECRET_KEY"))
    token, err := jwtService.GenerateToken(&user)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    // Send token and user data back to client
    c.JSON(201, gin.H{"token": token, "user": user})
}

func ProtectedRouteHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "You are authorized"})
}