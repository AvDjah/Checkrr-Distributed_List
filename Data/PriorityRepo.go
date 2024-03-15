package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

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
