package ws

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"os"
	"path"
	"tiny-tg/internal/config"
)

func uploadFile(conn *websocket.Conn, size int, mimeType string) (string, error) {
	filePath := path.Join(config.UPLOADS_DIR, uuid.NewString()+"."+mimeType)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	for size > 0 {

		mt, message, err := conn.ReadMessage()
		if err != nil {
			return "", err
		}

		if mt != websocket.BinaryMessage {
			return "", fmt.Errorf("expected %d{BinaryMessage}, got %d", websocket.BinaryMessage, mt)
		}

		_, err = file.Write(message)
		if err != nil {
			return "", err
		}

		size -= len(message)

		//conn.,.requestNextBlock()
	}

	file.Close()

	return filePath, nil
}
