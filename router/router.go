package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"self-payroll-rakamin/delivery"
	"self-payroll-rakamin/repository"
	"self-payroll-rakamin/service"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	positionRepository := repository.NewPositionRepository(db)
	positionService := service.NewPositionService(positionRepository)
	positionDelivery := delivery.NewPositionDelivery(positionService)

	api := e.Group("/api")

	positionRoute := api.Group("/position")
	positionDelivery.Mount(positionRoute)
}
