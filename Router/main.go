package Router

import "github.com/gin-gonic/gin"

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
}
