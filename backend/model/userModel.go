package model

import (
    "gorm.io/gorm"
    "github.com/google/uuid"
)

type User struct {
    gorm.Model
    Email              string `gorm:"uniqueIndex"`
    Password           string
    PublicID           string `gorm:"uniqueIndex"`
    Profile            Profile
    Location           Location
    SwipedLeft         []*Swipe `gorm:"many2many:user_swipes_left;"`
    SwipedRight        []*Swipe `gorm:"many2many:user_swipes_right;"`
    Matches            []Match
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
    user.PublicID = uuid.New().String()
    return nil
}
