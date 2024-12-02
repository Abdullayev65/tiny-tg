package ws

import (
	"errors"
	"fmt"
	"tiny-tg/internal/models"
	"tiny-tg/internal/models/types"
)

func (h *Hub) update(update *models.Update) error {
	var err error

	switch update.Action {
	case types.ActionCreateGroup:
		err = h.updateCreateGroup(update)
	case types.ActionJoinGroup:
		err = h.updateJoinGroup(update)
	case types.ActionLiveGroup:
		err = h.updateLiveGroup(update)

	case types.ActionSendMessage:
		err = h.updateSendMessage(update)
	case types.ActionEditMessage:
		err = h.updateEditMessage(update)
	case types.ActionDeleteMessage:
		err = h.updateDeleteMessage(update)

	case types.ActionMessageSeen:
		err = h.updateMessageSeen(update)

	default:
		err = errors.New("unknown action")
	}

	return err
}

func (h *Hub) updateCreateGroup(update *models.Update) error {
	if update.Group == nil {
		return errors.New("group cannot be null")
	}

	g := update.Group
	g.Type = types.ChatGroup
	g.OwnerId = update.FromUserId

	g, err := h.serv.Chats.Create(update.Group)
	if err != nil {
		return err
	}

	err = h.serv.Chats.CreateMembers(g.Id, []int{g.OwnerId})
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(`@%d created a group "%s"`, g.OwnerId, g.Name)
	err = h.sendEventMsg(msg, g.Id)
	if err != nil {
		return err
	}

	return nil
}

func (h *Hub) updateJoinGroup(update *models.Update) error { return nil }

func (h *Hub) updateLiveGroup(update *models.Update) error {
	return nil
}

func (h *Hub) updateSendMessage(update *models.Update) error {
	return nil
}

func (h *Hub) updateEditMessage(update *models.Update) error {
	return nil
}

func (h *Hub) updateDeleteMessage(update *models.Update) error {
	return nil
}

func (h *Hub) updateMessageSeen(update *models.Update) error {
	return nil
}
