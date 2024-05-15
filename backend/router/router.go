package router

import "github.com/gin-gonic/gin"

func CreateRoutes() *gin.Engine {
    r := gin.Default()

    r = SetupGetRoutes(r)
    r = SetupPostRoutes(r)

    return r
}
