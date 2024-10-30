package controllers

import (
	"fmt"
	"net/http"
	"smol/core"
	"smol/dtos"
	"smol/services"

	"github.com/labstack/echo/v4"
)

type urlController struct {
	app     *core.App
	service services.IUrlService
}

func bindUrlController(app *core.App, group *echo.Group) {
	urlService := services.InitUrlService(app)
	controller := urlController{app, urlService}

	subGroup := group.Group("/u")

	subGroup.GET("/:id", controller.redirect)
	subGroup.POST("", controller.shorten)
}

func (cont *urlController) shorten(c echo.Context) error {
	var dto dtos.ShortenUrlDto
	if err := c.Bind(&dto); err != nil {
		return echo.ErrBadRequest
	}

	shortened, err := cont.service.ShortenUrl(dto.Url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(http.StatusOK, dtos.ShortenUrlResponse{
		UrlPath: shortened,
	})
}

func (cont *urlController) redirect(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return echo.ErrBadRequest
	}

	longUrl, err := cont.service.GetLongUrl(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.Redirect(http.StatusTemporaryRedirect, longUrl)
}
