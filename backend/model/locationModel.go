package model

import "gorm.io/gorm"

type Location struct {
    gorm.Model
    UserID    uint    `gorm:"unique"`
    Latitude  float64
    Longitude float64
}

func (l *Location) GetUser(db *gorm.DB) (*User, error) {
    var user User
    if err := db.First(&user, l.UserID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
