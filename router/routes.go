package router

import (
	"github.com/brenoacf/gotags-api-test/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/tag", handler.ShowTagHandler)
		v1.POST("/tag", handler.CreateTagHandler)
		v1.DELETE("/tag", handler.DeleteTagHandler)
		v1.PUT("/tag", handler.PutTagHandler)
		v1.GET("/tags", handler.ListTagsHandler)
	}

}