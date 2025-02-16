package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(err error, validationMessages map[string]string, c *gin.Context) {
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				// Lookup error message or use default
				message := validationMessages[fe.Tag()]
				if message == "" {
					message = "Invalid value!"
				}
				out[i] = ApiError{Field: fe.Field(), Message: message}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
	}
}
