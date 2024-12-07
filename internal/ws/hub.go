package ws

import (
	"fmt"
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

func (h *Hub) ReadMessages(client *Client) error {

	for {
		update := new(models.Update)
		err := client.conn.ReadJSON(update)
		if err != nil {
			return fmt.Errorf("error reading message: %s", err)
		}

		if update.Message != nil {
			for _, att := range update.Message.Attachments {
				if att.Id > 0 {
					continue
				}

				att.FilePath, err = uploadFile(client.conn, att.Size, att.MimeType)
				if err != nil {
					return err
				}

			}
		}

		update.FromUserId = client.userId
		err = h.update(update)
		if err != nil {
			fmt.Println(err)
			//return err
		}

	}

}

func (h *Hub) WriteMessages(client *Client) {

	for {
		updates, ok := <-client.send
		if !ok {
			return
		}

		err := client.conn.WriteJSON(updates)
		if err != nil {
			fmt.Println(client)
			fmt.Println(updates)
			fmt.Println(err)
			return
		}

	}

}
