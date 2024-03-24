package Router

import (
	"Checkrr/Handlers"
	"Checkrr/middleware"
	"github.com/gin-gonic/gin"
)

// Routes Belonging to eventList Model

func EventListRoutes(r *gin.Engine) {
	g := r.Group("/eventList")
	g.Use(middleware.ValidAuth())
	{
		g.GET("/GetAllEvents", Handlers.GetAllUserEvents)
	}
}
