package Helpers

import "github.com/gin-gonic/gin"

type Request struct {
	UserId int64 `form:"userId"`
}

func GetUserIdFromQuery(c *gin.Context) int64 {
	var request Request

	err := c.ShouldBind(&request)
	if err != nil {
		Log(err, "Binding Query String")
		return 0
	} else {
		return request.UserId
	}
}
