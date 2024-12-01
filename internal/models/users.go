package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	AvatarPath   string    `json:"avatar_path"`
	Bio          string    `json:"bio"`
	LastActiveAt time.Time `json:"last_active_at"`
	CreatedAt    time.Time `json:"created_at"`
}
