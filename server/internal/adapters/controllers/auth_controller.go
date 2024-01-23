package controllers

import (
	"fmt"
	"net/http"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type AuthController interface {
	RegisterEmail(ctx Context)
	GetRegisterEmail(ctx Context)
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
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: nil})
}

func (a *authController) GetRegisterEmail(ctx Context) {

	token := ctx.Query("token")

	registerEmail, res := a.authService.GetRegisterEmail(token)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}
	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: registerEmail})
}

func (a *authController) VerifyCode(ctx Context) {

	token := ctx.Query("token")
	pinCode := ctx.Query("pin_code")
	fmt.Println(token, pinCode)

	res := a.authService.VerifyCode(&models.RegisterEmails{
		Token:   token,
		PinCode: pinCode,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
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

	foundUser, res := a.authService.Login(user)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("_refresh_token", foundUser.Token.RefreshToken, 24*30*60*60, "/", "localhost", true, true)

	ctx.JSON(res.StatusCode, presenters.NewResponse(foundUser))
}

func (a *authController) Signup(ctx Context) {

	firstName := ctx.PostForm("first_name")
	lastName := ctx.PostForm("last_name")
	displayName := ctx.PostForm("display_name")
	email := ctx.PostForm("email")
	tel := ctx.PostForm("tel")
	password := ctx.PostForm("password")
	postCode := ctx.PostForm("post_code")
	prefecture := ctx.PostForm("prefecture")
	city := ctx.PostForm("city")

	token := ctx.PostForm("token")

	u := &models.Users{
		FirstName:   firstName,
		LastName:    lastName,
		DisplayName: &displayName,
		Email:       email,
		Tel:         tel,
		Password:    password,
		PostCode:    postCode,
		Prefecture:  prefecture,
		City:        city,
	}

	user, res := a.authService.Create(u, token)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.ErrResponse{ErrorMessage: res.Err.Error()})
		return
	}

	// ctx.Header(constants.AuthorizationHeaderKey, user.Token.AccessToken)
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie("_refresh_token", user.Token.RefreshToken, 24*30*60*60, "/", "localhost", true, true)

	ctx.JSON(res.StatusCode, presenters.Response{Message: "success", Data: user})
}

func (a *authController) Logout(ctx Context) {

}
