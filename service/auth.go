package service

import (
	"errors"
	"fmt"
	"testMod/config"
	"testMod/dto"
	"testMod/pkg/rabbitMq"
	"testMod/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(loginDto dto.LoginServiceDto) (error, string)
	Otp(otpDto dto.OtpService) error
}

type authService struct {
	authRepo repository.AuthRepository
	config   config.Config
	rabbitMq rabbitMq.RabbitMQ
}

func NewAuthService(authRepo repository.AuthRepository,
	 config config.Config,
	  rabbitMq rabbitMq.RabbitMQ) AuthService {
	return authService{
		authRepo: authRepo,
		config:   config,
		rabbitMq: rabbitMq,
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
		"user_id": fmt.Sprintf("%d", user.ID)   ,
		"expired_at": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.AppConfig.Database.SecretKey))
	return err, tokenString

}
func (authService authService) Otp(otpDto dto.OtpService) error {
	var otpRepo dto.OtpRepositoryDto
	otpRepo.Mobile = otpDto.Mobile
	user,err := authService.authRepo.Otp(otpRepo)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("کاربر یافت نشد")
	}
	if err!=nil {
		return err	
	}
	authService.rabbitMq.PublishMessage("otp",fmt.Sprintf("%d", user.ID))
	return err
}
