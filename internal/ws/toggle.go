package ws

import "github.com/gorilla/websocket"

func (h *Hub) register(client *Client) {
	h.mu.Lock()

	h.clients[client.userId] = client

	h.mu.Unlock()
}

func (h *Hub) unregister(userId int) {
	client, ok := h.clients[userId]
	if !ok {
		return
	}

	h.mu.Lock()
	delete(h.clients, userId)
	h.mu.Unlock()

	close(client.send)
	_ = client.conn.WriteMessage(websocket.CloseMessage, []byte{})
	_ = client.conn.Close()

}
