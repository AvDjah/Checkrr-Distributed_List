package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

func GetEventsByEventListId(db *gorm.DB, eventListId int64) []Models.Event {
	var events []Models.Event
	result := db.Where(&Models.Event{EventListID: eventListId}).Find(&events)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting EventsByEventResultId")
	}
	return events
}
func GetAllEvents(db *gorm.DB) []Models.Event {
	var events []Models.Event
	result := db.Find(&events)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting All Events")
		return events
	}
	return events
}

func GetEventByUserId(db *gorm.DB, userId int64) []Models.Event {
	var events []Models.Event
	result := db.Where(&Models.Event{UserID: userId}).Find(&events)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Getting EventsByUserId")
	}
	return events
}

func GetEventById(db *gorm.DB, id int64) Models.Event {
	var event Models.Event
	result := db.Find(event, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Getting eventbyID")
	}
	return event
}

func UpsertEvent(db *gorm.DB, event Models.Event) bool {
	result := db.Save(event)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error Inserting User Event")
		return false
	} else {
		return true
	}
}

func DeleteEvent(db *gorm.DB, id int64) bool {
	result := db.Delete(&Models.Event{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting Category by ID")
		return false
	}
	return result.RowsAffected > 0
}
