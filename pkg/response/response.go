package response

import (
	"base-gin-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(ctx *gin.Context, data interface{}) {
	response(ctx, http.StatusOK, data)
}

func WithStatusCode(ctx *gin.Context, code int, data interface{}) {
	response(ctx, code, data)
}

func response(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, models.Response{
		Message: http.StatusText(code),
		Data:    data,
	})
}
