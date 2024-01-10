package controllers

import (
	"strconv"

	"github.com/takeuchi-shogo/ticket-api/internal/adapters/presenters"
	"github.com/takeuchi-shogo/ticket-api/internal/domain/models"
	"github.com/takeuchi-shogo/ticket-api/internal/usecase/services"
)

type ArtistsController interface {
	Get(ctx Context)
	GetList(ctx Context)
	Post(ctx Context)
}

type artistsController struct {
	artist services.ArtistService
}

func NewArtistsController(
	artist services.ArtistService,
) ArtistsController {
	return &artistsController{
		artist: artist,
	}
}

func (a *artistsController) Get(ctx Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	artist, res := a.artist.Get(id)
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.JSON(res.StatusCode, presenters.NewResponse(artist))
}

func (a *artistsController) GetList(ctx Context) {

	artists, res := a.artist.GetList()
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.JSON(res.StatusCode, presenters.NewResponse(artists))
}

func (a *artistsController) Post(ctx Context) {

	displayName := ctx.PostForm("display_name")
	screenName := ctx.PostForm("screen_name")
	description := ctx.PostForm("description")

	artist, res := a.artist.Create(&models.Artists{
		DisplayName: displayName,
		ScreenName:  screenName,
		Description: &description,
	})
	if res.Err != nil {
		ctx.JSON(res.StatusCode, presenters.NewErrResponse(res.Err.Error()))
		return
	}

	ctx.JSON(res.StatusCode, presenters.NewResponse(artist))
}
