package router

import (
	"montravail/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowEventHandler)
		v1.POST("/opening", handler.CreateEventHandler)
		v1.DELETE("/opening", handler.DeleteEventHandler)
		v1.PUT("/opening", handler.UpdateEventHandler)
		v1.GET("/openings", handler.ListEventsHandler)
	}
}
