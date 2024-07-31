package handler

import (
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

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
