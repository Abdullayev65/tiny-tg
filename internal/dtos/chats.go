package dtos

import "tiny-tg/internal/models"

type ChatColl struct {
	Users  []models.User `json:"users"`
	Groups []models.Chat `json:"groups"`
}
