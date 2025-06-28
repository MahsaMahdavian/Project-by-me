package service

import (
	"testMod/dto"
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
	dtoRepo.Age = userCreateService.Age
	dtoRepo.Email = userCreateService.Email
	dtoRepo.Name = userCreateService.Name
	dtoRepo.Family = userCreateService.Family
	err := userService.userRepo.Create(dtoRepo)
	return err
}

func (userService userService) Update(userUpdateService dto.UserUpdateService) error {
	var dtoRepo dto.UserUpdateRepository
	dtoRepo.Age = userUpdateService.Age
	dtoRepo.Email = userUpdateService.Email
	dtoRepo.Name = userUpdateService.Name
	dtoRepo.Family = userUpdateService.Family
	dtoRepo.Id=userUpdateService.Id
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
		user.Age=v.Age
		user.Email=v.Email
		user.Family=v.Family
		user.Name=v.Name
		user.Id=v.Id
		users = append(users, user)
	}
	return users, nil
}
