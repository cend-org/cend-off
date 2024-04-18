package model

import "time"

type UserMark struct {
	Id            uint       `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	UserId        uint       `json:"user_id"`
	AuthorId      uint       `json:"author_id"`
	AuthorComment string     `json:"author_comment"`
	AuthorMark    uint       `json:"author_mark"`
}

type UserMarkInput struct {
	UserId        uint   `json:"user_id"`
	AuthorId      uint   `json:"author_id"`
	AuthorComment string `json:"author_comment,omitempty"`
	AuthorMark    uint   `json:"author_mark"`
}
