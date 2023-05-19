package repository

import (
	"context"
	"gorm.io/gorm"
	"self-payroll-rakamin/app/domain/models"
)

type positionRepository struct {
	DB *gorm.DB
}

func NewPositionRepository(db *gorm.DB) models.PositionRepository {
	return &positionRepository{DB: db}
}

func (p *positionRepository) Create(ctx context.Context, position *models.Position) (*models.Position, error) {
	if err := p.DB.WithContext(ctx).Create(&position).Error; err != nil {
		return nil, err
	}
	return position, nil
}
