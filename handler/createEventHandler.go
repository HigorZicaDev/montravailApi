package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEventHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post Event",
	})
}
