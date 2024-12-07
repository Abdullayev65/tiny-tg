package ws

func (h *Hub) register(client *Client) {
	h.mu.Lock()

	prev, ok := h.clients[client.userId]
	h.clients[client.userId] = client

	h.mu.Unlock()

	if ok {
		prev.GoClose()
	}

}

func (h *Hub) unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	client, ok := h.clients[c.userId]
	if !ok || client != c {
		return
	}

	delete(h.clients, c.userId)

	client.GoClose()

}
