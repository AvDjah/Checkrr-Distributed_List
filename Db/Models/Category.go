package Models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	UserID      int64  `json:"userId"`
}
