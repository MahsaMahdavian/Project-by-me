package handler

import (
	"net/http"
	"strconv"
	"testMod/dto"
	"testMod/service"
	// "testMod/utils"
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
	var userCreateServiceDto dto.UserCreateService
	userCreateServiceDto.Age = request.Age
	userCreateServiceDto.Name = request.Name
	userCreateServiceDto.Family = request.Family
	userCreateServiceDto.Email = request.Email

	err = userhandler.userService.Create(userCreateServiceDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"request": request,
	})

}

func (userhandler userhandler) Update(ctx *gin.Context) {
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
	userUpdateService.Name = request.Name
	userUpdateService.Family = request.Family
	userUpdateService.Email = request.Email
	userUpdateService.Id = request.Id

	err =userhandler. userService.Update(userUpdateService)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"request": request,
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
	err = userhandler.userService.Delete(uint(id64))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "delete affected",
	})
}

func (userhandler userhandler) List(ctx *gin.Context) {
	rows, err := userhandler.userService.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, rows)
}
