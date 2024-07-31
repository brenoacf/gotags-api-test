package handler

import (
	"fmt"
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("tag from handler %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateTagResponse struct {
	Message string              `json:"message"`
	Data    schemas.TagResponse `json:"data"`
}

type DeleteTagResponse struct {
	Message string              `json:"message"`
	Data    schemas.TagResponse `json:"data"`
}

type ShowTagResponse struct {
	Message string              `json:"message"`
	Data    schemas.TagResponse `json:"data"`
}
