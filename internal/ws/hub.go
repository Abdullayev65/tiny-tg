package ws

import (
	"github.com/gorilla/websocket"
	"sync"
	"tiny-tg/internal/models"
	"tiny-tg/internal/service"
)

func NewHub(serv *service.Service) *Hub {
	return &Hub{
		serv: serv,
		//broadcast:  make(chan []byte),
		//clients:    make(map[*Client]bool),
	}
}

type Hub struct {
	serv      *service.Service
	clients   map[int]*Client
	upgrader  *websocket.Upgrader
	broadcast chan []*models.Update
	mu        sync.Mutex
}

type Client struct {
	userId int
	conn   *websocket.Conn
	send   chan []*models.Update
}

func (h *Hub) Start() {
	go h.Run()
}

func (h *Hub) Run() {
	for {
		select {
		case message := <-h.broadcast:
			for _, client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client.userId)
				}
			}
		}
	}
}
