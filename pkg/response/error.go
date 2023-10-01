package response

import (
	"base-gin-go/models"
	"base-gin-go/pkg/apierrors"
	"base-gin-go/pkg/i18n"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(ctx *gin.Context, err error) {
	var debugMsg string
	if gin.Mode() != gin.ReleaseMode {
		debugMsg = err.Error()
	}
	errType := apierrors.ErrType(err)

	ctx.JSON(errType.HTTPCode(), models.ErrorResponse{
		Code:         errType.Code(),
		DebugMessage: debugMsg,
		ErrorDetails: errorDetails(err, ctx.Query("locale")),
	})
}

func errorDetails(err error, locale string) (details []models.ErrorDetail) {
	var vErrs validator.ValidationErrors
	if errors.As(err, &vErrs) {
		fmt.Println("locale: ", locale)
		trans := i18n.GetTrans(locale)
		for _, err := range vErrs {
			details = append(details, models.ErrorDetail{
				Field:     err.Field(),
				ErrorCode: err.Tag(),
				// ErrorMessage: err.Error(),
				ErrorMessage: err.Translate(trans),
			})
		}
	}
	return
}
