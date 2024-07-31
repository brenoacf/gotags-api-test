package handler

import (
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List Tags
// @Description List all tags
// @Tags Tags
// @Accept json
// @Produce json
// @Success 200 {object} ListTagsResponse
// @Failure 500 {object} ErrorResponse
// @Router /tags [get]
func ListTagsHandler(ctx *gin.Context) {
	tags := []schemas.Tag{}

	if err := db.Find(&tags).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing tags")
		return
	}

	sendSuccess(ctx, "list-tags", tags)
}
