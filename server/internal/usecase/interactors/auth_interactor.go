package interactors

import (
	"fmt"
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type authInteractor struct {
	// DB  usecase.DBUsecase
	jwt token.JwtMakerInterface
}

func NewAuthInteractor(jwt token.JwtMakerInterface) services.AuthService {
	return &authInteractor{
		// DB:  db,
		jwt: jwt,
	}
}

func (a *authInteractor) Verify(userID int) (*models.Users, *usecase.ResultStatus) {
	fmt.Println(userID)

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
