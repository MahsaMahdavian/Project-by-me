package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testMod/dto"
	"testMod/service"
	"testMod/validator"
)

type AuthHandler interface {
	Otp(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return authHandler{
		authService: authService,
	}
}

func (authHandler authHandler) Otp(ctx *gin.Context) {
	var request dto.AuthOtpRequest
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	validator := validator.NewValidator()
	if !validator.Validator(&request) {
		ctx.JSON(http.StatusBadRequest, validator.GetErrors())
		return
	}

	var authOtpDtoService dto.OtpService
	authOtpDtoService.Mobile = request.Mobile

	err = authHandler.authService.Otp(authOtpDtoService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "کد ورود ارسال شد",
	})
}
func (authHandler authHandler) Login(ctx *gin.Context) {
var request dto.AuthLoginRequest
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	validator := validator.NewValidator()
	if !validator.Validator(&request) {
		ctx.JSON(http.StatusBadRequest, validator.GetErrors())
		return
	}

	var authLoginDtoServic dto.LoginServiceDto
	authLoginDtoServic.Mobile = request.Mobile
	authLoginDtoServic.OtpCode=request.OtpCode

	err,token:=authHandler.authService.Login(authLoginDtoServic)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"JWT": token,
	})
}
