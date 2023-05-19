package web

import validation "github.com/go-ozzo/ozzo-validation"

type (
	PositionRequest struct {
		Name   string `json:"name" validate:"required"`
		Salary int64  `json:"salary" validate:"required,number"`
	}
)

func (req PositionRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Salary, validation.Required),
	)
}
