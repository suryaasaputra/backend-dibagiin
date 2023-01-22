package helpers

type Response struct {
	IsError bool        `json:"error" example:"false"`
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data,omitempty"`
}

func GetResponse(isError bool, code int, message string, payload interface{}) Response {
	return Response{
		IsError: isError,
		Code:    code,
		Message: message,
		Data:    payload,
	}
}
