package Handlers

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Helpers"
	"github.com/gin-gonic/gin"
)

func GetAllHandler(c *gin.Context) {

	userId := Helpers.GetUserIdFromJWTClaim(c)
	if userId == 0 {
		c.JSON(401, "Not Logged In, Redirecting")
		return
	}

	db := Db.InitDb()
	result := Data.GetNotificationsByUserID(db, userId)

	c.JSON(200, result)

}
