package model

import (
    "gorm.io/gorm"
    "time"
)

type Match struct {
    gorm.Model
    UserID             uint
    User               User
    MatchedUserID      uint
    MatchedUser        User `gorm:"foreignKey:MatchedUserID"`
    MatchedAt          time.Time
    CompatibilityScore int
}
