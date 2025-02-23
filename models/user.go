package models

import "time"

type ConnectionType string

const (
	ConnectionTypeTelegram  ConnectionType = "telegram"
	ConnectionTypeTikTok    ConnectionType = "tiktok"
	ConnectionTypeDefault   ConnectionType = "pending"
)

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
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	ConnectionType ConnectionType `json:"connection_type"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Me struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}