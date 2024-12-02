package repository

import "gorm.io/gorm"

type Repo struct {
	Users    *Users
	Chats    *Chats
	Messages *Messages
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		&Users{db},
		&Chats{db},
		&Messages{db},
	}
}
