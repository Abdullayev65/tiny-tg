package service

import (
	"slices"
	"time"
	"tiny-tg/internal/models"
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
