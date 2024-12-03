package ws

import (
	"tiny-tg/internal/models"
	"tiny-tg/internal/models/types"
)

func (h *Hub) sendEventMsg(msg string, chatId int) error {
	m, err := h.serv.Messages.Create(&models.Message{Text: msg, ChatId: chatId})
	if err != nil {
		return err
	}

	m.IsEvent = true

	return h.sendMsg(m)
}

func (h *Hub) sendMsg(msg *models.Message) error {
	ids, err := h.serv.Chats.FindMemberIds(msg.ChatId)
	if err != nil {
		return err
	}

	var onlineMembers []*Client

	h.mu.Lock()
	for _, id := range ids {
		if c, ok := h.clients[id]; ok {
			onlineMembers = append(onlineMembers, c)
		}
	}
	h.mu.Unlock()

	updates := []*models.Update{{Action: types.ActionGetMessage, Message: msg}}

	for _, client := range onlineMembers {
		client.send <- updates
	}

	return nil
}

func (h *Hub) sendMsgSeen(ms *models.MessageSeen) error {
	panic("implement me")
}
