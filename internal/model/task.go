package model

import "time"

type Task struct {
	UUID          string    `json:"uuid"`
	Title         string    `json:"title"`
	Date          time.Time `json:"date"`
	FormattedDate string    `json:"formatted_date"`
	Description   string    `json:"description,omitempty"`
}
