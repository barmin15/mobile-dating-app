// controller/user_controller.go

package controller

import (
	"os"
	"github.com/barmin15/mobile-dating-app/backend/initializer"
	"github.com/barmin15/mobile-dating-app/backend/model"
	"github.com/barmin15/mobile-dating-app/backend/service"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    // Middleware ensures the user is authenticated
    jwtService := service.NewJWTService(os.Getenv("SECRET_KEY"))
    jwtMiddleware := jwtService.AuthenticateTokenMiddleware()

    // Apply JWT authentication middleware
    jwtMiddleware(c)

    // Check if an error occurred during authentication
    if _, exists := c.Get("error"); exists {
        return  // Return if there's an error
    }

    // Proceed fetching users
    var users []model.User
    initializer.DB.Find(&users)

    c.JSON(200, gin.H{
        "users": users,
    })
}

