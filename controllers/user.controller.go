package controllers

import (
	"api-dinsos/helpers"
	"api-dinsos/middleware"
	"api-dinsos/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*middleware.JwtCustomClaims)
	// id := claims.Id
	var start string = "0"
	var length string = "10"
	var keyword string
	var role string

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		// fmt.Println("ERROR GUYS")
		start = c.FormValue("start")
		length = c.FormValue("length")
		keyword = c.FormValue("keyword")
		role = c.FormValue("role")
	} else {
		//json_map has the JSON Payload decoded into a map
		start = fmt.Sprintf("%s", json_map["start"])
		length = fmt.Sprintf("%s", json_map["length"])
		keyword = fmt.Sprintf("%s", json_map["keyword"])
		role = fmt.Sprintf("%s", json_map["role"])
	}

	sql := "WHERE"

	if r := role; r != "" {
		sql = fmt.Sprintf("%s role_user =%s", sql, r)
	} else {
		sql = fmt.Sprintf("%s role_user IS NOT null", sql)
	}

	if s := keyword; s != "" {
		sql = fmt.Sprintf("%s AND (fullname LIKE '%%%s%%' OR email LIKE '%%%s%%' OR npsn LIKE '%%%s%%' OR nip LIKE '%%%s%%')", sql, s, s, s, s)
	}

	if p := start; p != "" {
		start = p
	} else {
		start = "0"
	}

	if l := length; l != "" {
		length = l
	} else {
		length = "10"
	}

	res, err := models.GetUsers(sql, start, length)
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

func GetUser(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid .."},
		)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	id := claims.Id

	result, err := models.GetUserDetail(id)
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

func GetUserDetail(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	id := c.Param("id")

	result, err := models.GetUserDetail(id)
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

func PostChangePassword(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")
	var old_password string
	var new_password string
	var ulangi_new_password string

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	id := claims.Id

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		// fmt.Println("ERROR GUYS")
		old_password = c.FormValue("old_password")
		new_password = c.FormValue("new_password")
		ulangi_new_password = c.FormValue("ulangi_new_password")
	} else {
		//json_map has the JSON Payload decoded into a map
		old_password = fmt.Sprintf("%s", json_map["old_password"])
		new_password = fmt.Sprintf("%s", json_map["new_password"])
		ulangi_new_password = fmt.Sprintf("%s", json_map["ulangi_new_password"])
	}

	result, err := models.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: err.Error()},
		)
	}

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)

	if old_password == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Kata sandi lama tidak boleh kosong."},
		)
	}

	if new_password == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Kata sandi baru tidak boleh kosong."},
		)
	}

	if ulangi_new_password == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Ulangi Kata sandi baru tidak boleh kosong."},
		)
	}

	if new_password != ulangi_new_password {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "New Password tidak valid."},
		)
	}

	ress, err := models.UpdatePassword(result.Id, old_password, new_password, currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: err.Error()},
		)
	}

	models.InsertAktifitas("Mengubah kata sandi", "Update Kata Sandi", "update", id)
	// if errs != nil {
	// 	fmt.Println(err.Error())
	// }

	return c.JSON(http.StatusOK, ress)
}

func EditFotoUser(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	file, err := c.FormFile("lampiran")
	if err != nil {
		// fmt.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "File upload tidak boleh kosong"},
		)
	}

	if string(filepath.Ext(file.Filename)) == ".jpg" || string(filepath.Ext(file.Filename)) == ".png" || string(filepath.Ext(file.Filename)) == ".jpeg" {

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*middleware.JwtCustomClaims)
		id := claims.Id

		_, err := models.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: err.Error()},
			)
		}

		dir, err := os.Getwd()
		if err != nil {
			// fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		// dst, err := os.Create(file.Filename)
		// if err != nil {
		// 	return err
		// }
		// defer dst.Close()
		loca, _ := time.LoadLocation("Asia/Jakarta")

		currentTime := time.Now().In(loca)

		newfilename := helpers.GenerateFilename("FOTO")
		filename := fmt.Sprintf("%s%s", newfilename, filepath.Ext(file.Filename))

		fileLocation := filepath.Join(dir, "uploads/user", filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			// fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}
		defer targetFile.Close()

		// Copy
		if _, err = io.Copy(targetFile, src); err != nil {
			// fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		if _, err := models.UpdateFotoProfil(id, filename, currentTime.Format("2006-01-02 15:04:05")); err != nil {
			os.Remove(fileLocation)
			// if e != nil {
			// 	fmt.Println("GAGAL MENGHAPUS FILE UPLOAD")
			// }
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: err.Error()},
			)
		}

		// models.InsertAktifitas("Mengubah profil", "Update Foto Profil", "update", id)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	fmt.Println(errs)
		// }

		return c.JSON(http.StatusOK,
			models.Response{Status: 200, Message: "Profil berhasil diubah."},
		)
	} else {
		// fmt.Println(filepath.Ext(file.Filename))
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Type file upload tidak diizinkan."},
		)
	}
}

func PostUser(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid .."},
		)
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	id := claims.Id

	var provinsi string
	var kabupaten string
	var kecamatan string
	var kelurahan string
	var alamat string
	var nama string
	var kk string
	var tempat_lahir string
	var tgl_lahir string
	var jenis_kelamin string

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		// fmt.Println("PARSING REQUEST FORM")
		nama = c.FormValue("nama")
		kk = c.FormValue("kk")
		provinsi = c.FormValue("provinsi")
		kabupaten = c.FormValue("kabupaten")
		kecamatan = c.FormValue("kecamatan")
		kelurahan = c.FormValue("kelurahan")
		alamat = c.FormValue("alamat")
		tempat_lahir = c.FormValue("tempat_lahir")
		tgl_lahir = c.FormValue("tgl_lahir")
		jenis_kelamin = c.FormValue("jenis_kelamin")
	} else {
		//json_map has the JSON Payload decoded into a map
		// fmt.Println("PARSING REQUEST JSON")
		nama = fmt.Sprintf("%s", json_map["nama"])
		kk = fmt.Sprintf("%s", json_map["kk"])
		provinsi = fmt.Sprintf("%s", json_map["provinsi"])
		kabupaten = fmt.Sprintf("%s", json_map["kabupaten"])
		kecamatan = fmt.Sprintf("%s", json_map["kecamatan"])
		kelurahan = fmt.Sprintf("%s", json_map["kelurahan"])
		alamat = fmt.Sprintf("%s", json_map["alamat"])
		tempat_lahir = fmt.Sprintf("%s", json_map["tempat_lahir"])
		tgl_lahir = fmt.Sprintf("%s", json_map["tgl_lahir"])
		jenis_kelamin = fmt.Sprintf("%s", json_map["jenis_kelamin"])
	}

	if nama == "" || kk == "" || provinsi == "" || kabupaten == "" || kecamatan == "" || kelurahan == "" || alamat == "" || tempat_lahir == "" || tgl_lahir == "" || jenis_kelamin == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Nama, KK, Provinsi, Kabupaten, Kecamatan, Kelurahan, Alamat, Tempat lahir, Tanggal lahir dan Jenis kelamin tidak boleh kosong."},
		)
	}

	_, errS := models.GetUserDetail(id)
	if errS != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	// if _, err = io.Copy(targetFile, src); err != nil {
	// 	fmt.Println(err.Error())
	// 	return c.JSON(http.StatusOK,
	// 		models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
	// 	)
	// }
	// if result.Data.(modelpeserta.Userdetail).Latitude != nil {
	// 	if latitude != *result.Data.(modelpeserta.Userdetail).Latitude && longitude != *result.Data.(modelpeserta.Userdetail).Longitude {
	// 		if result.Data.(modelpeserta.Userdetail).Edited_map == 1 {
	// 			return c.JSON(http.StatusOK,
	// 				models.Response{Status: 400, Message: "Mohon maaf, system mendeteksi anda sudah melakukan update koordinat tempat tinggal sebelumnya. Selanjutnya untuk dapat mengubah koordinat tempat tinggal, silahkan hubungi admin PPDB Dinas."},
	// 			)
	// 		}
	// 	}
	// }

	// if latitude != *result.Data.(modelpeserta.Userdetail).Latitude && longitude != *result.Data.(modelpeserta.Userdetail).Longitude {
	// 	if result.Data.(modelpeserta.Userdetail).Edited_map == 1 {
	// 		return c.JSON(http.StatusOK,
	// 			models.Response{Status: 400, Message: "Mohon maaf, system mendeteksi anda sudah melakukan update koordinat tempat tinggal sebelumnya. Selanjutnya untuk dapat mengubah koordinat tempat tinggal, silahkan hubungi admin PPDB Dinas."},
	// 		)
	// 	}
	// }

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)

	ress, err := models.UpdateUser(id, provinsi, kabupaten, kecamatan, kelurahan, kk, alamat, tempat_lahir, tgl_lahir, nama, jenis_kelamin, currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}
	fmt.Println(id)

	// resultA, err := modelpeserta.InsertAktifitas("Mengubah alamat tempat tinggal", "Update Data Alamat", "update", id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(resultA)
	// }

	return c.JSON(http.StatusOK, ress)
}
