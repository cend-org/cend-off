package model

import "time"

type QrCodeRegistry struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UserId    uint       `json:"user_id"`
	Xid       string     `json:"xid"`
	IsUsed    bool       `json:"is_used"`
}
