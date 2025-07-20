package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"testMod/dto"
	"testMod/service"
	"testMod/validator"
	"github.com/gin-gonic/gin"
)

type userhandler struct {
	userService service.UserService
}

type Userhandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

func NewUserHandler(userService service.UserService) Userhandler {
	return userhandler{
		userService: userService,
	}
}

func (userhandler userhandler) Create(ctx *gin.Context) {
	var request dto.UserRequest


	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validator:=validator.NewValidator()
	if !validator.Validator(&request){
		ctx.JSON(http.StatusBadRequest,validator.GetErrors())
		return
	}
	var userCreateServiceDto dto.UserCreateService
	userCreateServiceDto.FirstName = request.FirstName
	userCreateServiceDto.LastName = request.LastName
	userCreateServiceDto.Mobile = request.Mobile
	userCreateServiceDto.Gender= request.Gender
	userCreateServiceDto.Age = request.Age
	userCreateServiceDto.IsActive = request.IsActive
	if request.Email != nil {
		userCreateServiceDto.Email = request.Email
	} else {
		userCreateServiceDto.Email = nil
	}

	rowAffected,err := userhandler.userService.Create(userCreateServiceDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"rowAffected": rowAffected,
	})

}

func (h userhandler) Update(ctx *gin.Context){
	var request dto.UserRequest
	err := ctx.Bind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var userUpdateService dto.UserUpdateService
	userUpdateService.Age = request.Age
	userUpdateService.FirstName = request.FirstName
	userUpdateService.LastName = request.LastName
	userUpdateService.Mobile = request.Mobile
	userUpdateService.Email = request.Email
	userUpdateService.Gender=string(request.Gender)
	userUpdateService.IsActive = request.IsActive
	userUpdateService.Id = request.Id

	rowAffected,err :=h.userService.Update(userUpdateService)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"rowAffected": rowAffected,
	})

}

func (userhandler userhandler) Delete(ctx *gin.Context) {
	userId := ctx.Param("id")
	if userId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "missing userid",
		})
	}
	id64, err := strconv.ParseUint(userId, 10, 30)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	rowAffected,err := userhandler.userService.Delete(uint(id64))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"rowAffected": rowAffected,
	})
}

func (userhandler userhandler) List(ctx *gin.Context) {
	userId:=ctx.GetString("user_id")
	fmt.Println(userId)
	rows, err := userhandler.userService.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, rows)
}
