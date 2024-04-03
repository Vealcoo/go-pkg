package ginhelper

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON marshal any type with std json lib, include protobuf
func JSON(ctx *gin.Context, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, b)
	return nil
}
