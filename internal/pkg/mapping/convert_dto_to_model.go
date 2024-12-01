package mapping

import (
	"github.com/alihaqberdi/goga_go/internal/dtos"
	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/gorm"
)

func ConvertTenderDTOToModel(dto *dtos.Tender) *models.Tender {
	return &models.Tender{
		Model: gorm.Model{
			ID: dto.ID,
		},
		ClientId:    dto.ClientId,
		Title:       dto.Title,
		Description: dto.Description,
		Deadline:    dto.Deadline,
		Budget:      dto.Budget,
		Status:      dto.Status,
	}
}
