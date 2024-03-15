package Router

import "github.com/gin-gonic/gin"

func eventListRoutes(router *gin.Engine) {

	router.GET("/GetUserList", func(c *gin.Context) {

		// Get Query Params
		_ = c.Query("userId")

	})

}
