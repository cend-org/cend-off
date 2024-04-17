package model

import "time"

type Address struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Country     string     `json:"country"`
	City        string     `json:"city"`
	Latitude    float64    `json:"latitude"`
	Longitude   float64    `json:"longitude"`
	Street      string     `json:"street"`
	FullAddress string     `json:"full_address"`
	Xid         string     `json:"xid"`
}

type UserAddress struct {
	Id          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UserId      uint       `json:"user_id"`
	AddressId   uint       `json:"address_id"`
	AddressType string     `json:"address_type"`
}

type NewAddress struct {
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Street      string  `json:"street"`
	FullAddress string  `json:"full_address,omitempty"`
}
