package Router

import (
	"Checkrr/Handlers"
	"github.com/gin-gonic/gin"
)

func NotificationRoutes(r *gin.Engine) {
	group := r.Group("/notification")
	{
		group.GET("/GetAll", Handlers.GetAllHandler)
	}
}
