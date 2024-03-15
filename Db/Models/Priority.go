package Models

import "gorm.io/gorm"

type Priority struct {
	gorm.Model
	Id        int64       `json:"id"`
	Name      string      `json:"name"`
	Order     string      `json:"order"`
	Color     string      `json:"color"`
	Events    []Event     `json:"events"`
	EventList []EventList `json:"eventList"`
}
