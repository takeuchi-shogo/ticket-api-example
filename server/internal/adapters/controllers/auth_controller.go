package controllers

import (
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type AuthController interface {
	RegisterEmail(ctx Context)
	VerifyCode(ctx Context)
	Signup(ctx Context)
	Signin(ctx Context)
	Logout(ctx Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(s services.AuthService) AuthController {
	return &authController{
		authService: s,
	}
}

func (a *authController) RegisterEmail(ctx Context) {
	email := ctx.PostForm("email")

	res := a.authService.RegisterEmail(email)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: nil})
}

func (a *authController) VerifyCode(ctx Context) {

	token := ctx.PostForm("token")
	pinCode := ctx.PostForm("pin_code")

	res := a.authService.VerifyCode(&models.RegisterEmails{
		Token:   token,
		PinCode: pinCode,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: nil})
}

func (a *authController) Signin(ctx Context) {

	email := ctx.PostForm("email")
	pass := ctx.PostForm("password")

	user := &models.Users{
		Email:    email,
		Password: pass,
	}

	// if err := ctx.BindJSON(user); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, presenters.NewErrResponse(err.Error()))
	// 	return
	// }

	_, res := a.authService.Login(user)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	// ctx.Header(constants.AuthorizationHeaderKey, foundUser.Token)

	// ctx.JSON(res.StatusCode, presenters.NewResponse(foundUser.User))
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

func (a *authController) Logout(ctx Context) {

}
