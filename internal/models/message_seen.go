package models

import "time"

type MessageSeen struct {
	MessageID string    `json:"message_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
