package ws

import "fmt"

func (h *Hub) WriteMessages(client *Client) {

	for {
		updates := <-client.send
		err := client.conn.WriteJSON(updates)
		if err != nil {
			fmt.Println(client)
			fmt.Println(updates)
			fmt.Println(err)
		}

	}

}
