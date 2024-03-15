package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

func GetUserByUsername(db *gorm.DB, username string) Models.User {
	var user Models.User
	result := db.Where(&Models.User{UserId: username}).First(&user)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Getting user by userId")
		return user
	} else {
		return user
	}
}

func GetUserById(db *gorm.DB, id int64) Models.User {
	var user Models.User
	result := db.First(&user, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting user by id")
	}
	return user
}

func UpsertUser(db *gorm.DB, user Models.User) bool {
	result := db.Save(&user)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Upserting User")
		return false
	}
	return true
}

func DeleteUser(db *gorm.DB, id int64) bool {
	result := db.Delete(&Models.User{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting Priority by ID")
		return false
	}
	return result.RowsAffected > 0
}

func GetAllUser(db *gorm.DB) []Models.User {
	var users []Models.User
	result := db.Find(&users)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Getting All users")
	}
	return users
}
