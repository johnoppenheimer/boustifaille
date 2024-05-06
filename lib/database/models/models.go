package models

import "time"

type User struct {
	Name string
	ID   uint
}

type Place struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Slug      string `gorm:"uniqueIndex"`
	User      User
	ID        uint
	UserID    int
	Latitude  float64 `gorm:"index"`
	Longitude float64 `gorm:"index"`
}
