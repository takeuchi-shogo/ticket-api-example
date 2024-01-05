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

type MeController interface {
	Get(ctx Context)
	Post(ctx Context)
	GetMe(ctx Context)
}

type meController struct {
	me services.MeService
}

func NewMeController(me services.MeService) MeController {
	return &meController{
		me: me,
	}
}

func (m *meController) Get(ctx Context) {

	user := &models.Users{}

	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, presenters.NewErrResponse(err.Error()))
		return
	}

	foundUser, res := m.me.Get(user)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.Header(constants.AuthorizationHeaderKey, foundUser.Token)

	ctx.JSON(res.StatusCode, presenters.NewResponse(foundUser.User))
}

func (m *meController) Post(ctx Context) {

	user := &models.Users{}

	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, presenters.NewErrResponse(err.Error()))
		return
	}

	newUser, res := m.me.Create(user)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.Header(constants.AuthorizationHeaderKey, newUser.Token)

	ctx.JSON(res.StatusCode, presenters.NewResponse(newUser.User))
}

func (m *meController) GetMe(ctx Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	userID, _ := strconv.Atoi(authPayload.UserID)

	foundUser, res := m.me.GetMe(userID)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.JSON(res.StatusCode, presenters.NewResponse(foundUser))
}
