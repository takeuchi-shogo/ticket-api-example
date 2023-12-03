package presenters

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrResponse struct {
	ErrorMessage string `json:"error_message"`
}

func NewResponse(data interface{}) Response {
	return Response{
		Message: "success",
		Data:    data,
	}
}

func NewErrResponse(message string) ErrResponse {
	return ErrResponse{
		ErrorMessage: message,
	}
}
