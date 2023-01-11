package ginhttphelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponse(g *gin.Context, res any) {
	g.JSON(
		http.StatusOK,
		res,
	)
}

func SendError(g *gin.Context, status int, message string) {
	g.JSON(status, gin.H{
		"message": message,
	})
}
