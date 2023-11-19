package controllers

import (
	"net/http"
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
	"github.com/takeuchi-shogo/ticket-api/pkg/constants"
	"github.com/takeuchi-shogo/ticket-api/pkg/token"
)

type AuthController interface {
	Signin(ctx Context)
	Signup(ctx Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(s services.AuthService) AuthController {
	return &authController{
		authService: s,
	}
}

func (a *authController) Signin(ctx Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	userID, _ := strconv.Atoi(authPayload.ID)

	user, res := a.authService.Verify(userID)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: user})
}

func (a *authController) Signup(ctx Context) {

	u := &models.Users{}

	if err := ctx.BindJSON(u); err != nil {
		ctx.JSON(http.StatusBadRequest, presenters.ErrResponse{ErrorMessage: err.Error()})
		return
	}

	user, res := a.authService.Create(u)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: user})
}
