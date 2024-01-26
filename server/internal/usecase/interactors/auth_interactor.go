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
	userMailLog   usecase.UserMailLogUsecase
}

func NewAuthInteractor(
	db usecase.DBUsecase,
	jwt token.JwtMakerInterface,
	registerEmail usecase.RegisterEmailUsecase,
	user usecase.UserUsecase,
	userMailLog usecase.UserMailLogUsecase,
) services.AuthService {
	return &authInteractor{
		db:            db,
		jwt:           jwt,
		registerEmail: registerEmail,
		user:          user,
		userMailLog:   userMailLog,
	}
}

func (a *authInteractor) RegisterEmail(email string) *usecase.ResultStatus {

	db, err := a.db.Connect()
	if err != nil {
		return usecase.NewResultStatus(http.StatusInternalServerError, err)
	}
	if email == "" {
		return usecase.NewResultStatus(http.StatusBadRequest, errors.New("メールアドレスが入力されていません"))
	}

	registerEmail := &models.RegisterEmails{
		Email:   email,
		Token:   random.RandomString(30),
		PinCode: random.RandomPinCode(),
	}

	for {
		_, err := a.registerEmail.FindByToken(db, registerEmail.Token)
		if err != nil {
			break
		}
		registerEmail.Token = random.RandomString(30)
	}

	for {
		_, err := a.registerEmail.FindByPinCode(db, registerEmail.PinCode)
		if err != nil {
			break
		}
		registerEmail.Token = random.RandomString(30)
	}

	_, err = a.registerEmail.Create(db, registerEmail)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return usecase.NewResultStatus(http.StatusNoContent, nil)
}

func (a *authInteractor) GetRegisterEmail(token string) (*models.RegisterEmailsResponse, *usecase.ResultStatus) {
	db, _ := a.db.Connect()

	registerEmail, err := a.registerEmail.FindByToken(db, token)
	if err != nil {
		return &models.RegisterEmailsResponse{}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	return registerEmail.BuildForGet(), usecase.NewResultStatus(http.StatusOK, nil)
}

func (a *authInteractor) VerifyCode(registerEmail *models.RegisterEmails) *usecase.ResultStatus {

	db, err := a.db.Connect()
	if err != nil {
		return usecase.NewResultStatus(http.StatusInternalServerError, err)
	}

	foundRegisterEmail, err := a.registerEmail.FindByPinCode(db, registerEmail.PinCode)
	if err != nil {
		return usecase.NewResultStatus(http.StatusBadRequest, errors.New("入力されたコードに誤りがあります"))
	}

	if foundRegisterEmail.Token != registerEmail.Token {
		return usecase.NewResultStatus(http.StatusBadRequest, errors.New("無効なトークンです"))
	}
	if !foundRegisterEmail.IsValid {
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
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, errors.New("メールアドレス、またはパスワードが違います"))
	}

	if err := password.CheckPassword(user.Password, foundUser.Password); err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusUnauthorized, errors.New("メールアドレス、またはパスワードが違います"))
	}

	userID := strconv.FormatUint(foundUser.ID, 10)

	tokenPairs, err := a.jwt.GenerateJWT(userID)
	if err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusUnauthorized, err)
	}

	return &models.MeInteractorResponse{
		User:  foundUser.BuildForGet(),
		Token: tokenPairs,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}

func (a *authInteractor) Create(user *models.Users, verifyToken string) (*models.MeInteractorResponse, *usecase.ResultStatus) {

	db, _ := a.db.Transaction()

	if _, err := a.registerEmail.FindByToken(db, verifyToken); err != nil {
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	newUser, err := a.user.Create(db, user)
	if err != nil {
		db.Rollback()
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	userID := strconv.Itoa(int(newUser.ID))

	jwtToken, err := a.jwt.GenerateJWT(userID)
	if err != nil {
		db.Rollback()
		return &models.MeInteractorResponse{
			User:  &models.UsersResponse{},
			Token: &token.TokenPairs{},
		}, usecase.NewResultStatus(http.StatusBadRequest, err)
	}

	db.Commit()
	return &models.MeInteractorResponse{
		User:  newUser.BuildForGet(),
		Token: jwtToken,
	}, usecase.NewResultStatus(http.StatusOK, nil)
}
