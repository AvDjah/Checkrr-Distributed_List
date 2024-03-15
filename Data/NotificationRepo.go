package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

func GetAllNotifications(db *gorm.DB) []Models.Notification {
	var notifications []Models.Notification
	result := db.Find(&notifications)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting all notifications")
		return notifications
	}
	return notifications
}

func GetNotificationsByUserID(db *gorm.DB, userId int64) []Models.Notification {
	var notifications []Models.Notification
	result := db.Where(&Models.Notification{UserID: userId}).Find(&notifications)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting notifications by userID")
	}
	return notifications
}

func GetNotificationByID(db *gorm.DB, id int64) (Models.Notification, bool) {
	var notification Models.Notification
	result := db.Find(&notification, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting notification by ID")
		return notification, false
	}
	return notification, true
}

func UpsertNotification(db *gorm.DB, notification Models.Notification) bool {
	result := db.Save(notification)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error upserting notification")
		return false
	}
	return true
}

func DeleteNotification(db *gorm.DB, id int64) (bool, error) {
	result := db.Delete(&Models.Notification{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting notification by ID")
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
