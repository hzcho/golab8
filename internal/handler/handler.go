package handler

import (
	_ "golab8/docs"
	"golab8/internal/group"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Users API
// @version         1.0
// @description     This is an API server for working with users
//
//	@host			localhost:8080

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func InitRoutes(router *gin.Engine, groups group.Groups) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", groups.Auth.Register)
			auth.POST("/login", groups.Auth.SignIn)
		}
		users := api.Group("/users", groups.Middleware.AccountIdentity())
		{
			users.GET("/", groups.User.Get)
			users.GET("/:id", groups.User.GetById)
			users.POST("/", groups.User.Post)
			users.PUT("/:id", groups.User.Put)
			users.DELETE("/:id", groups.User.Delete)
		}
		admins := api.Group("/admins", groups.Middleware.AccountIdentity())
		{
			admins.GET("/", groups.Admin.Get)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
