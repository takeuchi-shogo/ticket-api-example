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

	db := m.db.Connect()

	foundUser, err := m.user.FindByEmail(db, user.Email)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	if err := password.CheckPassword(user.Password, foundUser.Password); err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusUnauthorized, err)
	}

	token, err := m.generateJWT(int(user.ID))
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return &models.MeInteractorResponse{
		User:  foundUser.BuildForGet(),
		Token: token,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (m *meInteractor) Create(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus) {

	db := m.db.Connect()

	generatePassword, err := password.GenerateFromPassword(user.Password)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	user.Password = generatePassword

	newUser, err := m.user.Create(db, user)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	token, err := m.generateJWT(int(newUser.ID))
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return &models.MeInteractorResponse{
		User:  newUser.BuildForGet(),
		Token: token,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (m *meInteractor) generateJWT(userID int) (string, error) {
	id := int(userID)

	token, err := m.jwt.GenerateJWT(strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return token, nil
}
