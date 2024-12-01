package repository

import "gorm.io/gorm"

type Repo struct {
	Users *Users
	Chats *Chats
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		&Users{db},
		&Chats{db},
	}
}
