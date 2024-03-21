package Router

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Helpers"
	"github.com/gin-gonic/gin"
)

func eventListRoutes(router *gin.Engine) {

	router.GET("/GetUserList", func(c *gin.Context) {

		// Get Query Params
		_ = c.Query("userId")

	})

	router.GET("/check", func(c *gin.Context) {
		db := Db.InitDb()
		data, err := Data.GetEventsForNextMinute(db)
		Helpers.Log(err, "Getting Upcoming Events")
		c.JSON(200, gin.H{
			"result": len(data),
		})
	})

}
