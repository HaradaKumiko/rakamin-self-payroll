package models

import (
	"context"
	"self-payroll-rakamin/app/domain/web"
)

type (
	Position struct {
		ID     int64  `json:"id" gorm:"primaryKey"`
		Name   string `json:"name"`
		Salary int64  `json:"salary"`
	}

	PositionRepository interface {
		Create(ctx context.Context, position *Position) (*Position, error)
	}

	PositionService interface {
		StorePosition(ctx context.Context, request *web.PositionRequest) (*Position, error)
	}
)
