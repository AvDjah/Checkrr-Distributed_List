package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id            int64           `json:"id"`
	Name          string          `json:"name"`
	UserId        string          `json:"userid" gorm:"uniqueIndex"`
	Password      string          `json:"password"`
	Categories    []Category      `json:"categories"`
	EventLists    []EventList     `json:"eventLists"`
	Events        []Event         `json:"Events"`
	Notifications []Notification  `json:"notifications"`
	Subscriptions []Subscriptions `json:"subscriptions"`
}

type UserResponse struct {
	Id            int64           `json:"id"`
	Name          string          `json:"name"`
	UserId        string          `json:"userId"`
	Categories    []Category      `json:"categories"`
	EventLists    []EventList     `json:"eventLists"`
	Events        []Event         `json:"Events"`
	Notifications []Notification  `json:"notifications"`
	Subscriptions []Subscriptions `json:"subscriptions"`
}
