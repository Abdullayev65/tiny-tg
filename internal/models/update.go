package models

import "tiny-tg/internal/models/types"

type Update struct {
	FromUserId  int          `json:"-"`
	Action      types.Action `json:"action"`
	RelatedId   int          `json:"relatedId,omitempty"`
	Group       *Chat        `json:"group,omitempty"`
	Message     *Message     `json:"message,omitempty"`
	MessageSeen *MessageSeen `json:"message_seen,omitempty"`
}
