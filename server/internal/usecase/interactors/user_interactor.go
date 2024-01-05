package interactors

import (
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
)

type UserInteractor struct {
	UserRepository usecase.UserUsecase
}

func NewUserInteractor(
	userRepository usecase.UserUsecase,
) *UserInteractor {
	return &UserInteractor{
		UserRepository: userRepository,
	}
}

func (u *UserInteractor) Get(id int) (*models.Users, error) {
	// user, err := u.UserRepository.FindByID(id)
	// if err != nil {
	// 	return &models.Users{}, err
	// }
	return &models.Users{}, nil
}
