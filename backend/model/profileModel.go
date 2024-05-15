package model

import (
	"errors"
	"gorm.io/gorm"
)

type Gender string

const (
    Male   Gender = "male"
    Female Gender = "female"
    Other  Gender = "other"
)

type Profile struct {
    gorm.Model
    UserID    uint   `gorm:"unique"`
    Bio       string
    Age                uint
    Gender             Gender
    DatingPreference   string
    PartnersCount      uint
    Photos    []string `gorm:"type:json"`
    Interests []string `gorm:"type:json"`
}

func (g Gender) Validate() error {
    if g != Male && g != Female && g != Other {
        return errors.New("invalid gender")
    }
    return nil
}

func (p *Profile) BeforeSave(tx *gorm.DB) (err error) {
    if err := p.Gender.Validate(); err != nil {
        return err
    }
    return nil
}

