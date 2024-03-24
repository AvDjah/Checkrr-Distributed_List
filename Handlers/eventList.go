package Handlers

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handlers for eventList model

func GetAllUserEvents(c *gin.Context) {

	db := Db.InitDb()

	// Get UserId from QueryString
	userId := Helpers.GetUserIdFromQuery(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "User Does Not Exist",
		})
		return
	}
	// Get Data from Db
	eventList := Data.GetEventListsByUserId(db, userId)
	c.JSON(200, eventList)
}
