package util

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	FormatterJSON struct {
		Success bool        `json:"success"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func ResponseSuccess(e echo.Context, message string, data interface{}) error {
	response := FormatterJSON{
		Success: true,
		Message: message,
		Data:    data,
	}
	return e.JSON(http.StatusOK, response)
}

func ResponseError(e echo.Context, message string, data interface{}) error {
	response := FormatterJSON{
		Success: false,
		Message: message,
		Data:    data,
	}
	return e.JSON(http.StatusBadRequest, response)
}
