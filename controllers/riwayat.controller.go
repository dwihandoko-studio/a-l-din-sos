package controllers

import (
	"api-dinsos/middleware"
	"api-dinsos/models"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetRiwayatLayanan(c echo.Context) error {
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
		sql = fmt.Sprintf("%s AND (layanan LIKE '%%%s%%' OR jenis LIKE '%%%s%%')", sql, s, s)
	}

	if p := c.QueryParam("page"); p != "" {
		page = p
	}

	res, err := models.GetRiwayatLayanan(sql, page)
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

func GetDetailRiwayatLayanan(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	idUser := claims.Id

	_, err := models.GetUser(idUser)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	id := c.Param("id")
	layanan := c.Param("layanan")
	status := c.Param("status")

	switch layanan {
	case "SKTM":
		result, err := models.GetDetailSktm(id, status)
		if err != nil {
			// fmt.Println("Error getting data")
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: "success",
				Data:    result,
			})
	case "PBI":
		result, err := models.GetDetailPbi(id, status)
		if err != nil {
			// fmt.Println("Error getting data")
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: "success",
				Data:    result,
			})
	case "LKS":
		result, err := models.GetDetailLks(id, status)
		if err != nil {
			// fmt.Println("Error getting data")
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: "success",
				Data:    result,
			})
	case "SKDTKS":
		result, err := models.GetDetailSkdtks(id, status)
		if err != nil {
			// fmt.Println("Error getting data")
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: "success",
				Data:    result,
			})
	default:
		result, err := models.GetDetailSktm(id, status)
		if err != nil {
			// fmt.Println("Error getting data")
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: "success",
				Data:    result,
			})
	}

}

func GetRiwayatPengaduan(c echo.Context) error {
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
		sql = fmt.Sprintf("%s AND (kategori LIKE '%%%s%%' OR kode_aduan LIKE '%%%s%%')", sql, s, s)
	}

	if p := c.QueryParam("page"); p != "" {
		page = p
	}

	res, err := models.GetRiwayatPengaduan(sql, page)
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

func GetDetailRiwayatPengaduan(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	idUser := claims.Id

	_, err := models.GetUser(idUser)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	id := c.Param("id")
	// kategori := c.Param("kategori")
	status := c.Param("status")

	// switch kategori {
	// case "SKTM":
	result, err := models.GetDetailPengaduan(id, status)
	if err != nil {
		// fmt.Println("Error getting data")
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "success",
			Data:    result,
		})
	// case "PBI":
	// 	result, err := models.GetDetailPbi(id, status)
	// 	if err != nil {
	// 		// fmt.Println("Error getting data")
	// 		return c.JSON(http.StatusOK,
	// 			models.Response{Status: 404, Message: err.Error()},
	// 		)
	// 	}

	// 	return c.JSON(http.StatusOK,
	// 		models.Response{
	// 			Status:  200,
	// 			Message: "success",
	// 			Data:    result,
	// 		})
	// case "LKS":
	// 	result, err := models.GetDetailLks(id, status)
	// 	if err != nil {
	// 		// fmt.Println("Error getting data")
	// 		return c.JSON(http.StatusOK,
	// 			models.Response{Status: 404, Message: err.Error()},
	// 		)
	// 	}

	// 	return c.JSON(http.StatusOK,
	// 		models.Response{
	// 			Status:  200,
	// 			Message: "success",
	// 			Data:    result,
	// 		})
	// case "SKDTKS":
	// 	result, err := models.GetDetailSkdtks(id, status)
	// 	if err != nil {
	// 		// fmt.Println("Error getting data")
	// 		return c.JSON(http.StatusOK,
	// 			models.Response{Status: 404, Message: err.Error()},
	// 		)
	// 	}

	// 	return c.JSON(http.StatusOK,
	// 		models.Response{
	// 			Status:  200,
	// 			Message: "success",
	// 			Data:    result,
	// 		})
	// default:
	// 	result, err := models.GetDetailSktm(id, status)
	// 	if err != nil {
	// 		// fmt.Println("Error getting data")
	// 		return c.JSON(http.StatusOK,
	// 			models.Response{Status: 404, Message: err.Error()},
	// 		)
	// 	}

	// 	return c.JSON(http.StatusOK,
	// 		models.Response{
	// 			Status:  200,
	// 			Message: "success",
	// 			Data:    result,
	// 		})
	// }

}
