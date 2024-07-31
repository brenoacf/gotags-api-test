package handler

import (
	"fmt"
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

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
