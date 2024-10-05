package handler

import (
	"golab8/internal/group"

	_ "golab8/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Building API
// @version         1.0
// @description     This is an API server for working with buildings
//
//	@host			localhost:8080
func InitRoutes(router *gin.Engine, groups group.Groups) {
	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/", groups.User.Get)
			users.GET("/:id", groups.User.GetById)
			users.POST("/", groups.User.Post)
			users.PUT("/:id", groups.User.Put)
			users.DELETE("/:id", groups.User.Delete)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
