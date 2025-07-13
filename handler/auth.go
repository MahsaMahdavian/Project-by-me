package handler

import (
	"testMod/service"

	"github.com/gin-gonic/gin"
)
type AuthHandler interface{

	Otp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authHandler struct{
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService)AuthHandler{
	return authHandler{
		authService: authService,
	}
}

func (authHandler authHandler)Otp(ctx *gin.Context){

}
func (authHandler authHandler)Login(ctx *gin.Context){

}