package dto

type UserRequest struct {
	Id    uint `query:"id" form:"id"`
	Name   string `query:"name" form:"name" validate:"required"`
	Family string `query:"family" form:"family"`
	Age    uint `query:"age" form:"age"`
	Email  string `query:"email" form:"email" validate:"required,email"`
}

type UserResponse struct {
	Id     uint
	Name   string
	Family string
	Age    uint
	Email  string
}

type UserCreateRepository struct {
	Name   string
	Family string
	Age    uint
	Email  string
}

type UserUpdateRepository struct {
	Id     uint
	Name   string
	Family string
	Age    uint
	Email  string
}
type UserGetRepository struct {
	Id     uint
	Name   string
	Family string
	Age    uint
	Email  string
}

type UserCreateService struct {
	Name   string
	Family string
	Age    uint
	Email  string
}

type UserUpdateService struct {
	Id     uint
	Name   string
	Family string
	Age    uint
	Email  string
}
type UserGetService struct {
	Id     uint
	Name   string
	Family string
	Age    uint
	Email  string
}