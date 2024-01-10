package interactors

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/password"
	"github.com/takeuchi-shogo/ticket-api/pkg/random"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type authInteractor struct {
	db            usecase.DBUsecase
	jwt           token.JwtMakerInterface
	registerEmail usecase.RegisterEmailUsecase
	user          usecase.UserUsecase
}

func NewAuthInteractor(
	db usecase.DBUsecase,
	jwt token.JwtMakerInterface,
	registerEmail usecase.RegisterEmailUsecase,
	user usecase.UserUsecase,
) services.AuthService {
	return &authInteractor{
		db:            db,
		jwt:           jwt,
		registerEmail: registerEmail,
		user:          user,
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

	db, err := a.db.Connect()
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

func (a *authInteractor) Login(user *models.Users) (*models.MeInteractorResponse, *usecase.ResultStatus) {
	db, _ := a.db.Connect()

	foundUser, err := a.user.FindByEmail(db, user.Email)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusBadRequest, errors.New("メールアドレス、またはパスワードが違います"))
	}

	if err := password.CheckPassword(user.Password, foundUser.Password); err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: "",
		}, usecase.NewResultStatus(http.StatusUnauthorized, errors.New("メールアドレス、またはパスワードが違います"))
	}

	userID := strconv.FormatUint(foundUser.ID, 10)
	fmt.Println(userID)

	token, err := a.jwt.GenerateJWT(userID)
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

func (a *authInteractor) Create(user *models.Users) (*models.Users, string, *usecase.ResultStatus) {

	db, _ := a.db.Transaction()

	newUser, err := a.user.Create(db, user)
	if err != nil {
		db.Rollback()
		return &models.Users{}, "", usecase.NewResultStatus(http.StatusUnauthorized, err)
	}

	userID := strconv.Itoa(int(newUser.ID))

	jwtToken, err := a.jwt.GenerateJWT(userID)
	if err != nil {
		db.Rollback()
		return &models.Users{}, "", usecase.NewResultStatus(http.StatusUnauthorized, err)
	}

	db.Commit()
	return newUser, jwtToken, usecase.NewResultStatus(http.StatusOK, nil)
}
