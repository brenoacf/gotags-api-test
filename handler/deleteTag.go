package handler

import (
	"fmt"
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete Tag
// @Description Delete a tag
// @Tags Tags
// @Accept json
// @Produce json
// @Param id query string true "Tag identification"
// @Success 200 {object} DeleteTagResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /tag [delete]
func DeleteTagHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	tag := schemas.Tag{}

	// Find Tag
	if err := db.First(&tag, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("tag with id: %s not found", id))
		return
	}

	// Delete Tag
	if err := db.Delete(&tag).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting tag with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-tag", tag)
}
