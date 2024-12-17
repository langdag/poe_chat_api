package models

type RegisterUser struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginUser struct{
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type DefaultUser struct{
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}