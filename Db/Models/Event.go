package Models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID           int64        `json:"id"`
	EventName    string       `json:"eventName"`
	StartTime    string       `json:"startTime"`
	WaitTime     string       `json:"waitTime"`
	WaitTimeUnit string       `json:"waitTimeUnit"`
	Status       string       `json:"status"`
	EventType    int          `json:"eventType"`
	EventListID  int64        `json:"eventListId"`
	Value        string       `json:"value"`
	Description  string       `json:"description"`
	CategoryID   int64        `json:"categoryID"`
	PriorityID   int64        `json:"priorityID"`
	UserID       int64        `json:"userID"`
	Notification Notification `json:"notification"`
}
