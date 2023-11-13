package controllers

import (
	"api-dinsos/middleware"
	"api-dinsos/models"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetAktifitas(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	id := claims.Id

	_, err := models.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	page := "1"

	sql := fmt.Sprintf("WHERE user_id ='%s'", id)

	if s := c.QueryParam("keyword"); s != "" {
		sql = fmt.Sprintf("%s AND (keterangan LIKE '%%%s%%' OR aksi LIKE '%%%s%%')", sql, s, s)
	}

	if p := c.QueryParam("page"); p != "" {
		page = p
	}

	res, err := models.GetAktifitas(sql, page)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Data tidak ditemukan."},
		)
	}

	if res.Total > 0 {
		if len(res.Data) > 0 {
			return c.JSON(http.StatusOK,
				models.Response{
					Status:  200,
					Message: "Data ditemukan.",
					Data:    res,
				})
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  204,
				Message: "Data ditemukan.",
			})
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  204,
			Message: "Tidak ada data ditemukan.",
			Data:    res,
		})
}
