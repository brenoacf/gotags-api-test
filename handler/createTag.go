package handler

import (
	"net/http"

	"github.com/brenoacf/gotags-api-test/schemas"
	"github.com/gin-gonic/gin"
)

func CreateTagHandler(ctx *gin.Context) {
	request := CreateTagRequest{}

	ctx.BindJSON(&request)

	if err := request.validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tag := schemas.Tag{
		Code:  request.Code,
		Tag:   request.Tag,
		Date:  request.Date,
		Hour:  request.Hour,
		Value: request.Value,
	}

	if err := db.Create(&tag).Error; err != nil {
		logger.Errorf("error creating tag: %+v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating tag on database")
		return
	}

	sendSuccess(ctx, "create-tags", tag)
}
