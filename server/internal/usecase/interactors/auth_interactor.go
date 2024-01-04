package interactors

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/random"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type authInteractor struct {
	DB            usecase.DBUsecase
	jwt           token.JwtMakerInterface
	registerEmail usecase.RegisterEmailUsecase
}

func NewAuthInteractor(
	db usecase.DBUsecase,
	jwt token.JwtMakerInterface,
	registerEmail usecase.RegisterEmailUsecase,
) services.AuthService {
	return &authInteractor{
		DB:            db,
		jwt:           jwt,
		registerEmail: registerEmail,
	}
}

func (a *authInteractor) RegisterEmail(email string) *usecase.ResultStatus {

	// db, err := a.DB.Connect()
	// if err != nil {
	// 	return usecase.NewResultStatus(http.StatusInternalServerError, err)
	// }

	registerEmail := &models.RegisterEmails{
		Email:   email,
		Token:   random.RandomString(200),
		PinCode: random.RandomPinCode(),
	}

	// fmt.Println(registerEmail)
	fmt.Printf("%+v\n", registerEmail)

	// newRegisterEmail, err := a.registerEmail.Create(db, registerEmail)
	// if err != nil {
	// 	return usecase.NewResultStatus(http.StatusBadRequest, err)
	// }

	// fmt.Printf("%+v\n", newRegisterEmail)

	return usecase.NewResultStatus(http.StatusOK, nil)
}

func (a *authInteractor) VerifyCode(registerEmail *models.RegisterEmails) *usecase.ResultStatus {

	db, err := a.DB.Connect()
	if err != nil {
		return usecase.NewResultStatus(http.StatusInternalServerError, err)
	}

	foundRegisterEmail, err := a.registerEmail.FindByPinCode(db, registerEmail.PinCode)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	if foundRegisterEmail.Token != registerEmail.Token {
		return usecase.NewResultStatus(http.StatusBadRequest, errors.New("無効なトークンです"))
	}
	if foundRegisterEmail.IsValid {
		return usecase.NewResultStatus(http.StatusBadRequest, errors.New("無効な操作です"))
	}
	if foundRegisterEmail.ExpireAt <= time.Now().Unix() {
		return usecase.NewResultStatus(http.StatusBadRequest, errors.New("有効期限が切れています"))
	}

	return usecase.NewResultStatus(http.StatusOK, nil)
}

func (a *authInteractor) Verify(userID int) (*models.Users, *usecase.ResultStatus) {
	fmt.Println(userID)

	// _ = a.DB.Connect()

	return &models.Users{}, usecase.NewResultStatus(200, nil)
}

func (a *authInteractor) Login(user *models.Users) (*models.Users, *usecase.ResultStatus) {
	fmt.Println(user)

	// _ = a.DB.Connect()

	return &models.Users{}, usecase.NewResultStatus(200, nil)
}

func (a *authInteractor) Create(*models.Users) (*models.Users, *usecase.ResultStatus) {

	jwtToken, err := a.jwt.GenerateJWT("1")
	fmt.Println("interactor")
	if err != nil {
		return &models.Users{}, usecase.NewResultStatus(http.StatusUnauthorized, err)
	}

	fmt.Println(jwtToken)
	return &models.Users{}, usecase.NewResultStatus(http.StatusOK, nil)
}
