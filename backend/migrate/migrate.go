package main

import (
    "github.com/barmin15/mobile-dating-app/backend/initializer"
    "github.com/barmin15/mobile-dating-app/backend/model"
)

func init() {
    initializer.LoadEnvVariables()
    initializer.ConnectToDB()
}

func main() {
    db := initializer.DB

//users, location, swipes

    // Auto-migrate the models
    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.Location{})
    db.AutoMigrate(&model.Profile{})
    db.AutoMigrate(&model.Match{})
}
