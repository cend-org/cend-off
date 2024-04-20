package model

import "time"

type Message struct {
	Id               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	ResourceType     int        `json:"resource_type"`
	ResourceNumber   int        `json:"resource_number"`
	ResourceValue    int        `json:"resource_value"`
	ResourceLabel    string     `json:"resource_label"`
	ResourceLanguage int        `json:"resource_language"`
}

type MessageInput struct {
	ResourceType     int    `json:"resource_type"`
	ResourceNumber   int    `json:"resource_number"`
	ResourceValue    int    `json:"resource_value"`
	ResourceLabel    string `json:"resource_label"`
	ResourceLanguage int    `json:"resource_language"`
}

type MessageUpdateInput struct {
	Id               uint   `json:"id"`
	ResourceType     int    `json:"resource_type"`
	ResourceNumber   int    `json:"resource_number"`
	ResourceValue    int    `json:"resource_value"`
	ResourceLabel    string `json:"resource_label"`
	ResourceLanguage int    `json:"resource_language"`
}
