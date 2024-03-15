package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

func GetAllSubscriptions(db *gorm.DB) []Models.Subscriptions {
	var subscriptions []Models.Subscriptions
	result := db.Find(&subscriptions)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting all subscriptions")
		return subscriptions
	}
	return subscriptions
}

func GetSubscriptionsByUserID(db *gorm.DB, userId int64) []Models.Subscriptions {
	var subscriptions []Models.Subscriptions
	result := db.Where(&Models.Subscriptions{UserID: userId}).Find(&subscriptions)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting subscriptions by userID")
	}
	return subscriptions
}

func GetSubscriptionByID(db *gorm.DB, id int64) (Models.Subscriptions, bool) {
	var subscription Models.Subscriptions
	result := db.Find(&subscription, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting subscription by ID")
		return subscription, false
	}
	return subscription, true
}

func UpsertSubscription(db *gorm.DB, subscription Models.Subscriptions) bool {
	result := db.Save(subscription)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error upserting subscription")
		return false
	}
	return true
}

func DeleteSubscription(db *gorm.DB, id int64) (bool, error) {
	result := db.Delete(&Models.Subscriptions{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting subscription by ID")
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
