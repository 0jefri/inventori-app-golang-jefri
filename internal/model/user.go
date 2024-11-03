package model

import (
	"time"
)

type User struct {
	ID          string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	Username    string `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required,alphanum"`
	Email       string `gorm:"type:varchar(255);not null;unique" json:"email,omitempty" binding:"required,email"`
	Password    string `gorm:"type:varchar(255);not null" json:"password,omitempty" binding:"required,alphanum"`
	FirstName   string `gorm:"type:varchar(255);not null" json:"firstName,omitempty" binding:"required,alpha"`
	LastName    string `gorm:"type:varchar(255);not null" json:"lastName,omitempty" binding:"required,alpha"`
	PhoneNumber string `gorm:"type:varchar(255);not null;unique" json:"phoneNumber,omitempty" binding:"required,numeric"`
	// Wallet           Wallet    `gorm:"references:ID" json:"wallet"`
	RegistrationDate time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"registrationDate" binding:"omitempty"`
	// ProfilePicture   []byte    `json:"profilePicture" binding:"omitempty"`
	LastLogin time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"lastLogin" binding:"omitempty"`
}
