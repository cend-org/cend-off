package model

import "time"

type UserAuthorizationLink struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	LinkType  uint       `json:"link_type"`
}

type UserAuthorizationLinkActor struct {
	Id                      uint       `json:"id"`
	CreatedAt               time.Time  `json:"created_at"`
	UpdatedAt               time.Time  `json:"updated_at"`
	DeletedAt               *time.Time `json:"deleted_at"`
	UserAuthorizationLinkId uint       `json:"user_authorization_link_id"`
	AuthorizationId         uint       `json:"authorization_id"`
}
