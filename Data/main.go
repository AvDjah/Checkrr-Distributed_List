package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

// EventList Model

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

// User Model

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

// Event Model

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

func GetEventsByEventListId(db *gorm.DB, eventListId int64) []Models.Event {
	var events []Models.Event
	result := db.Where(&Models.Event{EventListID: eventListId}).Find(&events)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error getting EventsByEventResultId")
	}
	return events
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

// Category Model

func GetCategories(db *gorm.DB) []Models.Category {
	var categories []Models.Category

	result := db.Find(&categories)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error querying: GetCategories")
		return []Models.Category{}
	}
	return categories
}

func GetCategoryById(db *gorm.DB, id int64) Models.Category {
	var category Models.Category

	result := db.First(&category, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error querying: CategoryById")
		return Models.Category{}
	}
	return category
}

func UpsertCategory(db *gorm.DB, category Models.Category) bool {
	result := db.Save(&category)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error upserting Category")
		return false
	}
	return true
}

func DeleteCategoryById(db *gorm.DB, id int64) bool {
	result := db.Delete(&Models.Category{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting Category by ID")
		return false
	}
	return result.RowsAffected > 0 // Check if a row was actually deleted
}

// Priority Model

func GetPriorities(db *gorm.DB) []Models.Priority {
	var priorities []Models.Priority

	result := db.Find(&priorities)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error querying: GetPriorities")
		return []Models.Priority{}
	}
	return priorities
}

func GetPriorityById(db *gorm.DB, id int64) Models.Priority {
	var priority Models.Priority

	result := db.First(&priority, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error querying: GetPriorityById")
		return Models.Priority{}
	}
	return priority
}

func UpsertPriority(db *gorm.DB, priority Models.Priority) bool {
	result := db.Save(&priority)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error upserting Priority")
		return false
	}
	return true
}

func DeletePriorityById(db *gorm.DB, id int64) bool {
	result := db.Delete(&Models.Priority{}, id)
	if result.Error != nil {
		Helpers.Log(result.Error, "Error deleting Priority by ID")
		return false
	}
	return result.RowsAffected > 0 // Check if a row was actually deleted
}
