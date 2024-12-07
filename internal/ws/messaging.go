package ws

import (
	"tiny-tg/internal/models"
	"tiny-tg/internal/models/types"
	"tiny-tg/internal/pkg/app_errors"
)

func (h *Hub) createAndSendEventMsg(msg string, chatId int) error {
	m, err := h.serv.Messages.Create(&models.Message{Text: msg, ChatId: chatId})
	if err != nil {
		return err
	}

	m.IsEvent = true

	return h.sendMsg(m)
}

func (h *Hub) sendMsg(msg *models.Message) error {
	onlineMembers, err := h.onlineMembers(msg.ChatId)
	if err != nil {
		return err
	}

	updates := []*models.Update{{Action: types.ActionGetMessage, Message: msg}}

	for _, client := range onlineMembers {
		client.send <- updates
	}

	return nil
}

func (h *Hub) sendMsgSeen(ms *models.MessageSeen) error {
	msg, err := h.serv.Messages.GetByID(ms.MessageId)
	if err != nil {
		return err
	}

	if msg.SenderId == nil {
		return app_errors.BadRequest
	}

	client, ok := h.getClient(*msg.SenderId)
	if !ok {
		return nil
	}

	updates := []*models.Update{{Action: types.ActionGetMessageSeen, MessageSeen: ms}}

	client.send <- updates

	return nil
}

func (h *Hub) onlineMembers(chatId int) ([]*Client, error) {
	ids, err := h.serv.Chats.FindMemberIds(chatId)
	if err != nil {
		return nil, err
	}

	var onlineMembers []*Client

	h.mu.Lock()
	for _, id := range ids {
		if c, ok := h.clients[id]; ok {
			onlineMembers = append(onlineMembers, c)
		}
	}
	h.mu.Unlock()

	return onlineMembers, nil
}

func (h *Hub) getClient(id int) (c *Client, ok bool) {
	h.mu.Lock()
	defer h.mu.Unlock()

	c, ok = h.clients[id]
	return
}
