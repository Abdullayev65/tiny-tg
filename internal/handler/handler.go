package handler

import (
	"tiny-tg/internal/pkg/jwt_manager"
	"tiny-tg/internal/service"
	"tiny-tg/internal/ws"
)

type Handler struct {
	service    *service.Service
	jwtManager *jwt_manager.JwtManager
	wsHub      *ws.Hub
}

func New(service *service.Service, manager *jwt_manager.JwtManager, wsHub *ws.Hub) *Handler {
	return &Handler{
		service:    service,
		jwtManager: manager,
		wsHub:      wsHub,
	}
}
