package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListEventsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Get Events",
	})
}
