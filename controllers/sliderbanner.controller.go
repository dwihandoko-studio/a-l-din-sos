package controllers

import (
	"api-dinsos/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetSliderBanner(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	res, err := models.GetSliderBanner()
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	if res.Status != http.StatusOK {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: res.Message},
		)
	}

	return c.JSON(http.StatusOK,
		res,
	)
}
