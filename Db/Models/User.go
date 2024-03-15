package Models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id         int64       `json:"id"`
	Name       string      `json:"name"`
	UserId     string      `json:"userid"`
	Password   string      `json:"password"`
	Categories []Category  `json:"categories"`
	Lists      []EventList `json:"eventLists"`
}
