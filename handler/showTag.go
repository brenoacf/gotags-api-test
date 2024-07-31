package handler

import (
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show Tag
// @Description Show a tag
// @Tags Tags
// @Accept json
// @Produce json
// @Param id query string true "Tag identification"
// @Success 200 {object} ShowTagResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /tag [get]
func ShowTagHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	tag := schemas.Tag{}

	if err := db.First(&tag, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "tag not found")
		return
	}

	sendSuccess(ctx, "show-tag", tag)
}
