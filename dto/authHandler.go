package dto

type AuthLoginRequest struct {
	Mobile string `query:"mobile" form:"mobile" validate:"required,mobile"`
	OtpCode string `query:"otp_code" form:"otp_code" validate:"required"`
}

type AuthOtpRequest struct {
	Mobile string `query:"mobile" form:"mobile" validate:"required,mobile"`
}
type AuthResponse struct {
}
