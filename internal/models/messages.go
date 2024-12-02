package models

import "time"

type Message struct {
	Id            int          `json:"id"`
	SenderId      int          `json:"sender_id"`
	ChatId        int          `json:"chat_id"`
	Text          string       `json:"text"`
	AttachmentIds []string     `json:"attachment_ids"`
	Attachments   []Attachment `json:"attachments"`
	ReplyToId     int          `json:"reply_to_id"`
	ForwardFromId int          `json:"forward_from_id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     *time.Time   `json:"deleted_at"`

	HasSeen bool `json:"has_seen"`
}
