package delivery

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"self-payroll-rakamin/app/domain/models"
	"self-payroll-rakamin/app/domain/web"
	"self-payroll-rakamin/util/response"
)

type (
	positionDelivery struct {
		positionService models.PositionService
	}

	PositionDelivery interface {
		Mount(group *echo.Group)
	}
)

func NewPositionDelivery(positionService models.PositionService) positionDelivery {
	return positionDelivery{positionService: positionService}
}

func (p *positionDelivery) Mount(group *echo.Group) {
	group.POST("", p.StorePositionHandler)
}

func (p *positionDelivery) StorePositionHandler(e echo.Context) error {
	ctx := e.Request().Context()

	var req web.PositionRequest

	if err := e.Bind(&req); err != nil {
		response := response.FormatterJSON{
			Success: false,
			Message: "Something Wrong" + err.Error(),
			Data:    nil,
		}
		return e.JSON(http.StatusUnprocessableEntity, response)
	}

	position, err := p.positionService.StorePosition(ctx, &req)
	if err != nil {
		response := response.FormatterJSON{
			Success: false,
			Message: "Something Wrong" + err.Error(),
			Data:    nil,
		}
		return e.JSON(http.StatusInternalServerError, response)
	}

	response := response.FormatterJSON{
		Success: true,
		Message: "Success Create Position",
		Data:    position,
	}
	return e.JSON(http.StatusOK, response)
}
