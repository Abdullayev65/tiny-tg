package repository

import (
	"gorm.io/gorm"
	"tiny-tg/internal/models"
)

type Messages struct {
	DB *gorm.DB
}

func (r *Messages) Create(m *models.Message) (*models.Message, error) {
	err := r.DB.Create(m).Error
	if err != nil {
		return nil, err
	}

	return m, nil
}

//func (r *Messages) GetList(data *dtos.Messages) ([]models.Message, error) {
//	var ms []models.Message
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

func (r *Messages) GetByID(id int) (*models.Message, error) {
	var m models.Message

	err := r.DB.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *Messages) Delete(id int) error {

	err := r.DB.Delete(&models.Message{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
