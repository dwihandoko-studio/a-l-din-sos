package controllers

import (
	"api-dinsos/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetProvinsi(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	res, err := models.GetProvinsi()
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

func GetKabupaten(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}
	id := c.Param("id")

	res, err := models.GetKabupaten(id)
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

func GetKecamatan(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}
	id := c.Param("id")

	res, err := models.GetKecamatan(id)
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

func GetKelurahan(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}
	id := c.Param("id")

	res, err := models.GetKelurahan(id)
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
