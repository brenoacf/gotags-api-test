package router

import (
	docs "github.com/brenoacf/gotags-api-test/docs"
	"github.com/brenoacf/gotags-api-test/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	// Docs
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group("/api/v1")
	{
		v1.GET("/tag", handler.ShowTagHandler)
		v1.POST("/tag", handler.CreateTagHandler)
		v1.DELETE("/tag", handler.DeleteTagHandler)
		v1.PUT("/tag", handler.UpdateTagHandler)
		v1.GET("/tags", handler.ListTagsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
