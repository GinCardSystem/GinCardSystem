package model

import "time"

type UserModel struct {
	ID            uint `gorm:"primary"`
	UserName      string
	PasswordHash  string
	Email         string
	IsAdmin       bool
	AccountStatus bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
