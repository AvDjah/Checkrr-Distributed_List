package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"fmt"
	"gorm.io/gorm"
	"math/rand"
	"time"
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
	result := db.Save(&event)
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

func GetUpcomingEvents(db *gorm.DB) ([]Models.Event, error) {
	// Get current time rounded to the nearest minute, ignoring seconds
	now := time.Now().Truncate(time.Minute)
	// Add 5 minutes to current time
	fiveMinutesFromNow := now.Add(time.Minute * 5)
	// Add 8 minutes to current time
	eightMinutesFromNow := now.Add(time.Minute * 8)

	var events []Models.Event
	// Query for events with start time between 5 and 8 minutes from now
	result := db.Where("start_time >= ? AND start_time < ?", fiveMinutesFromNow, eightMinutesFromNow).Find(&events)

	return events, result.Error
}

func InsertDummyEvents(db *gorm.DB) error {
	events := []Models.Event{}

	result := db.Create(&events)
	return result.Error
}

func generateRandomTime(minTime, maxTime time.Time) time.Time {
	// Truncate the minimum and maximum time to minutes, ignoring seconds
	minTime = minTime.Truncate(time.Minute)
	maxTime = maxTime.Truncate(time.Minute)

	// Generate a random duration between 0 and the difference between maxTime and minTime
	duration := time.Duration(rand.Intn(int(maxTime.Sub(minTime))))

	// Add the random duration to the truncated minimum time
	return minTime.Add(duration).Truncate(time.Minute)
}

func GetEventsForNextMinute(db *gorm.DB) ([]Models.Event, error) {
	// Get current local time rounded to the nearest minute, ignoring seconds
	now := time.Now().Truncate(time.Minute)
	// Add 1 minute to current time to represent the start of the next minute
	nextMinuteStart := now.Add(time.Minute)
	formattedTime := nextMinuteStart.Format("15:04:05")

	fmt.Println("Next Time: ", formattedTime)
	var events []Models.Event
	// Query for events with start time exactly equal to nextMinuteStart
	result := db.Where("start_time = ?", formattedTime).Find(&events)

	return events, result.Error
}
