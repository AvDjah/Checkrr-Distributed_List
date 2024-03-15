package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

func GetEventLists(db *gorm.DB) []Models.EventList {

	var eventList []Models.EventList

	result := db.Find(&eventList)
	if result.Error != nil {

		Helpers.Log(result.Error, "Error querying: GetUserEventLists")

		return eventList
	} else {
		return eventList
	}
}

func GetEventListById(db *gorm.DB, id int64) Models.EventList {
	var eventList Models.EventList

	result := db.First(&eventList, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error querying: GetEventListById")
		return Models.EventList{}
	} else {
		return eventList
	}
}

func GetEventListsByUserId(db *gorm.DB, userId int64) []Models.EventList {
	var eventLists []Models.EventList
	result := db.Where(&Models.EventList{UserID: userId}).Find(&eventLists)

	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting EventListsByUserId")
	}

	return eventLists
}

func UpsertEventList(db *gorm.DB, eventList Models.EventList) bool {
	result := db.Save(&eventList)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Inserting EventList")
		return false
	} else {
		return true
	}
}

func DeleteEventListById(db *gorm.DB, id int64) bool {
	result := db.Delete(&Models.EventList{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting Priority by ID")
		return false
	}
	return result.RowsAffected > 0 // Check if a row was actually deleted
}
