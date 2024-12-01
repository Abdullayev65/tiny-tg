package models

import "time"

type Message struct {
	ID            string       `json:"id"`
	SenderID      string       `json:"sender_id"`
	ChatID        string       `json:"chat_id"`
	Text          string       `json:"text"`
	AttachmentIDs []string     `json:"attachment_ids"`
	Attachments   []Attachment `json:"attachments"`
	ReplyToID     string       `json:"reply_to_id"`
	ForwardFromID string       `json:"forward_from_id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     *time.Time   `json:"deleted_at"`

	HasSeen bool `json:"has_seen"`
}
