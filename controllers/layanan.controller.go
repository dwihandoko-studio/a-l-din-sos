package controllers

import (
	"api-dinsos/helpers"
	"api-dinsos/middleware"
	"api-dinsos/models"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func PostLayananPbi(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userId := claims.Id

	var fileLain int
	var fileKtp int
	var fileKk int
	var fileFotorumah int
	var fileSktm int

	// var keperluan string
	// var keperluan_lain string

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	// 	fmt.Println("PARSING REQUEST FORM")
	keperluan := c.FormValue("keperluan")
	keperluan_lain := c.FormValue("keperluan_lain")
	// } else {
	// 	//json_map has the JSON Payload decoded into a map
	// 	fmt.Println("PARSING REQUEST JSON")
	// 	keperluan = fmt.Sprintf("%s", json_map["keperluan"])
	// 	keperluan_lain = fmt.Sprintf("%s", json_map["keperluan_lain"])
	// }

	if keperluan == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "jenis rekomendasi tidak boleh kosong."},
		)
	}

	if keperluan == "Lainnya" {
		if keperluan_lain == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Jenis rekomendasi tidak boleh kosong."},
			)
		}
	}

	filelainnya, err := c.FormFile("lainnya")
	if err != nil {
		fmt.Println(err.Error())
		fileLain = 1
	} else {

	}

	filektp, err := c.FormFile("ktp")
	if err != nil {
		fmt.Println(err.Error())
		fileKtp = 1
	} else {

	}

	filekk, err := c.FormFile("kk")
	if err != nil {
		fmt.Println(err.Error())
		fileKk = 1
	} else {

	}

	filefotorumah, err := c.FormFile("fotorumah")
	if err != nil {
		fmt.Println(err.Error())
		fileFotorumah = 1
	} else {

	}

	filesktm, err := c.FormFile("sktm")
	if err != nil {
		fmt.Println(err.Error())
		fileSktm = 1
	} else {

	}

	if fileLain == 0 {
		if string(filepath.Ext(filelainnya.Filename)) == ".jpg" ||
			string(filepath.Ext(filelainnya.Filename)) == ".pdf" ||
			string(filepath.Ext(filelainnya.Filename)) == ".png" ||
			string(filepath.Ext(filelainnya.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileKtp == 0 {
		if string(filepath.Ext(filektp.Filename)) == ".jpg" ||
			string(filepath.Ext(filektp.Filename)) == ".pdf" ||
			string(filepath.Ext(filektp.Filename)) == ".png" ||
			string(filepath.Ext(filektp.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file ktp tidak diizinkan"},
			)
		}
	}

	if fileKk == 0 {
		if string(filepath.Ext(filekk.Filename)) == ".jpg" ||
			string(filepath.Ext(filekk.Filename)) == ".pdf" ||
			string(filepath.Ext(filekk.Filename)) == ".png" ||
			string(filepath.Ext(filekk.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file kk tidak diizinkan"},
			)
		}
	}

	if fileFotorumah == 0 {
		if string(filepath.Ext(filefotorumah.Filename)) == ".jpg" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".pdf" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".png" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file foto rumah tidak diizinkan"},
			)
		}
	}

	if fileSktm == 0 {
		if string(filepath.Ext(filesktm.Filename)) == ".jpg" ||
			string(filepath.Ext(filesktm.Filename)) == ".pdf" ||
			string(filepath.Ext(filesktm.Filename)) == ".png" ||
			string(filepath.Ext(filesktm.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file sktm tidak diizinkan"},
			)
		}
	}

	result, err := models.GetUserDetail(userId)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	filenamelainnya := ""
	var fileLocationlainnya string

	filenamektp := ""
	var fileLocationktp string

	filenamekk := ""
	var fileLocationkk string

	filenamefotorumah := ""
	var fileLocationfotorumah string

	filenamesktm := ""
	var fileLocationsktm string

	if fileFotorumah == 0 || fileKk == 0 || fileKtp == 0 || fileLain == 0 || fileSktm == 0 {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		if fileLain == 0 {
			srclainnya, err := filelainnya.Open()
			if err != nil {
				return err
			}
			defer srclainnya.Close()

			newfilenamelainnya := helpers.GenerateFilename("LAINNYA" + *result.Nik)
			filenamelainnya = fmt.Sprintf("%s%s", newfilenamelainnya, filepath.Ext(filelainnya.Filename))

			fileLocationlainnya = filepath.Join(dir, "uploads/layanan/lainnya", filenamelainnya)
			targetFilelainnya, err := os.OpenFile(fileLocationlainnya, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilelainnya.Close()

			// Copy
			if _, err = io.Copy(targetFilelainnya, srclainnya); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKtp == 0 {
			srcktp, err := filektp.Open()
			if err != nil {
				return err
			}
			defer srcktp.Close()

			newfilenamektp := helpers.GenerateFilename("KTP" + *result.Nik)
			filenamektp = fmt.Sprintf("%s%s", newfilenamektp, filepath.Ext(filektp.Filename))

			fileLocationktp = filepath.Join(dir, "uploads/layanan/ktp", filenamektp)
			targetFilektp, err := os.OpenFile(fileLocationktp, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilektp.Close()

			// Copy
			if _, err = io.Copy(targetFilektp, srcktp); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKk == 0 {
			srckk, err := filekk.Open()
			if err != nil {
				return err
			}
			defer srckk.Close()

			newfilenamekk := helpers.GenerateFilename("KK" + *result.Nik)
			filenamekk = fmt.Sprintf("%s%s", newfilenamekk, filepath.Ext(filekk.Filename))

			fileLocationkk = filepath.Join(dir, "uploads/layanan/kk", filenamekk)
			targetFilekk, err := os.OpenFile(fileLocationkk, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilekk.Close()

			// Copy
			if _, err = io.Copy(targetFilekk, srckk); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileFotorumah == 0 {
			srcfotorumah, err := filefotorumah.Open()
			if err != nil {
				return err
			}
			defer srcfotorumah.Close()

			newfilenamefotorumah := helpers.GenerateFilename("FOTORUMAH-" + *result.Nik)
			filenamefotorumah = fmt.Sprintf("%s%s", newfilenamefotorumah, filepath.Ext(filefotorumah.Filename))

			fileLocationfotorumah = filepath.Join(dir, "uploads/layanan/fotorumah", filenamefotorumah)
			targetFilefotorumah, err := os.OpenFile(fileLocationfotorumah, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilefotorumah.Close()

			// Copy
			if _, err = io.Copy(targetFilefotorumah, srcfotorumah); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileSktm == 0 {
			srcsktm, err := filesktm.Open()
			if err != nil {
				return err
			}
			defer srcsktm.Close()

			newfilenamesktm := helpers.GenerateFilename("SKTM-" + *result.Nik)
			filenamesktm = fmt.Sprintf("%s%s", newfilenamesktm, filepath.Ext(filesktm.Filename))

			fileLocationsktm = filepath.Join(dir, "uploads/layanan/sktm", filenamesktm)
			targetFilesktm, err := os.OpenFile(fileLocationsktm, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilesktm.Close()

			// Copy
			if _, err = io.Copy(targetFilesktm, srcsktm); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}
	}

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	kode_permohonan := helpers.GenerateKodePermohonan("PBI-" + *result.Nik)

	var keperluan_fix string
	if keperluan == "Lainnya" {
		keperluan_fix = keperluan_lain
	} else {
		keperluan_fix = keperluan
	}

	if _, err := models.PostPbi(userId, kode_permohonan, *result.Kelurahan, "kadis", *result.Nik, *result.Fullname, keperluan_fix, "PBI", "0", currentTime.Format("2006-01-02 15:04:05"), filenamektp, filenamekk, filenamefotorumah, filenamelainnya, filenamesktm); err != nil {
		if filenamelainnya != "" {
			el := os.Remove(fileLocationlainnya)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE LAINNYA")
			}
		}
		if filenamektp != "" {
			el := os.Remove(fileLocationktp)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KTP")
			}
		}
		if filenamekk != "" {
			el := os.Remove(fileLocationkk)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KK")
			}
		}
		if filenamefotorumah != "" {
			el := os.Remove(fileLocationfotorumah)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE FOTO RUMAH")
			}
		}
		if filenamesktm != "" {
			el := os.Remove(fileLocationsktm)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE SKTM")
			}
		}
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	// resultA, err := modelpeserta.InsertAktifitas("Mendaftar via Jalur Prestasi, untuk diverifikasi berkas oleh sekolah tujuan.", "Daftar Jalur Prestasi", "submit", id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(resultA)
	// }

	return c.JSON(http.StatusOK,
		models.Response{Status: 200, Message: "Permohonan Berhasil di Ajukan."},
	)
}

func PostLayananSktm(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userId := claims.Id

	var fileLain int
	var fileKtp int
	var fileKk int
	var fileFotorumah int

	// var keperluan string
	// var keperluan_lain string

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	// 	fmt.Println("PARSING REQUEST FORM")
	keperluan := c.FormValue("keperluan")
	keperluan_lain := c.FormValue("keperluan_lain")
	// } else {
	// 	//json_map has the JSON Payload decoded into a map
	// 	fmt.Println("PARSING REQUEST JSON")
	// 	keperluan = fmt.Sprintf("%s", json_map["keperluan"])
	// 	keperluan_lain = fmt.Sprintf("%s", json_map["keperluan_lain"])
	// }

	if keperluan == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Keperluan tidak boleh kosong."},
		)
	}

	if keperluan == "Lainnya" {
		if keperluan_lain == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Keperluan keterangan DTKS tidak boleh kosong."},
			)
		}
	}

	filelainnya, err := c.FormFile("lainnya")
	if err != nil {
		fmt.Println(err.Error())
		fileLain = 1
	} else {

	}

	filektp, err := c.FormFile("ktp")
	if err != nil {
		fmt.Println(err.Error())
		fileKtp = 1
	} else {

	}

	filekk, err := c.FormFile("kk")
	if err != nil {
		fmt.Println(err.Error())
		fileKk = 1
	} else {

	}

	filefotorumah, err := c.FormFile("fotorumah")
	if err != nil {
		fmt.Println(err.Error())
		fileFotorumah = 1
	} else {

	}

	if fileLain == 0 {
		if string(filepath.Ext(filelainnya.Filename)) == ".jpg" ||
			string(filepath.Ext(filelainnya.Filename)) == ".pdf" ||
			string(filepath.Ext(filelainnya.Filename)) == ".png" ||
			string(filepath.Ext(filelainnya.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileKtp == 0 {
		if string(filepath.Ext(filektp.Filename)) == ".jpg" ||
			string(filepath.Ext(filektp.Filename)) == ".pdf" ||
			string(filepath.Ext(filektp.Filename)) == ".png" ||
			string(filepath.Ext(filektp.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileKk == 0 {
		if string(filepath.Ext(filekk.Filename)) == ".jpg" ||
			string(filepath.Ext(filekk.Filename)) == ".pdf" ||
			string(filepath.Ext(filekk.Filename)) == ".png" ||
			string(filepath.Ext(filekk.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileFotorumah == 0 {
		if string(filepath.Ext(filefotorumah.Filename)) == ".jpg" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".pdf" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".png" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	result, err := models.GetUserDetail(userId)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	filenamelainnya := ""
	var fileLocationlainnya string

	filenamektp := ""
	var fileLocationktp string

	filenamekk := ""
	var fileLocationkk string

	filenamefotorumah := ""
	var fileLocationfotorumah string

	if fileFotorumah == 0 || fileKk == 0 || fileKtp == 0 || fileLain == 0 {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		if fileLain == 0 {
			srclainnya, err := filelainnya.Open()
			if err != nil {
				return err
			}
			defer srclainnya.Close()

			newfilenamelainnya := helpers.GenerateFilename("LAINNYA" + *result.Nik)
			filenamelainnya = fmt.Sprintf("%s%s", newfilenamelainnya, filepath.Ext(filelainnya.Filename))

			fileLocationlainnya = filepath.Join(dir, "uploads/layanan/lainnya", filenamelainnya)
			targetFilelainnya, err := os.OpenFile(fileLocationlainnya, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilelainnya.Close()

			// Copy
			if _, err = io.Copy(targetFilelainnya, srclainnya); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKtp == 0 {
			srcktp, err := filektp.Open()
			if err != nil {
				return err
			}
			defer srcktp.Close()

			newfilenamektp := helpers.GenerateFilename("KTP" + *result.Nik)
			filenamektp = fmt.Sprintf("%s%s", newfilenamektp, filepath.Ext(filektp.Filename))

			fileLocationktp = filepath.Join(dir, "uploads/layanan/ktp", filenamektp)
			targetFilektp, err := os.OpenFile(fileLocationktp, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilektp.Close()

			// Copy
			if _, err = io.Copy(targetFilektp, srcktp); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKk == 0 {
			srckk, err := filekk.Open()
			if err != nil {
				return err
			}
			defer srckk.Close()

			newfilenamekk := helpers.GenerateFilename("KK" + *result.Nik)
			filenamekk = fmt.Sprintf("%s%s", newfilenamekk, filepath.Ext(filekk.Filename))

			fileLocationkk = filepath.Join(dir, "uploads/layanan/kk", filenamekk)
			targetFilekk, err := os.OpenFile(fileLocationkk, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilekk.Close()

			// Copy
			if _, err = io.Copy(targetFilekk, srckk); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileFotorumah == 0 {
			srcfotorumah, err := filefotorumah.Open()
			if err != nil {
				return err
			}
			defer srcfotorumah.Close()

			newfilenamefotorumah := helpers.GenerateFilename("FOTORUMAH-" + *result.Nik)
			filenamefotorumah = fmt.Sprintf("%s%s", newfilenamefotorumah, filepath.Ext(filefotorumah.Filename))

			fileLocationfotorumah = filepath.Join(dir, "uploads/layanan/fotorumah", filenamefotorumah)
			targetFilefotorumah, err := os.OpenFile(fileLocationfotorumah, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilefotorumah.Close()

			// Copy
			if _, err = io.Copy(targetFilefotorumah, srcfotorumah); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}
	}

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	kode_permohonan := helpers.GenerateKodePermohonan("SKTM-" + *result.Nik)

	var keperluan_fix string
	if keperluan == "Lainnya" {
		keperluan_fix = keperluan_lain
	} else {
		keperluan_fix = keperluan
	}

	if _, err := models.PostSktm(userId, kode_permohonan, *result.Kelurahan, "kakam", *result.Nik, *result.Fullname, keperluan_fix, "SKTM", "0", currentTime.Format("2006-01-02 15:04:05"), filenamektp, filenamekk, filenamefotorumah, filenamelainnya); err != nil {
		if filenamelainnya != "" {
			el := os.Remove(fileLocationlainnya)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE LAINNYA")
			}
		}
		if filenamektp != "" {
			el := os.Remove(fileLocationktp)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KTP")
			}
		}
		if filenamekk != "" {
			el := os.Remove(fileLocationkk)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KK")
			}
		}
		if filenamefotorumah != "" {
			el := os.Remove(fileLocationfotorumah)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE FOTO RUMAH")
			}
		}
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	// resultA, err := modelpeserta.InsertAktifitas("Mendaftar via Jalur Prestasi, untuk diverifikasi berkas oleh sekolah tujuan.", "Daftar Jalur Prestasi", "submit", id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(resultA)
	// }

	return c.JSON(http.StatusOK,
		models.Response{Status: 200, Message: "Permohonan Berhasil di Ajukan."},
	)
}

func PostLayananDtks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userId := claims.Id

	var fileLain int
	var fileKtp int
	var fileKk int
	var fileFotorumah int

	// var keperluan string
	// var keperluan_lain string

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	// 	fmt.Println("PARSING REQUEST FORM")
	keperluan := c.FormValue("keperluan")
	keperluan_lain := c.FormValue("keperluan_lain")
	// } else {
	// 	//json_map has the JSON Payload decoded into a map
	// 	fmt.Println("PARSING REQUEST JSON")
	// 	keperluan = fmt.Sprintf("%s", json_map["keperluan"])
	// 	keperluan_lain = fmt.Sprintf("%s", json_map["keperluan_lain"])
	// }

	if keperluan == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Keperluan tidak boleh kosong."},
		)
	}

	if keperluan == "Lainnya" {
		if keperluan_lain == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Keperluan keterangan DTKS tidak boleh kosong."},
			)
		}
	}

	filelainnya, err := c.FormFile("lainnya")
	if err != nil {
		fmt.Println(err.Error())
		fileLain = 1
	} else {

	}

	filektp, err := c.FormFile("ktp")
	if err != nil {
		fmt.Println(err.Error())
		fileKtp = 1
	} else {

	}

	filekk, err := c.FormFile("kk")
	if err != nil {
		fmt.Println(err.Error())
		fileKk = 1
	} else {

	}

	filefotorumah, err := c.FormFile("fotorumah")
	if err != nil {
		fmt.Println(err.Error())
		fileFotorumah = 1
	} else {

	}

	if fileLain == 0 {
		if string(filepath.Ext(filelainnya.Filename)) == ".jpg" ||
			string(filepath.Ext(filelainnya.Filename)) == ".pdf" ||
			string(filepath.Ext(filelainnya.Filename)) == ".png" ||
			string(filepath.Ext(filelainnya.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileKtp == 0 {
		if string(filepath.Ext(filektp.Filename)) == ".jpg" ||
			string(filepath.Ext(filektp.Filename)) == ".pdf" ||
			string(filepath.Ext(filektp.Filename)) == ".png" ||
			string(filepath.Ext(filektp.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileKk == 0 {
		if string(filepath.Ext(filekk.Filename)) == ".jpg" ||
			string(filepath.Ext(filekk.Filename)) == ".pdf" ||
			string(filepath.Ext(filekk.Filename)) == ".png" ||
			string(filepath.Ext(filekk.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	if fileFotorumah == 0 {
		if string(filepath.Ext(filefotorumah.Filename)) == ".jpg" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".pdf" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".png" ||
			string(filepath.Ext(filefotorumah.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file lain tidak diizinkan"},
			)
		}
	}

	result, err := models.GetUserDetail(userId)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	filenamelainnya := ""
	var fileLocationlainnya string

	filenamektp := ""
	var fileLocationktp string

	filenamekk := ""
	var fileLocationkk string

	filenamefotorumah := ""
	var fileLocationfotorumah string

	if fileFotorumah == 0 || fileKk == 0 || fileKtp == 0 || fileLain == 0 {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		if fileLain == 0 {
			srclainnya, err := filelainnya.Open()
			if err != nil {
				return err
			}
			defer srclainnya.Close()

			newfilenamelainnya := helpers.GenerateFilename("LAINNYA" + *result.Nik)
			filenamelainnya = fmt.Sprintf("%s%s", newfilenamelainnya, filepath.Ext(filelainnya.Filename))

			fileLocationlainnya = filepath.Join(dir, "uploads/layanan/lainnya", filenamelainnya)
			targetFilelainnya, err := os.OpenFile(fileLocationlainnya, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilelainnya.Close()

			// Copy
			if _, err = io.Copy(targetFilelainnya, srclainnya); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKtp == 0 {
			srcktp, err := filektp.Open()
			if err != nil {
				return err
			}
			defer srcktp.Close()

			newfilenamektp := helpers.GenerateFilename("KTP" + *result.Nik)
			filenamektp = fmt.Sprintf("%s%s", newfilenamektp, filepath.Ext(filektp.Filename))

			fileLocationktp = filepath.Join(dir, "uploads/layanan/ktp", filenamektp)
			targetFilektp, err := os.OpenFile(fileLocationktp, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilektp.Close()

			// Copy
			if _, err = io.Copy(targetFilektp, srcktp); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKk == 0 {
			srckk, err := filekk.Open()
			if err != nil {
				return err
			}
			defer srckk.Close()

			newfilenamekk := helpers.GenerateFilename("KK" + *result.Nik)
			filenamekk = fmt.Sprintf("%s%s", newfilenamekk, filepath.Ext(filekk.Filename))

			fileLocationkk = filepath.Join(dir, "uploads/layanan/kk", filenamekk)
			targetFilekk, err := os.OpenFile(fileLocationkk, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilekk.Close()

			// Copy
			if _, err = io.Copy(targetFilekk, srckk); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileFotorumah == 0 {
			srcfotorumah, err := filefotorumah.Open()
			if err != nil {
				return err
			}
			defer srcfotorumah.Close()

			newfilenamefotorumah := helpers.GenerateFilename("FOTORUMAH-" + *result.Nik)
			filenamefotorumah = fmt.Sprintf("%s%s", newfilenamefotorumah, filepath.Ext(filefotorumah.Filename))

			fileLocationfotorumah = filepath.Join(dir, "uploads/layanan/fotorumah", filenamefotorumah)
			targetFilefotorumah, err := os.OpenFile(fileLocationfotorumah, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilefotorumah.Close()

			// Copy
			if _, err = io.Copy(targetFilefotorumah, srcfotorumah); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}
	}

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	kode_permohonan := helpers.GenerateKodePermohonan("SKDTKS-" + *result.Nik)

	var keperluan_fix string
	if keperluan == "Lainnya" {
		keperluan_fix = keperluan_lain
	} else {
		keperluan_fix = keperluan
	}

	if _, err := models.PostDtks(userId, kode_permohonan, *result.Kelurahan, "kadis", *result.Nik, *result.Fullname, keperluan_fix, "SKDTKS", "0", currentTime.Format("2006-01-02 15:04:05"), filenamektp, filenamekk, filenamefotorumah, filenamelainnya); err != nil {
		if filenamelainnya != "" {
			el := os.Remove(fileLocationlainnya)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE LAINNYA")
			}
		}
		if filenamektp != "" {
			el := os.Remove(fileLocationktp)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KTP")
			}
		}
		if filenamekk != "" {
			el := os.Remove(fileLocationkk)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KK")
			}
		}
		if filenamefotorumah != "" {
			el := os.Remove(fileLocationfotorumah)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE FOTO RUMAH")
			}
		}
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	// resultA, err := modelpeserta.InsertAktifitas("Mendaftar via Jalur Prestasi, untuk diverifikasi berkas oleh sekolah tujuan.", "Daftar Jalur Prestasi", "submit", id)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	fmt.Println(resultA)
	// }

	return c.JSON(http.StatusOK,
		models.Response{Status: 200, Message: "Permohonan Berhasil di Ajukan."},
	)
}

func TestPostLayananDtks(c echo.Context) error {
	// var requestData map[string]interface{}
	// err := c.Bind(requestData)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal membaca data JSON"})
	// }

	keperluan := c.FormValue("keperluan")
	log.Println(keperluan)
	// Mengambil file dari request
	file, err := c.FormFile("ktp")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Gagal mengambil file"})
	}

	dst, err := os.Create("uploads/" + file.Filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan file"})
	}
	defer dst.Close()

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuka file"})
	}
	defer src.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyalin file"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Ggagal"})
}
