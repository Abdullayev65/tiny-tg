package ws

import (
	"fmt"
	"log"
	"net/http"
	"tiny-tg/internal/models"
)

func (h *Hub) Handle(userId int, w http.ResponseWriter, r *http.Request) {

	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %s", err)
		return
	}

	client := &Client{
		conn:   conn,
		userId: userId,
		send:   make(chan []*models.Update),
	}

	h.register(client)
	defer h.unregister(client.userId)

	go h.WriteMessages(client)

	err = h.ReadMessages(client)
	if err != nil {
		fmt.Println(err)
		return
	}

}
