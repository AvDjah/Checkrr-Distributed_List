package Handlers

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Db/Models"
	"Checkrr/Helpers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddSingleEvent(c *gin.Context) {

	var request struct {
		EventName   string `json:"eventName"`
		StartTime   string `json:"startTime"`
		WaitTime    int64  `json:"waitTime"`
		EventListID int64  `json:"eventListId"`
		Value       string `json:"message"`
		Description string `json:"description"`
		CategoryID  int64  `json:"categoryID"`
		PriorityID  string `json:"priorityID"`
		TimeUnit    string `json:"timeUnit"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		Helpers.Log(err, "Binding JSON")
		c.JSON(http.StatusBadRequest, "Bad Body Data")
		return
	}

	userId := Helpers.GetUserIdFromJWTClaim(c)
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	eventTime := Helpers.GetTimeFromUnixTime(request.StartTime)

	var event Models.Event
	{
		event.EventName = request.EventName
		event.Status = "Waiting"
		event.EventType = 1
		event.Value = request.Value
		event.WaitTime = strconv.FormatInt(request.WaitTime, 10)
		event.Description = request.Description
		event.PriorityID = Helpers.GetPriorityId(request.PriorityID)
		event.EventListID = request.EventListID
		event.StartTime = eventTime
		event.UserID = userId
		event.WaitTimeUnit = request.TimeUnit
	}

	db := Db.InitDb()
	result := Data.UpsertEvent(db, event)
	if result == true {
		c.JSON(201, "Successfully created")
	} else {
		c.JSON(http.StatusConflict, "Could not insert record")
	}
}
