package routes

import (
	"github.com/gin-gonic/gin"
	"traceability/controllers"
	"traceability/middleware"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/upload", controllers.UploadProduct)
		protected.GET("/product/:id", controllers.GetProduct)
		protected.DELETE("/product/:id", controllers.DeleteProduct)
	}
}
