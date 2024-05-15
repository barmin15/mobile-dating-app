// router/get_routes.go
package router

import (
    "github.com/barmin15/mobile-dating-app/backend/controller"
    "github.com/gin-gonic/gin"
)

func SetupGetRoutes(r *gin.Engine) *gin.Engine {
    r.GET("/users", controller.GetUsers)
    return r
}
