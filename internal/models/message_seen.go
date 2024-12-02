package models

import "time"

type MessageSeen struct {
	MessageId int       `json:"message_id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
