package Services

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Db/Models"
	"time"
)

func RunWorker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Run The Two Consumers

	for {
		select {
		case <-ticker.C:
			if time.Now().Second() == 55 {
				go worker()
			}
		}
	}
}

func worker() {
	db := Db.InitDb()
	events := Data.GetAllEvents(db)
	var triggeredEvents []Models.Event
	for _, event := range events {
		// Add a check to get the event which we are going to happen
		triggeredEvents = append(triggeredEvents, event)
	}
	// Send Triggered Notifications to be notified to Users in Real Time

	// Send Triggered Notifications to be stored in DB

}
