package tasks

import "github.com/takeuchi-shogo/ticket-api/internal/usecase/services/tasks"

type userMailLogInteractor struct{}

func NewUserMailLogInteractor() tasks.UserMailLogService {
	return &userMailLogInteractor{}
}
