package repository

import (
	"gorm.io/gorm"
	"tiny-tg/internal/models"
	"tiny-tg/internal/models/types"
)

type Chats struct {
	DB *gorm.DB
}

func (r *Chats) Create(m *models.Chat) (*models.Chat, error) {
	err := r.DB.Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

//func (r *Chats) GetList(data *dtos.Chats) ([]models.Chat, error) {
//	var ms []models.Chat
//	query := r.DB
//
//	if data.ContractorID > 0 && data.TenderID > 0 {
//		query = query.Where("contractor_id = ? AND tender_id=?", data.ContractorID, data.TenderID)
//	} else if data.ContractorID > 0 {
//		query = query.Where("contractor_id=?", data.ContractorID)
//	} else if data.TenderID > 0 {
//		query = query.Where("tender_id=?", data.TenderID)
//	}
//
//	err := query.Limit(data.Limit).Offset(data.Offset).Find(&ms).Error
//
//	if err != nil {
//		return nil, err
//	}
//
//	return ms, nil
//}

func (r *Chats) GetByID(id int) (*models.Chat, error) {
	var m models.Chat

	err := r.DB.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *Chats) Delete(id int) error {
	if err := r.DB.Delete(&models.Chat{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *Chats) GetPersonalChat(userIds [2]int) (*models.Chat, error) {
	m := new(models.Chat)
	err := r.DB.Raw(`SELECT ch.*
FROM chats AS ch
JOIN chat_members AS m ON ch.id = m.chat_id
WHERE ch."type" = ? AND m.id IN ?`, types.ChatPersonal, userIds).Scan(m).Error
	if err != nil {
		return nil, err
	}

	if m.Id == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return m, nil
}

func (r *Chats) CreateMembers(id int, memberIds []int) error {

	var members []ChatMember
	for _, memberId := range memberIds {
		members = append(members, ChatMember{ChatId: id, UserId: memberId})
	}

	err := r.DB.Create(members).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Chats) FindMemberIds(id int) ([]int, error) {
	var list []ChatMember
	err := r.DB.Where("chat_id = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}

	ids := make([]int, len(list))
	for i, m := range list {
		ids[i] = m.UserId
	}

	return ids, nil
}
