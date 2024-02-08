package models

import "time"

type Player struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UUID        string    `json:"uuid" gorm:"unique"`
	Surename    string    `json:"surename"`
	Firstname   string    `json:"firstname"`
	DateOfBirth time.Time `json:"dob"`
	Number      uint8     `json:"number"`
	Position    string    `json:"position"`
	Nationality string    `json:"nationality"`
	TeamID      int       `json:"team_id"`
	Team        Team
}
