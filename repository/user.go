package repository

import (
	"testMod/dto"
	"testMod/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userCreateRepository dto.UserCreateRepository) error
	Update(userUpdateRepository dto.UserUpdateRepository) error
	Delete(id uint) error
	List() ([]dto.UserGetRepository, error)
}

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) UserRepository {
	return userRepository{
		conn: conn,
	}
}
func (userRepo userRepository) Create(userCreateRepository dto.UserCreateRepository) error {
	err:=userRepo.conn.Model(&models.User{}).Create(models.User{
		FirstName:userCreateRepository.FirstName ,
		LastName: userCreateRepository.LastName,
		Email: userCreateRepository.Email,
		Age: userCreateRepository.Age,
		Mobile: userCreateRepository.Mobile,
		Gender: userCreateRepository.Gender,
		IsActive: userCreateRepository.IsActive,	
	}).Error
	
	return err
}

func (userRepo userRepository) Update(userUpdateRepository dto.UserUpdateRepository) error {

	err:=userRepo.conn.Model(&models.User{}).Where("id",userUpdateRepository.Id).Updates(map[string]interface{}{
		"firs_name":   userUpdateRepository.FirstName,
		"last-name": userUpdateRepository.LastName,
		"email":  userUpdateRepository.Email,
		"age":    userUpdateRepository.Age,
		"mobile": userUpdateRepository.Mobile,
		"gender": userUpdateRepository.Gender,
		"is_active": userUpdateRepository.IsActive,
	}).Error

	return err
}
func (userRepo userRepository) Delete(id uint) error {
	err:=userRepo.conn.Where("id",id).Delete(&models.User{}).Error
	return err
}

func (userRepo userRepository) List() ([]dto.UserGetRepository, error) {

	var users []dto.UserGetRepository
	
	err:=userRepo.conn.Model(&models.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
