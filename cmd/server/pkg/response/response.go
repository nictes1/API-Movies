package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// utils
func Ok(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, Response{
		Status: http.StatusText(statusCode),
		Message: message,
		Data: data,
	})
}

func Error(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, Response{
		Status: http.StatusText(statusCode),
		Message: err.Error(),
		Data: nil,
	})
}