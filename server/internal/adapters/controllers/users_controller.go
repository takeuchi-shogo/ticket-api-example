package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/gateways/database"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/interactors"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type UserController interface {
	Get(ctx Context)
}

type usersController struct {
	userService services.UserService
}

func NewUsersController() UserController {
	return &usersController{
		userService: &interactors.UserInteractor{
			UserRepository: database.NewUserRepository(),
		},
	}
}

func (u *usersController) Get(ctx Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res, _ := u.userService.Get(id)

	ctx.JSON(200, res)
}
