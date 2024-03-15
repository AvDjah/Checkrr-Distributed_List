package Data

import (
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"gorm.io/gorm"
)

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
