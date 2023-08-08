package helpers

import (
	"github.com/gin-gonic/gin"
)

type messageType string

type jsonResponse struct {
	Status  bool         `json:"status"`
	Data    any          `json:"data,omitempty"`
	Message *messageType `json:"message"`
}

func response(c *gin.Context, statusCode int, data *jsonResponse) {
	c.JSON(statusCode, data)
}

func SuccessResponse(c *gin.Context, statusCode int, data any, message string) {
	tmp := messageType(message)
	msg := &tmp
	if message == "" {
		msg = nil
	}
	jsonRes := new(jsonResponse)
	jsonRes.Status = true
	jsonRes.Data = data
	jsonRes.Message = msg
	response(c, statusCode, jsonRes)
}

func FailedResponse(c *gin.Context, statusCode int, data any, message string) {
	tmp := messageType(message)
	msg := &tmp
	if message == "" {
		msg = nil
	}
	jsonRes := new(jsonResponse)
	jsonRes.Status = false
	jsonRes.Data = data
	jsonRes.Message = msg
	response(c, statusCode, jsonRes)
}
