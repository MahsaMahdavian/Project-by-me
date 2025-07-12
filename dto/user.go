package dto

import "testMod/models"

type UserRequest struct {
	Id        uint          `query:"id" form:"id"`
	FirstName string        `query:"first_name" form:"first_name" validate:"required"`
	LastName  string        `query:"last_name" form:"last_name" validate:"required"`
	Email     *string       `query:"email" form:"email" validate:"email"`
	Mobile    string        `query:"mobile" form:"mobile" validate:"required"`
	Age       uint          `query:"age" form:"age" validate:"required"`
	Gender    models.Gender `query:"gender" form:"gender" validate:"required,oneof=male female"`
	IsActive  bool          `query:"is_active" form:"is_active"`
}

type UserResponse struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Age       uint   `json:"age"`
	Gender    string `json:"gender"`
	IsActive  bool   `json:"is_active"`
}

type UserCreateRepository struct {
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	Age       uint
	Gender    models.Gender
	IsActive  bool
}

type UserUpdateRepository struct {
	Id        uint
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	Age       uint
	Gender    models.Gender
	IsActive  bool
}
type UserGetRepository struct {
	Id        uint
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	Age       uint
	Gender    models.Gender
	IsActive  bool
}

type UserCreateService struct {
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	Age       uint
	Gender    models.Gender
	IsActive  bool
}

type UserUpdateService struct {
	Id        uint
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	Age       uint
	Gender    string
	IsActive  bool
}
type UserGetService struct {
	Id        uint
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	Age       uint
	Gender    string
	IsActive  bool
}
