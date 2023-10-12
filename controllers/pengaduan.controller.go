package controllers

import (
	"api-dinsos/helpers"
	"api-dinsos/middleware"
	"api-dinsos/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func PostPengaduan(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userId := claims.Id

	var filePengaduan1 int
	var filePengaduan2 int
	var filePengaduan3 int
	var filePengaduan4 int
	var filePengaduan5 int

	// var keperluan string
	// var keperluan_lain string

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	// 	fmt.Println("PARSING REQUEST FORM")
	keperluan := c.FormValue("keperluan")
	keperluan_lain := c.FormValue("keperluan_lain")
	subject := c.FormValue("subject")
	nama := c.FormValue("nama")
	nik := c.FormValue("nik")
	nohp := c.FormValue("nohp")
	alamat := c.FormValue("alamat")
	kecamatan := c.FormValue("kecamatan")
	kelurahan := c.FormValue("kelurahan")
	uraian := c.FormValue("uraian")
	// } else {
	// 	//json_map has the JSON Payload decoded into a map
	// 	fmt.Println("PARSING REQUEST JSON")
	// 	keperluan = fmt.Sprintf("%s", json_map["keperluan"])
	// 	keperluan_lain = fmt.Sprintf("%s", json_map["keperluan_lain"])
	// }

	if keperluan == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "jenis pengaduan tidak boleh kosong."},
		)
	}

	if keperluan == "Lainnya" {
		if keperluan_lain == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Jenis pengaduan tidak boleh kosong."},
			)
		}
	}

	if subject == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "subject aduan tidak boleh kosong."},
		)
	}

	if uraian == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "uraian aduan tidak boleh kosong."},
		)
	}

	if subject == "2" {
		if nama == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "nama aduan tidak boleh kosong."},
			)
		}
		if nik == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "nik aduan tidak boleh kosong."},
			)
		}
		if nohp == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "nohp aduan tidak boleh kosong."},
			)
		}
		if alamat == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "alamat aduan tidak boleh kosong."},
			)
		}
		if kecamatan == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "kecamatan aduan tidak boleh kosong."},
			)
		}
		if kelurahan == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "kelurahan aduan tidak boleh kosong."},
			)
		}
	}

	filepengaduan1, err := c.FormFile("pengaduan1")
	if err != nil {
		fmt.Println(err.Error())
		filePengaduan1 = 1
	} else {

	}

	filepengaduan2, err := c.FormFile("pengaduan2")
	if err != nil {
		fmt.Println(err.Error())
		filePengaduan2 = 1
	} else {

	}

	filepengaduan3, err := c.FormFile("pengaduan3")
	if err != nil {
		fmt.Println(err.Error())
		filePengaduan3 = 1
	} else {

	}

	filepengaduan4, err := c.FormFile("pengaduan4")
	if err != nil {
		fmt.Println(err.Error())
		filePengaduan4 = 1
	} else {

	}

	filepengaduan5, err := c.FormFile("pengaduan5")
	if err != nil {
		fmt.Println(err.Error())
		filePengaduan5 = 1
	} else {

	}

	if filePengaduan1 == 0 {
		if string(filepath.Ext(filepengaduan1.Filename)) == ".jpg" ||
			string(filepath.Ext(filepengaduan1.Filename)) == ".pdf" ||
			string(filepath.Ext(filepengaduan1.Filename)) == ".png" ||
			string(filepath.Ext(filepengaduan1.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file pengaudan 1 tidak diizinkan"},
			)
		}
	}

	if filePengaduan2 == 0 {
		if string(filepath.Ext(filepengaduan2.Filename)) == ".jpg" ||
			string(filepath.Ext(filepengaduan2.Filename)) == ".pdf" ||
			string(filepath.Ext(filepengaduan2.Filename)) == ".png" ||
			string(filepath.Ext(filepengaduan2.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file pengaduan 2 tidak diizinkan"},
			)
		}
	}

	if filePengaduan3 == 0 {
		if string(filepath.Ext(filepengaduan3.Filename)) == ".jpg" ||
			string(filepath.Ext(filepengaduan3.Filename)) == ".pdf" ||
			string(filepath.Ext(filepengaduan3.Filename)) == ".png" ||
			string(filepath.Ext(filepengaduan3.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file pengaduan 3 tidak diizinkan"},
			)
		}
	}

	if filePengaduan4 == 0 {
		if string(filepath.Ext(filepengaduan4.Filename)) == ".jpg" ||
			string(filepath.Ext(filepengaduan4.Filename)) == ".pdf" ||
			string(filepath.Ext(filepengaduan4.Filename)) == ".png" ||
			string(filepath.Ext(filepengaduan4.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file pengaduan 4 tidak diizinkan"},
			)
		}
	}

	if filePengaduan5 == 0 {
		if string(filepath.Ext(filepengaduan5.Filename)) == ".jpg" ||
			string(filepath.Ext(filepengaduan5.Filename)) == ".pdf" ||
			string(filepath.Ext(filepengaduan5.Filename)) == ".png" ||
			string(filepath.Ext(filepengaduan5.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file pengaduan 5 tidak diizinkan"},
			)
		}
	}

	result, err := models.GetUserDetail(userId)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	var nama_fix string
	if subject == "2" {
		nama_fix = nama
	} else {
		nama_fix = *result.Fullname
	}

	var nik_fix string
	if subject == "2" {
		nik_fix = nik
	} else {
		nik_fix = *result.Nik
	}

	var nohp_fix string
	if subject == "2" {
		nohp_fix = nohp
	} else {
		nohp_fix = *result.No_hp
	}

	var alamat_fix string
	if subject == "2" {
		alamat_fix = alamat
	} else {
		alamat_fix = *result.Alamat
	}

	var kecamatan_fix string
	if subject == "2" {
		kecamatan_fix = kecamatan
	} else {
		kecamatan_fix = *result.Kecamatan
	}

	var kelurahan_fix string
	if subject == "2" {
		kelurahan_fix = kelurahan
	} else {
		kelurahan_fix = *result.Kelurahan
	}

	var subject_fix string
	if subject == "2" {
		subject_fix = "beda"
	} else {
		subject_fix = "sama"
	}

	filenamepengaduan1 := ""
	var fileLocationpengaduan1 string

	filenamepengaduan2 := ""
	var fileLocationpengaduan2 string

	filenamepengaduan3 := ""
	var fileLocationpengaduan3 string

	filenamepengaduan4 := ""
	var fileLocationpengaduan4 string

	filenamepengaduan5 := ""
	var fileLocationpengaduan5 string

	if filePengaduan1 == 0 || filePengaduan2 == 0 || filePengaduan3 == 0 || filePengaduan4 == 0 || filePengaduan5 == 0 {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		if filePengaduan1 == 0 {
			srcpengaduan1, err := filepengaduan1.Open()
			if err != nil {
				return err
			}
			defer srcpengaduan1.Close()

			newfilenamepengaduan1 := helpers.GenerateFilename("PENGADUAN-1" + *result.Nik)
			filenamepengaduan1 = fmt.Sprintf("%s%s", newfilenamepengaduan1, filepath.Ext(filepengaduan1.Filename))

			fileLocationpengaduan1 = filepath.Join(dir, "uploads/aduan", filenamepengaduan1)
			targetFilepengaduan1, err := os.OpenFile(fileLocationpengaduan1, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepengaduan1.Close()

			// Copy
			if _, err = io.Copy(targetFilepengaduan1, srcpengaduan1); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if filePengaduan2 == 0 {
			srcpengaduan2, err := filepengaduan2.Open()
			if err != nil {
				return err
			}
			defer srcpengaduan2.Close()

			newfilenamepengaduan2 := helpers.GenerateFilename("PENGADUAN-2" + *result.Nik)
			filenamepengaduan2 = fmt.Sprintf("%s%s", newfilenamepengaduan2, filepath.Ext(filepengaduan2.Filename))

			fileLocationpengaduan2 = filepath.Join(dir, "uploads/aduan", filenamepengaduan2)
			targetFilepengaduan2, err := os.OpenFile(fileLocationpengaduan2, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepengaduan2.Close()

			// Copy
			if _, err = io.Copy(targetFilepengaduan2, srcpengaduan2); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if filePengaduan3 == 0 {
			srcpengaduan3, err := filepengaduan3.Open()
			if err != nil {
				return err
			}
			defer srcpengaduan3.Close()

			newfilenamepengaduan3 := helpers.GenerateFilename("PENGADUAN-3" + *result.Nik)
			filenamepengaduan3 = fmt.Sprintf("%s%s", newfilenamepengaduan3, filepath.Ext(filepengaduan3.Filename))

			fileLocationpengaduan3 = filepath.Join(dir, "uploads/aduan", filenamepengaduan3)
			targetFilepengaduan3, err := os.OpenFile(fileLocationpengaduan3, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepengaduan3.Close()

			// Copy
			if _, err = io.Copy(targetFilepengaduan3, srcpengaduan3); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if filePengaduan4 == 0 {
			srcpengaduan4, err := filepengaduan4.Open()
			if err != nil {
				return err
			}
			defer srcpengaduan4.Close()

			newfilenamepengaduan4 := helpers.GenerateFilename("PENGADUAN-4" + *result.Nik)
			filenamepengaduan4 = fmt.Sprintf("%s%s", newfilenamepengaduan4, filepath.Ext(filepengaduan4.Filename))

			fileLocationpengaduan4 = filepath.Join(dir, "uploads/aduan", filenamepengaduan4)
			targetFilepengaduan4, err := os.OpenFile(fileLocationpengaduan4, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepengaduan4.Close()

			// Copy
			if _, err = io.Copy(targetFilepengaduan4, srcpengaduan4); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if filePengaduan5 == 0 {
			srcpengaduan5, err := filepengaduan5.Open()
			if err != nil {
				return err
			}
			defer srcpengaduan5.Close()

			newfilenamepengaduan5 := helpers.GenerateFilename("PENGADUAN-5" + *result.Nik)
			filenamepengaduan5 = fmt.Sprintf("%s%s", newfilenamepengaduan5, filepath.Ext(filepengaduan5.Filename))

			fileLocationpengaduan5 = filepath.Join(dir, "uploads/aduan", filenamepengaduan5)
			targetFilepengaduan5, err := os.OpenFile(fileLocationpengaduan5, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepengaduan5.Close()

			// Copy
			if _, err = io.Copy(targetFilepengaduan5, srcpengaduan5); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

	}

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	kode_permohonan := helpers.GenerateKodePermohonan("ADUAN-" + nik_fix)

	var keperluan_fix string
	if keperluan == "Lainnya" {
		keperluan_fix = keperluan_lain
	} else {
		keperluan_fix = keperluan
	}

	if _, err := models.PostAduan(userId, kode_permohonan, *result.Fullname, *result.Nik, *result.No_hp, *result.Alamat, *result.Kelurahan, *result.Kecamatan, nama_fix, nik_fix, nohp_fix, alamat_fix, kelurahan_fix, kecamatan_fix, keperluan_fix, subject_fix, uraian, 0, currentTime.Format("2006-01-02 15:04:05"), filenamepengaduan1, filenamepengaduan2, filenamepengaduan3, filenamepengaduan4, filenamepengaduan5); err != nil {
		if filenamepengaduan1 != "" {
			el := os.Remove(fileLocationpengaduan1)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE PENGADUAN 1")
			}
		}
		if filenamepengaduan2 != "" {
			el := os.Remove(fileLocationpengaduan2)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE PENGADUAN 2")
			}
		}
		if filenamepengaduan3 != "" {
			el := os.Remove(fileLocationpengaduan3)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE PENGADUAN 3")
			}
		}
		if filenamepengaduan4 != "" {
			el := os.Remove(fileLocationpengaduan4)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE PENGADUAN 4")
			}
		}
		if filenamepengaduan5 != "" {
			el := os.Remove(fileLocationpengaduan5)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE PENGADUAN 5")
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
