package Models

import "gorm.io/gorm"

type Subscriptions struct {
	gorm.Model
	ID          int64 `json:"id"`
	UserID      int64 `json:"userID"`
	EventListID int64 `json:"eventListID"`
}
