package main

import (
	initializers "github.com/barmin15/mobile-dating-app/backend/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Run()
}