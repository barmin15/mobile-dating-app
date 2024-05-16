// router/post_routes.go
package router

import (
    "github.com/barmin15/mobile-dating-app/backend/controller"
    "github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine) *gin.Engine {
    r.POST("/register/", controller.RegisterUser)
    return r
}
