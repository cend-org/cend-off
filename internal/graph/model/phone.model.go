package model

import "time"

type PhoneNumber struct {
	Id                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `json:"deleted_at"`
	MobilePhoneNumber string     `json:"mobile_phone_number"`
	IsUrgency         bool       `json:"is_urgency"`
}

type UserPhoneNumber struct {
	Id            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	UserId        uint       `json:"user_id"`
	PhoneNumberId uint       `json:"phone_number_id"`
}

type NewPhoneNumber struct {
	MobilePhoneNumber string `json:"mobile_phone_number"`
	IsUrgency         bool   `json:"is_urgency,omitempty"`
}
