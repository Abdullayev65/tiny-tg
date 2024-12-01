package ws

import (
	"fmt"
	"tiny-tg/internal/models"
)

func (h *Hub) ReadMessages(client *Client) error {

	for {
		update := new(models.Update)
		err := client.conn.ReadJSON(update)
		if err != nil {
			return fmt.Errorf("error reading message: %s", err)
		}

		update.FromUserId = client.userId
		h.update(update)

	}

}
