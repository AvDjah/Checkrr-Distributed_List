package Models

import (
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	ID           int64        `json:"id"`
	EventName    string       `json:"eventName"`
	StartTime    time.Time    `json:"startTime"`
	WaitTime     time.Time    `json:"waitTime"`
	Status       string       `json:"status"`
	EventType    int          `json:"eventType"`
	EventListID  int64        `json:"eventListId"`
	Value        float64      `json:"value"`
	Description  string       `json:"description"`
	CategoryID   int64        `json:"categoryID"`
	PriorityID   int64        `json:"priorityID"`
	UserID       int64        `json:"userID"`
	Notification Notification `json:"notification"`
}
