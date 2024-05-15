package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/barmin15/mobile-dating-app/backend/model"
    "github.com/barmin15/mobile-dating-app/backend/initializer"
)

func PostUser(c *gin.Context) {
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

    c.JSON(201, gin.H{"user": user})
}
