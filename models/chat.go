package models

import "time"

type Chat struct {
	ID         int     `json:"id"`
	SenderID   int     `json:"sender_id"`
	ReceiverID int     `json:"receiver_id"`
	Message    string  `json:"message"`
	CreatedAt  string  `json:"created_at"`
}

type CreateChat struct {
	SenderID   int     `json:"sender_id" validate:"required"`
	ReceiverID int     `json:"receiver_id" validate:"required"`
	Message    string  `json:"message" validate:"required"`
}

type EstablishedConnection struct {
	ID              int       `json:"id"`
    WebConnectionID int       `json:"web_connection_id"`
	CreatedAt       time.Time `json:"created_at"`
}

type ChannelRegistrationStatus struct {
    Registered bool `json:"registered"`
}


