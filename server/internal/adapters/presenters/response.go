package presenters

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrResponse struct {
	ErrorMessage string `json:"error_message"`
}
