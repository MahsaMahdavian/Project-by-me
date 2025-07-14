package repository

import (
	"testMod/dto"
	"testMod/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(loginDto dto.LoginRepositoryDto) (models.User,error)
	Otp(otpDto dto.OtpRepositoryDto) (error)

}

type authRepository struct {
	conn *gorm.DB
}
func NewAuthRepository(conn *gorm.DB) AuthRepository {
	return authRepository{
		conn: conn,
	}
}


func (authRepo authRepository) Login(loginDto dto.LoginRepositoryDto) (models.User,error) {
var user models.User
	err:=authRepo.conn.Where("mobile",loginDto.Mobile).Where("otp_code",loginDto.OtpCode).First(&user).Error
	return user,err
}

func (authRepo authRepository) Otp(otpDto dto.OtpRepositoryDto) (error) {

	var user models.User
	err:=authRepo.conn.Where("mobile",otpDto.Mobile).First(&user).Error
	return err
}