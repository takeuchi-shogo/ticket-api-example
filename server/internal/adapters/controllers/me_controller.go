package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	Patch(ctx Context)
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

type UserRequest struct {
	DisplayName *string `json:"display_name"`
	Email       *string `json:"email"`
}

func (m *meController) Patch(ctx Context) {

	authPayload := ctx.MustGet(constants.AuthorizationPayloadKey).(*token.CustomClaims)

	id, _ := strconv.ParseUint(authPayload.UserID, 10, 64)

	user := &models.Users{
		ID: id,
	}

	body, _ := ctx.GetRawData()
	u := &UserRequest{}

	if err := json.Unmarshal(body, u); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	email := *u.Email
	fmt.Println("user reqest", email)
	if displayName, ok := ctx.GetPostForm("displayName"); ok {
		user.DisplayName = displayName
	}
	if email, ok := ctx.GetPostForm("email"); ok {
		fmt.Println("email", email)
		user.Email = email
	}

	me, res := m.me.Save(user)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.JSON(res.StatusCode, presenters.NewResponse(me))
}
