package handler

import (
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateTagHandler(ctx *gin.Context) {
	request := UpdateTagRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	// Update Tag
	if request.Code > 0 {
		tag.Code = request.Code
	}

	if request.Tag != "" {
		tag.Tag = request.Tag
	}

	if request.Date != "" {
		tag.Date = request.Date
	}

	if request.Hour != "" {
		tag.Hour = request.Hour
	}

	if request.Value > 0 {
		tag.Value = request.Value
	}

	// Save tag
	if err := db.Save(&tag).Error; err != nil {
		logger.Errorf("error updating tag: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating tag")
		return
	}

	sendSuccess(ctx, "update-tag", tag)

}
