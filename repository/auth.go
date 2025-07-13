package repository

import (
	"gorm.io/gorm"
	"testMod/dto"
)

type AuthRepository interface {
	Login(loginDto dto.LoginRepositoryDto) (error)
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


func (authRepo authRepository) Login(loginDto dto.LoginRepositoryDto) (error) {
	return nil
}

func (authRepo authRepository) Otp(otpDto dto.OtpRepositoryDto) (error) {
	return nil
}