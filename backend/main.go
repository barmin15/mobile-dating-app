// main.go
package main

import (
    "github.com/barmin15/mobile-dating-app/backend/initializer"
    "github.com/barmin15/mobile-dating-app/backend/router"
)

func init() {
    initializer.LoadEnvVariables()
    initializer.ConnectToDB()
}

func main() {
    r := router.CreateRoutes()
    r.Run()
}
