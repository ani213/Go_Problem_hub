package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func NewCustomError(statusCode int, message string, details string) CustomError {
	return CustomError{
		StatusCode: statusCode,

		Message: message,
		Details: details,
	}
}

func SomethingWentWrong(err error) CustomError {
	return NewCustomError(
		http.StatusInternalServerError, // 500
		"Something went wrong",
		err.Error(),
	)
}

func SendError(c *gin.Context, customErr CustomError) {
	c.JSON(customErr.StatusCode, customErr)
}
