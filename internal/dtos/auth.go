package dtos

import "tiny-tg/internal/models"

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRes struct {
	User  *models.User `json:"user"`
	Token string       `json:"token"`
}
