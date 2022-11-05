package helpers

type Response struct {
	IsError bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
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
