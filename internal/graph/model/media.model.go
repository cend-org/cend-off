package model

import "time"

const (
	CV                = 0
	CoverLetter       = 1
	PresentationVideo = 2
	UserProfileImage  = 3
)

type Media struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	FileName    string     `json:"file_name"`
	Extension   string     `json:"extension"`
	Xid         string     `json:"xid"`
	ContentType uint       `json:"content_type"`
}

type UserMediaDetail struct {
	Id           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	OwnerId      uint       `json:"owner_id"`
	DocumentType uint       `json:"document_type"`
}
