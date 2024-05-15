package controller

import (
	"github.com/barmin15/mobile-dating-app/backend/initializer"
	"github.com/barmin15/mobile-dating-app/backend/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
    var users []model.User

    initializer.DB.Find(&users)

    c.JSON(200, gin.H{
        "users": users,
    })
}
