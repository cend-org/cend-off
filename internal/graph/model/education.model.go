package model

import "time"

type Education struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
}

type Subject struct {
	Id               uint       `json:"id"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	EducationLevelId uint       `json:"education_level_id"`
	Name             string     `json:"name"`
}

type UserEducationLevelSubject struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UserId    uint       `json:"user_id"`
	SubjectId uint       `json:"subject_id"`
}

type SubjectInput struct {
	Id               uint   `json:"id,omitempty"`
	EducationLevelId uint   `json:"education_level_id"`
	Name             string `json:"name"`
}
