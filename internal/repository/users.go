package repository

import (
	"gorm.io/gorm"
	"tiny-tg/internal/dtos"
	"tiny-tg/internal/models"
)

type Users struct {
	DB *gorm.DB
}

func (r *Users) Create(bid *models.User) (*models.User, error) {
	if err := r.DB.Create(bid).Error; err != nil {
		return nil, err
	}

	return bid, nil
}

//func (r *Users) GetList(data *dtos.Users) ([]models.User, error) {
//	var bids []models.User
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
//	err := query.Limit(data.Limit).Offset(data.Offset).Find(&bids).Error
//
//	if err != nil {
//		return nil, err
//	}
//
//	return bids, nil
//}

func (r *Users) GetByID(id int) (*models.User, error) {
	var bid models.User
	if err := r.DB.First(&bid, id).Error; err != nil {
		return nil, err
	}

	return &bid, nil
}

func (r *Users) Delete(id int) error {
	if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *Users) GetByUsername(username string) (*models.User, error) {
	var m models.User
	err := r.DB.Where("username = ?", username).First(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *Users) Search(data *dtos.ListOpts) ([]models.User, error) {
	var ms []models.User
	query := r.DB

	err := query.Limit(data.Limit).Offset(data.Offset).Find(&ms).Error

	if err != nil {
		return nil, err
	}

	return ms, nil
}
