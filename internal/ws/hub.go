package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"tiny-tg/internal/models"
	"tiny-tg/internal/service"
)

func NewHub(serv *service.Service) *Hub {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return &Hub{
		serv:      serv,
		clients:   make(map[int]*Client),
		upgrader:  upgrader,
		broadcast: make(chan []*models.Update),
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
