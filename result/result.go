package result

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func ErrorResponse(code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
