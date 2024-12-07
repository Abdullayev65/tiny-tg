package service

import (
	"slices"
	"time"
	"tiny-tg/internal/models"
	"tiny-tg/internal/models/types"
	"tiny-tg/internal/pkg/app_errors"
	"tiny-tg/internal/repository"
)

type Messages struct {
	Repo *repository.Repo
}

func (s *Messages) Create(m *models.Message) (*models.Message, error) {
	if m.Text == "" {
		m.Text = "It is " + time.Now().Format("Jan 02 15:04:05")
	}

	if m.ChatId == 0 {
		return nil, app_errors.BadRequest
	}

	memberIds, err := s.Repo.Chats.FindMemberIds(m.ChatId)
	if err != nil {
		return nil, err
	}

	if m.SenderId != nil && !slices.Contains(memberIds, *m.SenderId) {
		return nil, app_errors.AccessDenied

	}

	m, err = s.Repo.Messages.Create(m)
	if err != nil {
		return nil, err
	}

	m, err = s.Repo.Messages.GetByID(m.Id)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Messages) Update(data *models.Message) (*models.Message, error) {
	m, err := s.Repo.Messages.GetByID(data.Id)
	if err != nil {
		return nil, err
	}

	if m.DeletedAt != nil {
		return nil, app_errors.AccessDenied
	}

	if !equal(data.SenderId, m.SenderId) || m.ForwardFromId != nil {
		return nil, app_errors.AccessDenied
	}

	m.Text = data.Text
	m.Attachments = data.Attachments
	m.UpdatedAt = time.Now()

	err = s.Repo.Messages.Update(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Messages) SoftDelete(id, userId int) (*models.Message, error) {
	m, err := s.Repo.Messages.GetByID(id)
	if err != nil {
		return nil, err
	}

	if !equal(m.SenderId, &userId) || m.DeletedAt != nil {
		return nil, app_errors.AccessDenied
	}

	now := time.Now()
	m.DeletedAt = &now

	err = s.Repo.Messages.Update(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Messages) CreateMsgSeen(m *models.MessageSeen) (*models.MessageSeen, error) {
	msg, err := s.Repo.Messages.GetByID(m.MessageId)
	if err != nil {
		return nil, err
	}

	chat, err := s.Repo.Chats.GetByID(msg.ChatId)
	if err != nil {
		return nil, err
	}

	if chat.Type == types.ChatPersonal {
		ids, err := s.Repo.Chats.FindMemberIds(msg.ChatId)
		if err != nil {
			return nil, err
		}

		if !slices.Contains(ids, m.UserId) {
			return nil, app_errors.AccessDenied
		}
	}

	m, err = s.Repo.Messages.CreateMsgSeen(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Messages) GetByID(id int) (*models.Message, error) {
	return s.Repo.Messages.GetByID(id)
}
