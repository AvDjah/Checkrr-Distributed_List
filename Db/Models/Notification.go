package Models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ID               int64  `json:"id"`
	UserID           int64  `json:"userID"`
	Status           int64  `json:"status"`
	EventID          int64  `json:"eventID"`
	NotificationType string `json:"notificationType"`
	Message          string `json:"message"`
}
