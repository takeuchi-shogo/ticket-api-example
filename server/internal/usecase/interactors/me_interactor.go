package interactors

import (
	"net/http"
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/password"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type meInteractor struct {
	db   usecase.DBUsecase
	jwt  token.JwtMakerInterface
	user usecase.UserUsecase
}

func NewMeInteractor(
	db usecase.DBUsecase,
	jwt token.JwtMakerInterface,
	user usecase.UserUsecase,
) services.MeService {
	return &meInteractor{
		db:   db,
		jwt:  jwt,
		user: user,
	}
}

type MeInteractorResponse struct {
	User  *models.UsersResponse
	Token string
}

func (m *meInteractor) Get(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus) {

	db, _ := m.db.Connect()

	foundUser, err := m.user.FindByEmail(db, user.Email)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	if err := password.CheckPassword(user.Password, foundUser.Password); err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusUnauthorized, err)
	}

	tokenPairs, err := m.generateJWT(int(user.ID))
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return &models.MeInteractorResponse{
		User:  foundUser.BuildForGet(),
		Token: tokenPairs,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (m *meInteractor) Create(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus) {

	db, _ := m.db.Connect()

	generatePassword, err := password.GenerateFromPassword(user.Password)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	user.Password = generatePassword

	newUser, err := m.user.Create(db, user)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	tokenPairs, err := m.generateJWT(int(newUser.ID))
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return &models.MeInteractorResponse{
		User:  newUser.BuildForGet(),
		Token: tokenPairs,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (m *meInteractor) GetMe(userID int) (*models.UsersResponse, *usecase.ResultStatus) {

	db, _ := m.db.Connect()

	foundUser, err := m.user.FindByID(db, userID)
	if err != nil {
		return &models.UsersResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return foundUser.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (m *meInteractor) generateJWT(userID int) (*token.TokenPairs, error) {
	id := int(userID)

	tokenPairs, err := m.jwt.GenerateJWT(strconv.Itoa(id))
	if err != nil {
		return &token.TokenPairs{}, err
	}
	return tokenPairs, nil
}

func (m *meInteractor) Save(user *models.Users) (*models.UsersResponse, *usecase.ResultStatus) {

	db, _ := m.db.Connect()

	foundUser, err := m.user.FindByID(db, int(user.ID))
	if err != nil {
		return &models.UsersResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	foundUser.DisplayName = user.DisplayName

	updatedUser, err := m.user.Save(db, foundUser)
	if err != nil {
		return &models.UsersResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return updatedUser.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}
