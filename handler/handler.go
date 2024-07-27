package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowTagHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Tag",
	})
}

func CreateTagHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Tag",
	})
}

func DeleteTagHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Tag",
	})
}

func PutTagHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Tag",
	})
}

func ListTagsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Tags",
	})
}
