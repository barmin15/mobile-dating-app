package model

import "gorm.io/gorm"

type Swipe struct {
    gorm.Model
    UserID uint
    SwipedUserID uint
}
