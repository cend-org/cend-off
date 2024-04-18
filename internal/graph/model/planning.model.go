package model

import "time"

type CalendarPlanning struct {
	Id              uint       `json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	AuthorizationId uint       `json:"authorization_id"`
	StartDateTime   time.Time  `json:"start_date_time"`
	EndDateTime     time.Time  `json:"end_date_time"`
	Description     string     `json:"description"`
}

type CalendarPlanningActor struct {
	Id                 uint       `json:"id"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
	AuthorizationId    uint       `json:"authorization_id"`
	CalendarPlanningId uint       `json:"calendar_planning_id"`
}

type NewCalendarPlanning struct {
	StartDateTime time.Time `json:"start_date_time"`
	EndDateTime   time.Time `json:"end_date_time"`
	Description   string    `json:"description"`
}
