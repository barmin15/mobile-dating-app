package model

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name               string
    Email              string `gorm:"uniqueIndex"`
    Password           string
    PublicID           string `gorm:"uniqueIndex"`
    Profile            Profile
    Location           Location
    SwipedLeft         []*Swipe `gorm:"many2many:user_swipes_left;"`
    SwipedRight        []*Swipe `gorm:"many2many:user_swipes_right;"`
    Matches            []Match
}