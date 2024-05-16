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
    UserID            uint   `gorm:"unique"`
    Name               string
    Bio               string
    Age               uint
    Gender            Gender
    DatingPreference  Gender
    PartnersCount     uint
    Photos            [3]string `gorm:"type:json"`
    Interests         []string `gorm:"type:json"`
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
    if err := p.DatingPreference.Validate(); err != nil {
        return err
    }
    return nil
}