package response

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(message string, data interface{}) Response {
	return Response{
		Status:  200,
		Message: message,
		Data:    data,
	}
}

func Error(status int, message string) Response {
	return Response{
		Status:  status,
		Message: message,
	}
}
