package service

//
//import (
//	"errors"
//	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
//	"time"
//
//	"github.com/alihaqberdi/goga_go/internal/dtos"
//	"github.com/alihaqberdi/goga_go/internal/models/types"
//	"github.com/alihaqberdi/goga_go/internal/pkg/mapping"
//	"github.com/alihaqberdi/goga_go/internal/repo"
//	"github.com/alihaqberdi/goga_go/internal/service/caching"
//)
//
//type Tenders struct {
//	Repo  *repo.Repo
//	Cache *caching.Cache
//}
//
//func (s *Tenders) Create(tender *dtos.Tender) (*dtos.Tender, error) {
//	tender.Status = types.TenderStatusOpen
//
//	if err := s.validateTender(tender); err != nil {
//		return nil, err
//	}
//
//	tenderModel := mapping.ConvertTenderDTOToModel(tender)
//
//	createdTenderModel, err := s.Repo.Tenders.Create(tenderModel)
//	if err != nil {
//		return nil, err
//	}
//
//	tenderDTO := &dtos.Tender{
//		ID:          createdTenderModel.ID,
//		ClientId:    createdTenderModel.ClientId,
//		Title:       createdTenderModel.Title,
//		Description: createdTenderModel.Description,
//		Deadline:    createdTenderModel.Deadline,
//		Budget:      createdTenderModel.Budget,
//		Status:      createdTenderModel.Status,
//	}
//
//	return tenderDTO, nil
//}
//
//func (s *Tenders) update(data *dtos.Tender) (*dtos.Tender, error) {
//	model, err := s.Repo.Tenders.GetByID(data.ID)
//	if err != nil {
//		return nil, app_errors.TenderNotFound
//	}
//
//	if model.ClientId != data.ClientId {
//		return nil, err
//	}
//
//	{
//		if data.Title == "" {
//			data.Title = model.Title
//		}
//		if data.Description == "" {
//			data.Description = model.Description
//		}
//		if data.Deadline.IsZero() {
//			data.Deadline = model.Deadline
//		}
//		if data.Budget == 0 {
//			data.Budget = model.Budget
//		}
//		if data.Status == "" {
//			data.Status = model.Status
//		}
//
//	}
//
//	err = s.validateTender(data)
//	if err != nil {
//		return nil, err
//	}
//
//	model = mapping.ConvertTenderDTOToModel(data)
//	err = s.Repo.Tenders.update(model)
//
//	return nil, err
//}
//
//func (s *Tenders) Delete(id, clientID uint) error {
//	tender, err := s.Repo.Tenders.GetByID(id)
//	if err != nil || tender.ClientId != clientID {
//		return app_errors.TenderNotFoundOrAccessDenied
//	}
//
//	err = s.Repo.Tenders.Delete(id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *Tenders) GetList(data *dtos.Tenders) ([]dtos.Tender, error) {
//	if data.Limit == 0 {
//		data.Limit = 10
//	}
//
//	tenders, err := s.Repo.Tenders.GetList(data)
//	if err != nil {
//		return nil, err
//	}
//
//	tenderDTOs := make([]dtos.Tender, len(tenders))
//	for i, model := range tenders {
//		tenderDTOs[i] = dtos.Tender{
//			ID:          model.ID,
//			ClientId:    model.ClientId,
//			Title:       model.Title,
//			Description: model.Description,
//			Deadline:    model.Deadline,
//			Budget:      model.Budget,
//			Status:      model.Status,
//		}
//	}
//
//	return tenderDTOs, nil
//}
//
//func (s *Tenders) validateTender(tender *dtos.Tender) error {
//	if tender.Budget <= 0 {
//		return app_errors.TenderInvalidData
//	}
//
//	if tender.Deadline.Before(time.Now()) {
//		return errors.New("deadline must be in the future")
//	}
//
//	// Ensure that the status is valid (you can expand this based on your business rules)
//	if !tender.Status.Valid() {
//		return app_errors.TenderInvalidStatus
//	}
//
//	return nil
//}
//
//func (s *Tenders) GetListTendersByUser(userID, limit, offset int) ([]dtos.Tender, error) {
//	tenders, err := s.Repo.Tenders.GetListByUser(userID, limit, offset)
//	if err != nil {
//		return nil, err
//	}
//
//	tenderDTOs := make([]dtos.Tender, len(tenders))
//	for i, model := range tenders {
//		tenderDTOs[i] = dtos.Tender{
//			ID:          model.ID,
//			ClientId:    model.ClientId,
//			Title:       model.Title,
//			Description: model.Description,
//			Deadline:    model.Deadline,
//			Budget:      model.Budget,
//			Status:      model.Status,
//		}
//	}
//
//	return tenderDTOs, nil
//}
