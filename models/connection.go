package models

import "time"

type ConnectionType struct {
	TikTok    int
	Instagram int
}

const (
	ConnectionTypeTikTokInt    = 1
	ConnectionTypeInstagramInt = 2
)

var connectionTypes = ConnectionType{
	TikTok:    ConnectionTypeTikTokInt,
	Instagram: ConnectionTypeInstagramInt,
}


type Connection struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	ConnectionType int       `json:"connection_type"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreateConnection struct {
	UserID         int `json:"user_id" validate:"required"`
	ConnectionType string `json:"connection_type" validate:"required,oneof=tiktok instagram"`
}
