package tasks

type UserMailLogsController interface{}

type userMailLogsController struct {
}

func NewUserMailLogsController() UserMailLogsController {
	return &userMailLogsController{}
}
