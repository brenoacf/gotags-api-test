package handler

import (
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

func ListTagsHandler(ctx *gin.Context) {
	tags := []schemas.Tag{}

	if err := db.Find(&tags).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing tags")
		return
	}

	sendSuccess(ctx, "list-tags", tags)
}
