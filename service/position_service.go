package service

import (
	"context"
	"self-payroll-rakamin/app/domain/models"
	"self-payroll-rakamin/app/domain/web"
)

type positionService struct {
	positionRepository models.PositionRepository
}

func NewPositionService(position models.PositionRepository) models.PositionService {
	return &positionService{positionRepository: position}
}

func (p positionService) StorePosition(ctx context.Context, request *web.PositionRequest) (*models.Position, error) {
	newPosition := &models.Position{
		Name:   request.Name,
		Salary: request.Salary,
	}

	position, err := p.positionRepository.Create(ctx, newPosition)
	if err != nil {
		return nil, err
	}
	return position, nil
}
