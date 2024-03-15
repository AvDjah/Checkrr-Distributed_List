package Models

import "gorm.io/gorm"

type EventList struct {
	gorm.Model
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Events      []Event `json:"events"`
	UserID      int64   `json:"userId"`
	PriorityID  int64   `json:"priorityID"`
}
