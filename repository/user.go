package repository

import (
	"testMod/dto"
	"testMod/models"
	"testMod/utils"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userCreateRepository dto.UserCreateRepository) (int64, error)
	Update(userUpdateRepository dto.UserUpdateRepository) (int64, error)
	Delete(id uint) (int64, error)
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
func (userRepo userRepository) Create(userCreateRepository dto.UserCreateRepository) (int64, error) {
	res := userRepo.conn.Model(&models.User{}).Create(&models.User{
		FirstName: userCreateRepository.FirstName,
		LastName:  userCreateRepository.LastName,
		Email:     userCreateRepository.Email,
		Age:       userCreateRepository.Age,
		Mobile:    userCreateRepository.Mobile,
		Gender:    userCreateRepository.Gender,
		IsActive:  userCreateRepository.IsActive,
		OtpCode: uint(utils.GenerateRandomNumber()),
		OtpCodeExpiredAt: time.Now().Add(time.Minute*1),
	})

	err := res.Error
	rowAffected := res.RowsAffected

	if err != nil {
		return rowAffected, err
	}
	return rowAffected, nil
}

func (userRepo userRepository) Update(userUpdateRepository dto.UserUpdateRepository) (int64, error) {

	res := userRepo.conn.Model(&models.User{}).Where("id", userUpdateRepository.Id).Updates(map[string]interface{}{
		"first_name": userUpdateRepository.FirstName,
		"last_name": userUpdateRepository.LastName,
		"email":     userUpdateRepository.Email,
		"age":       userUpdateRepository.Age,
		"mobile":    userUpdateRepository.Mobile,
		"gender":    userUpdateRepository.Gender,
		"is_active": userUpdateRepository.IsActive,
	})

	err := res.Error
	rowAffected := res.RowsAffected
	if err != nil {
		return rowAffected, err
	}
	return rowAffected, err
}
func (userRepo userRepository) Delete(id uint) (int64, error) {
	res := userRepo.conn.Where("id", id).Delete(&models.User{})
	err := res.Error
	rowAffected := res.RowsAffected
	if err != nil {
		return rowAffected, err
	}
	return rowAffected, err
}

func (userRepo userRepository) List() ([]dto.UserGetRepository, error) {

	var users []dto.UserGetRepository

	err := userRepo.conn.Model(&models.User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
