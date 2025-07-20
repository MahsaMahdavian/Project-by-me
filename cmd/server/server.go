package server

import (
	"testMod/config"
	"testMod/handler"
	"github.com/gin-gonic/gin"
	"testMod/middleware"
)

func StartServer(config config.Config, userhandler handler.Userhandler,authHandler handler.AuthHandler) error {

	r := gin.Default()
	userGroup := r.Group("/users")
	{
		userGroup.Use(middleware.AuthMiddleware(config))
		userGroup.Use(middleware.LoggingMiddleware())
		userGroup.GET("list", userhandler.List)
		userGroup.DELETE("delete/:id", userhandler.Delete)
		userGroup.POST("create", userhandler.Create)
		userGroup.PUT("update", userhandler.Update)
	}

		authGroup := r.Group("/auth")
	{
	
		authGroup.POST("otp", authHandler.Otp)
		authGroup.POST("login", authHandler.Login)
	}

	err := r.Run(config.AppPort)
	return err
}


