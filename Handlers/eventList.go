package Handlers

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handlers for eventList model

func GetAllUserEvents(c *gin.Context) {

	db := Db.InitDb()

	// Get UserId from JWT
	userId := Helpers.GetUserIdFromJWTClaim(c)

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

func AddUserEvent(c *gin.Context) {

	userId := Helpers.GetUserIdFromJWTClaim(c)
	if userId == 0 {
		c.JSON(http.StatusBadRequest, "User is not logged in.")
		return
	}
	// Body Data struct
	var data struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Priority    string `json:"priority" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, "Bad Body Data")
		return
	}

	var eventList Models.EventList
	eventList.UserID = userId
	eventList.Name = data.Name
	eventList.Description = data.Description
	eventList.PriorityID = Helpers.GetPriorityId(data.Priority)

	if eventList.PriorityID == 0 {
		c.JSON(400, "Bad Body Data")
		return
	}

	db := Db.InitDb()
	result := Data.UpsertEventList(db, eventList)
	if result == false {
		c.JSON(400, "Could not Insert Data in DB")
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
		"id":     eventList,
	})
}
