package models

import (
	"time"
	"tiny-tg/internal/models/types"
)

type Chat struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Type      types.Chat `json:"type"`
	OwnerId   int        `json:"owner_id"`
	MemberIds []int      `json:"member_ids" gorm:"-"`
	Info      string     `json:"info,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}
