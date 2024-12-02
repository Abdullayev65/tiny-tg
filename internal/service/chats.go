package service

import (
	"time"
	"tiny-tg/internal/models"
	"tiny-tg/internal/models/types"
	"tiny-tg/internal/pkg/app_errors"
	"tiny-tg/internal/repository"
)

// import (
//
//	"errors"
//
//	"github.com/alihaqberdi/goga_go/internal/dtos"
//	"github.com/alihaqberdi/goga_go/internal/models"
//	"github.com/alihaqberdi/goga_go/internal/models/types"
//	"github.com/alihaqberdi/goga_go/internal/pkg/app_errors"
//	"github.com/alihaqberdi/goga_go/internal/repo"
//	"github.com/alihaqberdi/goga_go/internal/service/caching"
//
// )
type Chats struct {
	Repo *repository.Repo
}

func (s *Chats) GetGroupChat(id int) (*models.Chat, error) {
	model, err := s.Repo.Chats.GetByID(id)
	if err != nil {
		return nil, err
	}

	if model.Type != types.ChatGroup {
		return nil, app_errors.AccessDenied
	}

	model.MemberIds, err = s.Repo.Chats.FindMemberIds(model.Id)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Chats) MustGetPersonalChat(userIds [2]int) (*models.Chat, error) {
	model, err := s.Repo.Chats.GetPersonalChat(userIds)
	if err != nil {
		m, err := s.Repo.Chats.Create(&models.Chat{
			Type:      types.ChatPersonal,
			OwnerId:   userIds[0],
			CreatedAt: time.Now(),
		})
		if err != nil {
			return nil, err
		}

		err = s.Repo.Chats.CreateMembers(m.Id, userIds[:])
		if err != nil {
			return nil, err
		}

		return m, nil
	}

	model.MemberIds, err = s.Repo.Chats.FindMemberIds(model.Id)

	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *Chats) Create(m *models.Chat) (*models.Chat, error) {
	if m.Name == "" {
		m.Name = "Group " + time.Now().Format("Jan 02 15:04")
	}

	m, err := s.Repo.Chats.Create(m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Chats) CreateMembers(gropId int, memberIds []int) error {
	return s.Repo.Chats.CreateMembers(gropId, memberIds)
}

//
//func (s *Bids) Delete(id uint, contractorId uint) error {
//	bid, err := s.Repo.Bids.GetByID(id)
//	if err != nil || bid.ContractorId != contractorId {
//		return app_errors.BidNotFoundOrAccessDenied
//	}
//
//	return s.Repo.Bids.Delete(id)
//}
//
//func (s *Bids) UserBids(userID uint) ([]models.Bid, error) {
//	return s.Repo.Bids.UserBids(userID)
//}
//
//func (s *Bids) AwardBid(tenderID, id, clientID uint) error {
//	tender, err := s.Repo.Tenders.GetByID(tenderID)
//	_ = tender
//	if err != nil || tender.ClientId != clientID {
//		return app_errors.TenderNotFoundOrAccessDenied
//	}
//
//	bid, err := s.Repo.Bids.GetByID(id)
//	if err != nil {
//		return app_errors.BidNotFound
//	}
//
//	if bid.Status != types.BidStatusPending {
//		return app_errors.BidNotPending
//	}
//	//if tender.Status != types.TenderStatusClosed {
//	//	return app_errors.TenderNotClosed
//	//}
//
//	return s.Repo.Bids.AwardBid(id)
//}
//
//func (s *Bids) GetList(data *dtos.Bids) ([]dtos.BidList, error) {
//	if data.Limit == 0 {
//		data.Limit = 10
//	}
//
//	if data.TenderID > 0 {
//		_, err := s.Repo.Tenders.GetByID(data.TenderID)
//		if err != nil {
//			return nil, app_errors.TenderNotFoundOrAccessDenied
//		}
//	}
//
//	list, err := s.Repo.Bids.GetList(data)
//	if err != nil {
//		return nil, err
//	}
//
//	dtoList := make([]dtos.BidList, len(list))
//	for i, bid := range list {
//		dtoList[i] = *s.mapper(&bid)
//	}
//
//	return dtoList, nil
//}
//
//func (s *Bids) validateBid(bid *dtos.BidCreate) error {
//
//	if bid.Price <= 0 {
//		return app_errors.BidInvalidData
//	}
//
//	if bid.Status != types.BidStatusPending {
//		return errors.New("invalid status, must be 'pending'")
//	}
//	tender, err := s.Repo.Tenders.GetByID(bid.TenderID)
//
//	if err != nil {
//		return app_errors.TenderNotFound
//	}
//
//	if tender.Status != types.TenderStatusOpen {
//		return app_errors.BidTenderIsNotOpen
//	}
//	// You can add more validation rules as needed like rate limiting, etc.
//	return nil
//}
//
//func (s *Bids) mapper(m *models.Bid) *dtos.BidList {
//	return &dtos.BidList{
//		BidsBase: dtos.BidsBase{
//			TenderID:     m.TenderId,
//			ContractorID: m.ContractorId,
//			Price:        m.Price,
//			DeliveryTime: m.DeliveryTime,
//			Comments:     m.Comments,
//			Status:       m.Status,
//		},
//		ID: m.ID,
//	}
//}
