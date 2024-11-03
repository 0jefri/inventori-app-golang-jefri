package dto

import (
	"time"
	// "github.com/sakupay-apps/internal/model"
)

type UserResponse struct {
	ID               string    `json:"id,omitempty"`
	Username         string    `json:"username,omitempty"`
	Email            string    `json:"email,omitempty"`
	Password         string    `json:"password,omitempty"`
	FirstName        string    `json:"firstName,omitempty"`
	LastName         string    `json:"lastName,omitempty"`
	PhoneNumber      string    `json:"phoneNumber,omitempty"`
	RegistrationDate time.Time `json:"registrationDate,omitempty"`
	// ProfilePicture   []byte          `json:"profilePicture,omitempty"`
	LastLogin time.Time `json:"lastLogin,omitempty"`
	// Wallet           model.Wallet    `json:"wallet,omitempty"`
	// Cards            []model.Card    `json:"cards,omitempty"`
	// Contacts         []model.Contact `json:"contacts,omitempty"`
}
