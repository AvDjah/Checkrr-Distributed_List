package Router

import (
	"Checkrr/Handlers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(g *gin.Engine) {

	group := g.Group("/event")
	{
		group.POST("/addSingleEvent", Handlers.AddSingleEvent)
	}

}
