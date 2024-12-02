package service

import (
	"tiny-tg/internal/pkg/jwt_manager"
	"tiny-tg/internal/repository"
)

type Service struct {
	Auth     *Auth
	Chats    *Chats
	Messages *Messages
}

func New(repo *repository.Repo, jwtManager *jwt_manager.JwtManager) *Service {
	return &Service{
		&Auth{repo, jwtManager},
		&Chats{repo},
		&Messages{repo},
	}
}
