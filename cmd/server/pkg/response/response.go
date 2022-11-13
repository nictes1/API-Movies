package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode string
	Message    string
	Data       any
}

// utils
func Ok(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, Response{
		StatusCode: http.StatusText(statusCode),
		Message: message,
		Data: data,
	})
}

func Err(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, Response{
		StatusCode: http.StatusText(statusCode),
		Message: err.Error(),
		Data: nil,
	})
}