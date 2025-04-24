package models

import "time"

const (
	ConnectionTypeTikTokInt    = 1
	ConnectionTypeInstagramInt = 2
)

var ConnectionTypes = map[string]int{
    "tiktok":    ConnectionTypeTikTokInt,
    "instagram": ConnectionTypeInstagramInt,
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
