package service

import (
	"testMod/dto"
	"testMod/repository"
)
type AuthService interface {
	Login(loginDto dto.LoginServiceDto) error
	Otp(otpDto dto.OtpService) error
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(authRepo repository.AuthRepository) AuthService {
	return authService{
		authRepo: authRepo,
	}
}

func (authService authService) Login(loginDto dto.LoginServiceDto) error {
	var loginRepo dto.LoginRepositoryDto
	return authService.authRepo.Login(loginRepo)
}
func (authService authService) Otp(otpDto dto.OtpService) error {
	var otpRepo dto.OtpRepositoryDto
	return authService.authRepo.Otp(otpRepo)
}