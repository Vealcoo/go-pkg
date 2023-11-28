package ginhttphelper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendOKResponse(g *gin.Context, res any) {
	g.JSON(http.StatusOK, res)
}

var errMap map[error]int

func SetErrorMap(m map[error]int) {
	errMap = m
}

func SendErrorResponse(g *gin.Context, err error) {
	g.Error(err)
	if code, ok := errMap[err]; ok {
		g.JSON(code, gin.H{
			"message":    err.Error(),
			"statusText": http.StatusText(code),
		})
	} else {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message":    "ServerError",
			"statusText": http.StatusText(http.StatusInternalServerError),
		})
	}
}
