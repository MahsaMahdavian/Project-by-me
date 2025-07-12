package service

import (
	"testMod/dto"
	"testMod/models"
	"testMod/repository"
)

type UserService interface {
	Create(userCreateService dto.UserCreateService) error
	Update(userUpdateService dto.UserUpdateService) error
	Delete(id uint) error
	List() ([]dto.UserGetService, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{
		userRepo: userRepo,
	}
}

func (userService userService) Create(userCreateService dto.UserCreateService) error {
	var dtoRepo dto.UserCreateRepository
	dtoRepo.FirstName = userCreateService.FirstName
	dtoRepo.LastName = userCreateService.LastName
	dtoRepo.Mobile = userCreateService.Mobile
	dtoRepo.Email = userCreateService.Email
	dtoRepo.Age = userCreateService.Age
	dtoRepo.Gender = models.Gender(userCreateService.Gender)
	dtoRepo.IsActive = userCreateService.IsActive
	err := userService.userRepo.Create(dtoRepo)
	return err
}

func (userService userService) Update(userUpdateService dto.UserUpdateService) error {
	var dtoRepo dto.UserUpdateRepository
	dtoRepo.Id = userUpdateService.Id
	dtoRepo.FirstName = userUpdateService.FirstName
	dtoRepo.LastName = userUpdateService.LastName
	dtoRepo.Mobile = userUpdateService.Mobile
	dtoRepo.Age = userUpdateService.Age
	dtoRepo.Email = userUpdateService.Email
	dtoRepo.Gender = models.Gender(userUpdateService.Gender)
	dtoRepo.IsActive = userUpdateService.IsActive
	err := userService.userRepo.Update(dtoRepo)
	return err
}

func (userService userService) Delete(id uint) error {
	err := userService.userRepo.Delete(id)
	return err
}

func (userService userService) List() ([]dto.UserGetService, error) {
	var users []dto.UserGetService
	repoList,err:=userService.userRepo.List()
	if err!=nil{
		return users,nil
	}
	for _,v:=range repoList{
		var user dto.UserGetService
		user.FirstName=v.FirstName
		user.LastName=v.LastName
		user.Mobile=v.Mobile
		user.Age=v.Age
		user.Email=v.Email
		user.Gender=string(v.Gender)
		user.IsActive=v.IsActive
		user.Id=v.Id
		users = append(users, user)
	}
	return users, nil
}
