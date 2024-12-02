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
			return err
		}

	}

}
