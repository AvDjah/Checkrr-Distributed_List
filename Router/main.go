package Router

import (
	"Checkrr/Data"
	"Checkrr/Db"
	"Checkrr/Helpers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true

	r.Use(cors.New(config))

	//r.Use(gin.Recovery())

	AddRoutes(r)
	UserGroup(r)
	DebugRoutes(r)
	EventListRoutes(r)

	return r
}

func DebugRoutes(r *gin.Engine) {
	g := r.Group("/debug")
	{
		g.GET("/check", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"result": "success",
			})
		})
	}
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

func UserGroup(router *gin.Engine) {
	g := router.Group("/user")
	{
		g.POST("/authenticate", LoginUser)
		g.GET("/details", GetUserDetails)
		g.POST("/register", RegisterUser)
	}
}

func EventsGroup(router *gin.Engine) {
	g := router.Group("/event")
	{
		g.GET("/all")
		g.GET("/alladmin")
		g.GET("/")
	}
}
