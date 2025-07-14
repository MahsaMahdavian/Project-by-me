package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"testMod/config"
	"testMod/dto"
	"testMod/repository"
	"time"
)

type AuthService interface {
	Login(loginDto dto.LoginServiceDto) (error, string)
	Otp(otpDto dto.OtpService) error
}

type authService struct {
	authRepo repository.AuthRepository
	config   config.Config
}

func NewAuthService(authRepo repository.AuthRepository, config config.Config) AuthService {
	return authService{
		authRepo: authRepo,
		config:   config,
	}
}

func (authService authService) Login(loginDto dto.LoginServiceDto) (error, string) {
	var loginRepo dto.LoginRepositoryDto
	loginRepo.Mobile = loginDto.Mobile
	loginRepo.OtpCode = loginDto.OtpCode
	user, err := authService.authRepo.Login(loginRepo)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("کاربر یافت نشد"), ""
	}
	claims := jwt.MapClaims{
		"userId":     user.ID,
		"expired_at": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.Database.SecretKey))
	return err, tokenString

}
func (authService authService) Otp(otpDto dto.OtpService) error {
	var otpRepo dto.OtpRepositoryDto
	otpRepo.Mobile = otpDto.Mobile
	err := authService.authRepo.Otp(otpRepo)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("کاربر یافت نشد")
	}
	return err
}
