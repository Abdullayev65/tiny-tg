package models

import "time"

type Message struct {
	Id            int        `json:"id"`
	SenderId      *int       `json:"sender_id,omitempty"`
	ChatId        int        `json:"chat_id"`
	Text          string     `json:"text"`
	ReplyToId     *int       `json:"reply_to_id,omitempty"`
	ForwardFromId *int       `json:"forward_from_id,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`

	//AttachmentIds []string     `json:"attachment_ids" gorm:"-"`
	Attachments []Attachment `json:"attachments,omitempty" gorm:"-"`
	HasSeen     bool         `json:"has_seen" gorm:"-"`
	IsEvent     bool         `json:"is_event" gorm:"-"`
}
