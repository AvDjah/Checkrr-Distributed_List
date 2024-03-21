package Router

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Helpers"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	AddRoutes(r)
	return r
}

func AddRoutes(router *gin.Engine) {

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"result": "Heelo World",
		})
	})

	router.GET("/check", func(c *gin.Context) {
		db := Db.InitDb()
		data, err := Data.GetEventsForNextMinute(db)
		Helpers.Log(err, "Getting Upcoming Events")
		c.JSON(200, gin.H{
			"result": len(data),
		})
	})

	router.GET("/insert", func(context *gin.Context) {

		db := Db.InitDb()
		err := Data.InsertDummyEvents(db)
		if err != nil {
			Helpers.Log(err, "Inserting Dummy Events")
			context.JSON(200, gin.H{
				"result": "failure",
			})
			return
		}

		context.JSON(200, gin.H{
			"result": "success",
		})
	})

}
