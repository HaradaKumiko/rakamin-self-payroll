package response

type (
	FormatterJSON struct {
		Success bool        `json:"success"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data"`
	}
)
