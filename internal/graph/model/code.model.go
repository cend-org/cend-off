package model

import "time"

type Code struct {
	Id               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	UserId           uint       `json:"user_id"`
	VerificationCode int        `json:"value"`
}
