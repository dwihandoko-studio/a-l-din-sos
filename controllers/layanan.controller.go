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
	"strconv"
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

			newfilenamelainnya := helpers.GenerateFilename("LAINNYA-" + *result.Nik)
			filenamelainnya = fmt.Sprintf("%s%s", newfilenamelainnya, filepath.Ext(filelainnya.Filename))

			fileLocationlainnya = filepath.Join(dir, "uploads/pbi", filenamelainnya)
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

			newfilenamektp := helpers.GenerateFilename("KTP-" + *result.Nik)
			filenamektp = fmt.Sprintf("%s%s", newfilenamektp, filepath.Ext(filektp.Filename))

			fileLocationktp = filepath.Join(dir, "uploads/pbi", filenamektp)
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

			newfilenamekk := helpers.GenerateFilename("KK-" + *result.Nik)
			filenamekk = fmt.Sprintf("%s%s", newfilenamekk, filepath.Ext(filekk.Filename))

			fileLocationkk = filepath.Join(dir, "uploads/pbi", filenamekk)
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

			fileLocationfotorumah = filepath.Join(dir, "uploads/pbi", filenamefotorumah)
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

			fileLocationsktm = filepath.Join(dir, "uploads/pbi", filenamesktm)
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

func PostLayananLks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	userId := claims.Id

	var fileKtpKetua int
	var fileKtpSekretaris int
	var fileKtpBendahara int
	var fileAktaNotaris int
	var filePengesahanKemenkumham int
	var fileAdrt int
	var fileKeteranganDomisili int
	var fileAkreditasi int
	var fileStrukturOrganisasi int
	var fileNpwp int
	var fileFotoLokasi int
	var fileFotoUsahaEkonomi int
	var fileLogoLembaga int
	var fileDataBinaan int

	// var keperluan string
	// var keperluan_lain string

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	// 	fmt.Println("PARSING REQUEST FORM")
	nama_lembaga := c.FormValue("nama_lembaga")
	jenis_lembaga := c.FormValue("jenis_lembaga")
	tanggal_berdiri := c.FormValue("tanggal_berdiri")
	nama_notaris := c.FormValue("nama_notaris")
	no_akta_notaris := c.FormValue("no_akta_notaris")
	no_kemenkumham := c.FormValue("no_kemenkumham")
	akreditasi_lembaga := c.FormValue("akreditasi_lembaga")
	no_akreditasi_lembaga := c.FormValue("no_akreditasi_lembaga")
	tgl_habis_akreditasi_lembaga := c.FormValue("tgl_habis_akreditasi_lembaga")
	npwp_lembaga := c.FormValue("npwp_lembaga")
	modal_usaha := c.FormValue("modal_usaha")
	status_lembaga := c.FormValue("status_lembaga")
	lingkup_wilayah_lembaga := c.FormValue("lingkup_wilayah_lembaga")
	bidang_kegiatan := c.FormValue("bidang_kegiatan")
	notelp_lembaga := c.FormValue("notelp_lembaga")
	email_lembaga := c.FormValue("email_lembaga")
	alamat_lembaga := c.FormValue("alamat_lembaga")
	rt_lembaga := c.FormValue("rt_lembaga")
	rw_lembaga := c.FormValue("rw_lembaga")
	kabupaten_lembaga := c.FormValue("kabupaten_lembaga")
	kecamatan_lembaga := c.FormValue("kecamatan_lembaga")
	kelurahan_lembaga := c.FormValue("kelurahan_lembaga")
	latitude_lembaga := c.FormValue("latitude_lembaga")
	longitude_lembaga := c.FormValue("longitude_lembaga")
	nama_pengurus_ketua := c.FormValue("nama_pengurus_ketua")
	nik_pengurus_ketua := c.FormValue("nik_pengurus_ketua")
	nohp_pengurus_ketua := c.FormValue("nohp_pengurus_ketua")
	nama_pengurus_sekretaris := c.FormValue("nama_pengurus_sekretaris")
	nik_pengurus_sekretaris := c.FormValue("nik_pengurus_sekretaris")
	nohp_pengurus_sekretaris := c.FormValue("nohp_pengurus_sekretaris")
	nama_pengurus_bendahara := c.FormValue("nama_pengurus_bendahara")
	nik_pengurus_bendahara := c.FormValue("nik_pengurus_bendahara")
	nohp_pengurus_bendahara := c.FormValue("nohp_pengurus_bendahara")
	jumlah_pengurus := c.FormValue("jumlah_pengurus")
	jumlah_binaan_dalam := c.FormValue("jumlah_binaan_dalam")
	jumlah_binaan_luar := c.FormValue("jumlah_binaan_luar")
	// } else {
	// 	//json_map has the JSON Payload decoded into a map
	// 	fmt.Println("PARSING REQUEST JSON")
	// 	keperluan = fmt.Sprintf("%s", json_map["keperluan"])
	// 	keperluan_lain = fmt.Sprintf("%s", json_map["keperluan_lain"])
	// }

	if nama_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nama lembaga tidak boleh kosong."},
		)
	}

	if jenis_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "jenis lembaga tidak boleh kosong."},
		)
	}

	if tanggal_berdiri == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "tanggal berdiri lembaga tidak boleh kosong."},
		)
	}

	if nama_notaris == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nama notaris tidak boleh kosong."},
		)
	}

	if no_akta_notaris == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "no akta notaris tidak boleh kosong."},
		)
	}

	if no_kemenkumham == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "no registrasi / pengesahan kemenkumham tidak boleh kosong."},
		)
	}

	if akreditasi_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "akreditasi lembaga tidak boleh kosong."},
		)
	}

	if akreditasi_lembaga != "Belum Terakreditasi" {
		if no_akreditasi_lembaga == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "no akreditasi lembaga tidak boleh kosong."},
			)
		}
		if tgl_habis_akreditasi_lembaga == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "tanggal habis akreditasi lembaga tidak boleh kosong."},
			)
		}
	}

	if npwp_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "npwp lembaga tidak boleh kosong."},
		)
	}

	if modal_usaha == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "modal usaha lembaga tidak boleh kosong."},
		)
	}

	if status_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "status lembaga tidak boleh kosong."},
		)
	}

	if lingkup_wilayah_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "lingkup wilayah kerja lembaga tidak boleh kosong."},
		)
	}

	if bidang_kegiatan == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "bidang kegiatan lembaga tidak boleh kosong."},
		)
	}

	if notelp_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "no telp lembaga tidak boleh kosong."},
		)
	}

	if email_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "email lembaga tidak boleh kosong."},
		)
	}

	if alamat_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "alamat lembaga tidak boleh kosong."},
		)
	}

	if rt_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "rt lembaga tidak boleh kosong."},
		)
	}

	if rw_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "rw lembaga tidak boleh kosong."},
		)
	}

	if kabupaten_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "kabupaten lembaga tidak boleh kosong."},
		)
	}

	if kecamatan_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "kecamatan lembaga tidak boleh kosong."},
		)
	}

	if kelurahan_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "kelurahan lembaga tidak boleh kosong."},
		)
	}

	if latitude_lembaga == "" || longitude_lembaga == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "koordinat lembaga tidak boleh kosong."},
		)
	}

	if nama_pengurus_ketua == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nama pengurus ketua lembaga tidak boleh kosong."},
		)
	}

	if nik_pengurus_ketua == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nik pengurus ketua lembaga tidak boleh kosong."},
		)
	}

	if nohp_pengurus_ketua == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nohp pengurus ketua lembaga tidak boleh kosong."},
		)
	}

	if nama_pengurus_sekretaris == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nama pengurus sekretaris lembaga tidak boleh kosong."},
		)
	}

	if nik_pengurus_sekretaris == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nik pengurus sekretaris lembaga tidak boleh kosong."},
		)
	}

	if nohp_pengurus_sekretaris == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nohp pengurus sekretaris lembaga tidak boleh kosong."},
		)
	}

	if nama_pengurus_bendahara == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nama pengurus bendahara lembaga tidak boleh kosong."},
		)
	}

	if nik_pengurus_bendahara == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nik pengurus bendahara lembaga tidak boleh kosong."},
		)
	}

	if nohp_pengurus_bendahara == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "nohp pengurus bendahara lembaga tidak boleh kosong."},
		)
	}

	if jumlah_pengurus == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "jumlah pengurus lembaga tidak boleh kosong."},
		)
	}

	if jumlah_binaan_dalam == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "jumlah binaan dalam lembaga tidak boleh kosong."},
		)
	}

	if jumlah_binaan_luar == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "jumlah binaan luar lembaga tidak boleh kosong."},
		)
	}

	filektp_ketua, err := c.FormFile("lampiran_ktp_ketua")
	if err != nil {
		fmt.Println(err.Error())
		fileKtpKetua = 1
	} else {

	}

	filektp_sekretaris, err := c.FormFile("lampiran_ktp_sekretaris")
	if err != nil {
		fmt.Println(err.Error())
		fileKtpSekretaris = 1
	} else {

	}

	filektp_bendahara, err := c.FormFile("lampiran_ktp_bendahara")
	if err != nil {
		fmt.Println(err.Error())
		fileKtpBendahara = 1
	} else {

	}

	fileakta_notaris, err := c.FormFile("lampiran_akta_notaris")
	if err != nil {
		fmt.Println(err.Error())
		fileAktaNotaris = 1
	} else {

	}

	filepengesahan_kemenkumham, err := c.FormFile("lampiran_pengesahan_kemenkumham")
	if err != nil {
		fmt.Println(err.Error())
		filePengesahanKemenkumham = 1
	} else {

	}

	fileadrt, err := c.FormFile("lampiran_adrt")
	if err != nil {
		fmt.Println(err.Error())
		fileAdrt = 1
	} else {

	}

	fileketerangan_domisili, err := c.FormFile("lampiran_keterangan_domisili")
	if err != nil {
		fmt.Println(err.Error())
		fileKeteranganDomisili = 1
	} else {

	}

	fileakreditasi, err := c.FormFile("lampiran_akreditasi")
	if err != nil {
		fmt.Println(err.Error())
		fileAkreditasi = 1
	} else {

	}

	filestruktur_organisasi, err := c.FormFile("lampiran_struktur_organisasi")
	if err != nil {
		fmt.Println(err.Error())
		fileStrukturOrganisasi = 1
	} else {

	}

	filenpwp, err := c.FormFile("lampiran_npwp")
	if err != nil {
		fmt.Println(err.Error())
		fileNpwp = 1
	} else {

	}

	filefoto_lokasi, err := c.FormFile("lampiran_foto_lokasi")
	if err != nil {
		fmt.Println(err.Error())
		fileFotoLokasi = 1
	} else {

	}

	filefoto_usaha_ekonomi, err := c.FormFile("lampiran_foto_usaha_ekonomi")
	if err != nil {
		fmt.Println(err.Error())
		fileFotoUsahaEkonomi = 1
	} else {

	}

	filelogo_lembaga, err := c.FormFile("lampiran_logo_lembaga")
	if err != nil {
		fmt.Println(err.Error())
		fileLogoLembaga = 1
	} else {

	}

	filedata_binaan, err := c.FormFile("lampiran_data_binaan")
	if err != nil {
		fmt.Println(err.Error())
		fileDataBinaan = 1
	} else {

	}

	if fileKtpKetua == 0 {
		if string(filepath.Ext(filektp_ketua.Filename)) == ".jpg" ||
			string(filepath.Ext(filektp_ketua.Filename)) == ".pdf" ||
			string(filepath.Ext(filektp_ketua.Filename)) == ".png" ||
			string(filepath.Ext(filektp_ketua.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file ktp ketua tidak diizinkan"},
			)
		}
	}

	if fileKtpSekretaris == 0 {
		if string(filepath.Ext(filektp_sekretaris.Filename)) == ".jpg" ||
			string(filepath.Ext(filektp_sekretaris.Filename)) == ".pdf" ||
			string(filepath.Ext(filektp_sekretaris.Filename)) == ".png" ||
			string(filepath.Ext(filektp_sekretaris.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file ktp sekretaris tidak diizinkan"},
			)
		}
	}

	if fileKtpBendahara == 0 {
		if string(filepath.Ext(filektp_bendahara.Filename)) == ".jpg" ||
			string(filepath.Ext(filektp_bendahara.Filename)) == ".pdf" ||
			string(filepath.Ext(filektp_bendahara.Filename)) == ".png" ||
			string(filepath.Ext(filektp_bendahara.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file ktp bendahara tidak diizinkan"},
			)
		}
	}

	if fileAktaNotaris == 0 {
		if string(filepath.Ext(fileakta_notaris.Filename)) == ".jpg" ||
			string(filepath.Ext(fileakta_notaris.Filename)) == ".pdf" ||
			string(filepath.Ext(fileakta_notaris.Filename)) == ".png" ||
			string(filepath.Ext(fileakta_notaris.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file akta notaris tidak diizinkan"},
			)
		}
	}

	if filePengesahanKemenkumham == 0 {
		if string(filepath.Ext(filepengesahan_kemenkumham.Filename)) == ".jpg" ||
			string(filepath.Ext(filepengesahan_kemenkumham.Filename)) == ".pdf" ||
			string(filepath.Ext(filepengesahan_kemenkumham.Filename)) == ".png" ||
			string(filepath.Ext(filepengesahan_kemenkumham.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file pengesahan/registrasi kemenkumham tidak diizinkan"},
			)
		}
	}

	if fileAdrt == 0 {
		if string(filepath.Ext(fileadrt.Filename)) == ".jpg" ||
			string(filepath.Ext(fileadrt.Filename)) == ".pdf" ||
			string(filepath.Ext(fileadrt.Filename)) == ".png" ||
			string(filepath.Ext(fileadrt.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file adrt tidak diizinkan"},
			)
		}
	}

	if fileKeteranganDomisili == 0 {
		if string(filepath.Ext(fileketerangan_domisili.Filename)) == ".jpg" ||
			string(filepath.Ext(fileketerangan_domisili.Filename)) == ".pdf" ||
			string(filepath.Ext(fileketerangan_domisili.Filename)) == ".png" ||
			string(filepath.Ext(fileketerangan_domisili.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file keterangan domisili tidak diizinkan"},
			)
		}
	}

	if fileAkreditasi == 0 {
		if string(filepath.Ext(fileakreditasi.Filename)) == ".jpg" ||
			string(filepath.Ext(fileakreditasi.Filename)) == ".pdf" ||
			string(filepath.Ext(fileakreditasi.Filename)) == ".png" ||
			string(filepath.Ext(fileakreditasi.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file akreditasi tidak diizinkan"},
			)
		}
	}

	if fileStrukturOrganisasi == 0 {
		if string(filepath.Ext(filestruktur_organisasi.Filename)) == ".jpg" ||
			string(filepath.Ext(filestruktur_organisasi.Filename)) == ".pdf" ||
			string(filepath.Ext(filestruktur_organisasi.Filename)) == ".png" ||
			string(filepath.Ext(filestruktur_organisasi.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file struktur organisasi tidak diizinkan"},
			)
		}
	}

	if fileNpwp == 0 {
		if string(filepath.Ext(filenpwp.Filename)) == ".jpg" ||
			string(filepath.Ext(filenpwp.Filename)) == ".pdf" ||
			string(filepath.Ext(filenpwp.Filename)) == ".png" ||
			string(filepath.Ext(filenpwp.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file npwp tidak diizinkan"},
			)
		}
	}

	if fileFotoLokasi == 0 {
		if string(filepath.Ext(filefoto_lokasi.Filename)) == ".jpg" ||
			string(filepath.Ext(filefoto_lokasi.Filename)) == ".pdf" ||
			string(filepath.Ext(filefoto_lokasi.Filename)) == ".png" ||
			string(filepath.Ext(filefoto_lokasi.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file foto lokasi tidak diizinkan"},
			)
		}
	}

	if fileFotoUsahaEkonomi == 0 {
		if string(filepath.Ext(filefoto_usaha_ekonomi.Filename)) == ".jpg" ||
			string(filepath.Ext(filefoto_usaha_ekonomi.Filename)) == ".pdf" ||
			string(filepath.Ext(filefoto_usaha_ekonomi.Filename)) == ".png" ||
			string(filepath.Ext(filefoto_usaha_ekonomi.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file foto usaha ekonomi produktif tidak diizinkan"},
			)
		}
	}

	if fileLogoLembaga == 0 {
		if string(filepath.Ext(filelogo_lembaga.Filename)) == ".jpg" ||
			string(filepath.Ext(filelogo_lembaga.Filename)) == ".pdf" ||
			string(filepath.Ext(filelogo_lembaga.Filename)) == ".png" ||
			string(filepath.Ext(filelogo_lembaga.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file logo lembaga tidak diizinkan"},
			)
		}
	}

	if fileDataBinaan == 0 {
		if string(filepath.Ext(filedata_binaan.Filename)) == ".jpg" ||
			string(filepath.Ext(filedata_binaan.Filename)) == ".pdf" ||
			string(filepath.Ext(filedata_binaan.Filename)) == ".png" ||
			string(filepath.Ext(filedata_binaan.Filename)) == ".jpeg" {

		} else {
			return c.JSON(http.StatusOK,
				models.Response{Status: 400, Message: "Type file data binaan tidak diizinkan"},
			)
		}
	}

	result, err := models.GetUserDetail(userId)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: err.Error()},
		)
	}

	filenamektp_ketua := ""
	var fileLocationktp_ketua string

	filenamektp_sekretaris := ""
	var fileLocationktp_sekretaris string

	filenamektp_bendahara := ""
	var fileLocationktp_bendahara string

	filenameakta_notaris := ""
	var fileLocationakta_notaris string

	filenamepengesahan_kemenkumham := ""
	var fileLocationpengesahan_kemenkumham string

	filenameadrt := ""
	var fileLocationadrt string

	filenameketerangan_domisili := ""
	var fileLocationketerangan_domisili string

	filenameakreditasi := ""
	var fileLocationakreditasi string

	filenamestruktur_organisasi := ""
	var fileLocationstruktur_organisasi string

	filenamenpwp := ""
	var fileLocationnpwp string

	filenamefoto_lokasi := ""
	var fileLocationfoto_lokasi string

	filenamefoto_usaha_ekonomi := ""
	var fileLocationfoto_usaha_ekonomi string

	filenamelogo_lembaga := ""
	var fileLocationlogo_lembaga string

	filenamedata_binaan := ""
	var fileLocationdata_binaan string

	if fileKtpKetua == 0 || fileKtpSekretaris == 0 || fileKtpBendahara == 0 || fileAktaNotaris == 0 || filePengesahanKemenkumham == 0 || fileAdrt == 0 || fileKeteranganDomisili == 0 || fileStrukturOrganisasi == 0 || fileNpwp == 0 || fileFotoLokasi == 0 || fileFotoUsahaEkonomi == 0 || fileLogoLembaga == 0 || fileDataBinaan == 0 {

		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusOK,
				models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
			)
		}

		if akreditasi_lembaga != "Belum Terakreditasi" {
			if fileAkreditasi == 0 {
				srcakreditasi, err := fileakreditasi.Open()
				if err != nil {
					fmt.Println(err.Error())
					return c.JSON(http.StatusOK,
						models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
					)
					// return err
				}
				defer srcakreditasi.Close()

				newfilenameakreditasi := helpers.GenerateFilename("AKREDITASI" + *result.Nik)
				filenameakreditasi = fmt.Sprintf("%s%s", newfilenameakreditasi, filepath.Ext(fileakreditasi.Filename))

				fileLocationakreditasi = filepath.Join(dir, "uploads/lks", filenameakreditasi)
				targetFileakreditasi, err := os.OpenFile(fileLocationakreditasi, os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					fmt.Println(err.Error())
					return c.JSON(http.StatusOK,
						models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
					)
				}
				defer targetFileakreditasi.Close()

				// Copy
				if _, err = io.Copy(targetFileakreditasi, srcakreditasi); err != nil {
					fmt.Println(err.Error())
					return c.JSON(http.StatusOK,
						models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
					)
				}
			}

		} else {

		}

		if fileKtpKetua == 0 {
			srcktp_ketua, err := filektp_ketua.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcktp_ketua.Close()

			newfilenamektp_ketua := helpers.GenerateFilename("KTP-KETUA-" + *result.Nik)
			filenamektp_ketua = fmt.Sprintf("%s%s", newfilenamektp_ketua, filepath.Ext(filektp_ketua.Filename))

			fileLocationktp_ketua = filepath.Join(dir, "uploads/lks", filenamektp_ketua)
			targetFilektp_ketua, err := os.OpenFile(fileLocationktp_ketua, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilektp_ketua.Close()

			// Copy
			if _, err = io.Copy(targetFilektp_ketua, srcktp_ketua); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKtpSekretaris == 0 {
			srcktp_sekretaris, err := filektp_sekretaris.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcktp_sekretaris.Close()

			newfilenamektp_sekretaris := helpers.GenerateFilename("KTP-SEKRETARIS-" + *result.Nik)
			filenamektp_sekretaris = fmt.Sprintf("%s%s", newfilenamektp_sekretaris, filepath.Ext(filektp_sekretaris.Filename))

			fileLocationktp_sekretaris = filepath.Join(dir, "uploads/lks", filenamektp_sekretaris)
			targetFilektp_sekretaris, err := os.OpenFile(fileLocationktp_sekretaris, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilektp_sekretaris.Close()

			// Copy
			if _, err = io.Copy(targetFilektp_sekretaris, srcktp_sekretaris); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKtpBendahara == 0 {
			srcktp_bendahara, err := filektp_bendahara.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcktp_bendahara.Close()

			newfilenamektp_bendahara := helpers.GenerateFilename("KTP-BENDAHARA-" + *result.Nik)
			filenamektp_bendahara = fmt.Sprintf("%s%s", newfilenamektp_bendahara, filepath.Ext(filektp_bendahara.Filename))

			fileLocationktp_bendahara = filepath.Join(dir, "uploads/lks", filenamektp_bendahara)
			targetFilektp_bendahara, err := os.OpenFile(fileLocationktp_bendahara, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilektp_bendahara.Close()

			// Copy
			if _, err = io.Copy(targetFilektp_bendahara, srcktp_bendahara); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileAktaNotaris == 0 {
			srcakta_notaris, err := fileakta_notaris.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcakta_notaris.Close()

			newfilenameakta_notaris := helpers.GenerateFilename("AKTA-NOTARIS-" + *result.Nik)
			filenameakta_notaris = fmt.Sprintf("%s%s", newfilenameakta_notaris, filepath.Ext(fileakta_notaris.Filename))

			fileLocationakta_notaris = filepath.Join(dir, "uploads/lks", filenameakta_notaris)
			targetFileakta_notaris, err := os.OpenFile(fileLocationakta_notaris, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFileakta_notaris.Close()

			// Copy
			if _, err = io.Copy(targetFileakta_notaris, srcakta_notaris); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if filePengesahanKemenkumham == 0 {
			srcpengesahan_kemenkumham, err := filepengesahan_kemenkumham.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcpengesahan_kemenkumham.Close()

			newfilenamepengesahan_kemenkumham := helpers.GenerateFilename("PENGESAHAN-KEMENKUMHAM-" + *result.Nik)
			filenamepengesahan_kemenkumham = fmt.Sprintf("%s%s", newfilenamepengesahan_kemenkumham, filepath.Ext(filepengesahan_kemenkumham.Filename))

			fileLocationpengesahan_kemenkumham = filepath.Join(dir, "uploads/lks", filenamepengesahan_kemenkumham)
			targetFilepengesahan_kemenkumham, err := os.OpenFile(fileLocationpengesahan_kemenkumham, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepengesahan_kemenkumham.Close()

			// Copy
			if _, err = io.Copy(targetFilepengesahan_kemenkumham, srcpengesahan_kemenkumham); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileAdrt == 0 {
			srcadrt, err := fileadrt.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcadrt.Close()

			newfilenameadrt := helpers.GenerateFilename("ADRT" + *result.Nik)
			filenameadrt = fmt.Sprintf("%s%s", newfilenameadrt, filepath.Ext(fileadrt.Filename))

			fileLocationadrt = filepath.Join(dir, "uploads/lks", filenameadrt)
			targetFileadrt, err := os.OpenFile(fileLocationadrt, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFileadrt.Close()

			// Copy
			if _, err = io.Copy(targetFileadrt, srcadrt); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileKeteranganDomisili == 0 {
			srcketerangan_domisili, err := fileketerangan_domisili.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcketerangan_domisili.Close()

			newfilenameketerangan_domisili := helpers.GenerateFilename("KETERANGAN-DOMISILI-" + *result.Nik)
			filenameketerangan_domisili = fmt.Sprintf("%s%s", newfilenameketerangan_domisili, filepath.Ext(fileketerangan_domisili.Filename))

			fileLocationketerangan_domisili = filepath.Join(dir, "uploads/lks", filenameketerangan_domisili)
			targetFileketerangan_domisili, err := os.OpenFile(fileLocationketerangan_domisili, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFileketerangan_domisili.Close()

			// Copy
			if _, err = io.Copy(targetFileketerangan_domisili, srcketerangan_domisili); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileStrukturOrganisasi == 0 {
			srcstruktur_organisasi, err := filestruktur_organisasi.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcstruktur_organisasi.Close()

			newfilenamestruktur_organisasi := helpers.GenerateFilename("STRUKTUR-ORGANISASI-" + *result.Nik)
			filenamestruktur_organisasi = fmt.Sprintf("%s%s", newfilenamestruktur_organisasi, filepath.Ext(filestruktur_organisasi.Filename))

			fileLocationstruktur_organisasi = filepath.Join(dir, "uploads/lks", filenamestruktur_organisasi)
			targetFilestruktur_organisasi, err := os.OpenFile(fileLocationstruktur_organisasi, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilestruktur_organisasi.Close()

			// Copy
			if _, err = io.Copy(targetFilestruktur_organisasi, srcstruktur_organisasi); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileNpwp == 0 {
			srcnpwp, err := filenpwp.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcnpwp.Close()

			newfilenamenpwp := helpers.GenerateFilename("NPWP-" + *result.Nik)
			filenamenpwp = fmt.Sprintf("%s%s", newfilenamenpwp, filepath.Ext(filenpwp.Filename))

			fileLocationnpwp = filepath.Join(dir, "uploads/lks", filenamenpwp)
			targetFilenpwp, err := os.OpenFile(fileLocationnpwp, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilenpwp.Close()

			// Copy
			if _, err = io.Copy(targetFilenpwp, srcnpwp); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileFotoLokasi == 0 {
			srcfoto_lokasi, err := filefoto_lokasi.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcfoto_lokasi.Close()

			newfilenamefoto_lokasi := helpers.GenerateFilename("FOTO-LOKASI-" + *result.Nik)
			filenamefoto_lokasi = fmt.Sprintf("%s%s", newfilenamefoto_lokasi, filepath.Ext(filefoto_lokasi.Filename))

			fileLocationfoto_lokasi = filepath.Join(dir, "uploads/lks", filenamefoto_lokasi)
			targetFilefoto_lokasi, err := os.OpenFile(fileLocationfoto_lokasi, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilefoto_lokasi.Close()

			// Copy
			if _, err = io.Copy(targetFilefoto_lokasi, srcfoto_lokasi); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileFotoUsahaEkonomi == 0 {
			srcfoto_usaha_ekonomi, err := filefoto_usaha_ekonomi.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcfoto_usaha_ekonomi.Close()

			newfilenamefoto_usaha_ekonomi := helpers.GenerateFilename("FOTO-USAHA-EKONOMI-" + *result.Nik)
			filenamefoto_usaha_ekonomi = fmt.Sprintf("%s%s", newfilenamefoto_usaha_ekonomi, filepath.Ext(filefoto_usaha_ekonomi.Filename))

			fileLocationfoto_usaha_ekonomi = filepath.Join(dir, "uploads/lks", filenamefoto_usaha_ekonomi)
			targetFilefoto_usaha_ekonomi, err := os.OpenFile(fileLocationfoto_usaha_ekonomi, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilefoto_usaha_ekonomi.Close()

			// Copy
			if _, err = io.Copy(targetFilefoto_usaha_ekonomi, srcfoto_usaha_ekonomi); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileLogoLembaga == 0 {
			srclogo_lembaga, err := filelogo_lembaga.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srclogo_lembaga.Close()

			newfilenamelogo_lembaga := helpers.GenerateFilename("LOGO-LEMBAGA-" + *result.Nik)
			filenamelogo_lembaga = fmt.Sprintf("%s%s", newfilenamelogo_lembaga, filepath.Ext(filelogo_lembaga.Filename))

			fileLocationlogo_lembaga = filepath.Join(dir, "uploads/lks", filenamelogo_lembaga)
			targetFilelogo_lembaga, err := os.OpenFile(fileLocationlogo_lembaga, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilelogo_lembaga.Close()

			// Copy
			if _, err = io.Copy(targetFilelogo_lembaga, srclogo_lembaga); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}

		if fileDataBinaan == 0 {
			srcdata_binaan, err := filedata_binaan.Open()
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam membaca file."},
				)
			}
			defer srcdata_binaan.Close()

			newfilenamedata_binaan := helpers.GenerateFilename("DATA-BINAAN-" + *result.Nik)
			filenamedata_binaan = fmt.Sprintf("%s%s", newfilenamedata_binaan, filepath.Ext(filedata_binaan.Filename))

			fileLocationdata_binaan = filepath.Join(dir, "uploads/lks", filenamedata_binaan)
			targetFiledata_binaan, err := os.OpenFile(fileLocationdata_binaan, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFiledata_binaan.Close()

			// Copy
			if _, err = io.Copy(targetFiledata_binaan, srcdata_binaan); err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
		}
	}

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	kode_permohonan := helpers.GenerateKodePermohonan("LKS-" + *result.Nik)

	//var keperluan_fix string
	//if keperluan == "Lainnya" {
	//keperluan_fix = keperluan_lain
	//} else {
	//	keperluan_fix = keperluan
	//}
	latLong := latitude_lembaga + "," + longitude_lembaga
	jp, errc := strconv.ParseInt(jumlah_pengurus, 10, 64)
	if errc != nil {
		fmt.Println("Error:", errc)
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "gagal convert int."},
		)
	}
	mu, errc := strconv.ParseInt(modal_usaha, 10, 64)
	if errc != nil {
		fmt.Println("Error:", errc)
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "gagal convert int."},
		)
	}
	jbd, errc := strconv.ParseInt(jumlah_binaan_dalam, 10, 64)
	if errc != nil {
		fmt.Println("Error:", errc)
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "gagal convert int."},
		)
	}
	jbl, errc := strconv.ParseInt(jumlah_binaan_luar, 10, 64)
	if errc != nil {
		fmt.Println("Error:", errc)
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "gagal convert int."},
		)
	}

	dataLks := models.LksIns{
		Nama_lembaga:                   &nama_lembaga,
		Jenis_lembaga:                  &jenis_lembaga,
		Tgl_berdiri_lembaga:            &tanggal_berdiri,
		Nama_notaris_lembaga:           &nama_notaris,
		Nomor_notaris_lembaga:          &no_akta_notaris,
		Nomor_Kemenkumham_lembaga:      &no_kemenkumham,
		Akreditasi_lembaga:             &akreditasi_lembaga,
		Nomor_surat_akreditasi_lembaga: &no_akreditasi_lembaga,
		Tgl_expired_akreditasi_lembaga: &tgl_habis_akreditasi_lembaga,
		Npwp_lembaga:                   &npwp_lembaga,
		Modal_usaha_lembaga:            mu,
		Status_lembaga:                 &status_lembaga,
		Lingkup_wilayah_kerja_lembaga:  &lingkup_wilayah_lembaga,
		Bidang_kegiatan_lembaga:        &bidang_kegiatan,
		No_telp_lembaga:                &notelp_lembaga,
		Email_lembaga:                  &email_lembaga,
		Lat_long_lembaga:               &latLong,
		Alamat_lembaga:                 &alamat_lembaga,
		Rt_lembaga:                     &rt_lembaga,
		Rw_lembaga:                     &rw_lembaga,
		Kecamatan_lembaga:              &kecamatan_lembaga,
		Kelurahan_lembaga:              &kelurahan_lembaga,
		Nama_ketua_pengurus:            &nama_pengurus_ketua,
		Nik_ketua_pengurus:             &nik_pengurus_ketua,
		Nohp_ketua_pengurus:            &nohp_pengurus_ketua,
		Nama_sekretaris_pengurus:       &nama_pengurus_sekretaris,
		Nik_sekretaris_pengurus:        &nik_pengurus_sekretaris,
		Nohp_sekretaris_pengurus:       &nohp_pengurus_sekretaris,
		Nama_bendahara_pengurus:        &nama_pengurus_bendahara,
		Nik_bendahara_pengurus:         &nik_pengurus_bendahara,
		Nohp_bendahara_pengurus:        &nohp_pengurus_bendahara,
		Jumlah_pengurus:                jp,
		Jumlah_binaan_dalam:            jbd,
		Jumlah_binaan_luar:             jbl,
	}

	if _, err := models.PostLks(userId, kode_permohonan, *result.Kelurahan, "kadis", *result.Nik, *result.Fullname, "Rekomendasi LKS/LKSA", "LKS", "0", currentTime.Format("2006-01-02 15:04:05"), filenamektp_ketua, filenamektp_sekretaris, filenamektp_bendahara, filenameakta_notaris, filenamepengesahan_kemenkumham, filenameadrt, filenameketerangan_domisili, filenameakreditasi, filenamestruktur_organisasi, filenamenpwp, filenamefoto_lokasi, filenamefoto_usaha_ekonomi, filenamelogo_lembaga, filenamedata_binaan, dataLks); err != nil {
		if filenamektp_ketua != "" {
			el := os.Remove(fileLocationktp_ketua)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KTP KETUA")
			}
		}
		if filenamektp_sekretaris != "" {
			el := os.Remove(fileLocationktp_sekretaris)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KTP sekretaris")
			}
		}
		if filenamektp_bendahara != "" {
			el := os.Remove(fileLocationktp_bendahara)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KTP bendahara")
			}
		}
		if filenameakta_notaris != "" {
			el := os.Remove(fileLocationakta_notaris)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE AKTA NOTARIS")
			}
		}
		if filenamepengesahan_kemenkumham != "" {
			el := os.Remove(fileLocationpengesahan_kemenkumham)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KEMENKUMHAM")
			}
		}
		if filenameadrt != "" {
			el := os.Remove(fileLocationadrt)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE ADRT")
			}
		}
		if filenameketerangan_domisili != "" {
			el := os.Remove(fileLocationketerangan_domisili)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE KETERANGAN DOMISILI")
			}
		}
		if filenameakreditasi != "" {
			el := os.Remove(fileLocationakreditasi)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE AKREDITASI")
			}
		}
		if filenamestruktur_organisasi != "" {
			el := os.Remove(fileLocationstruktur_organisasi)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE STRUKTUR ORGANISASI")
			}
		}
		if filenamenpwp != "" {
			el := os.Remove(fileLocationnpwp)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE NPWP")
			}
		}
		if filenamefoto_lokasi != "" {
			el := os.Remove(fileLocationfoto_lokasi)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE FOTO LOKASI")
			}
		}
		if filenamefoto_usaha_ekonomi != "" {
			el := os.Remove(fileLocationfoto_usaha_ekonomi)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE FOTO USAHA EKONOMI")
			}
		}
		if filenamelogo_lembaga != "" {
			el := os.Remove(fileLocationlogo_lembaga)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE LOGO LEMBAGA")
			}
		}
		if filenamedata_binaan != "" {
			el := os.Remove(fileLocationdata_binaan)
			if el != nil {
				fmt.Println("GAGAL MENGHAPUS FILE DATA BINAAN")
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
	var filePernyataan int

	// var keperluan string
	// var keperluan_lain string

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	// 	fmt.Println("PARSING REQUEST FORM")
	keperluan := c.FormValue("keperluan")
	keperluan_lain := c.FormValue("keperluan_lain")
	indikator1 := c.FormValue("indikator1")
	indikator2 := c.FormValue("indikator2")
	indikator3 := c.FormValue("indikator3")
	indikator4 := c.FormValue("indikator4")
	indikator5 := c.FormValue("indikator5")
	indikator6 := c.FormValue("indikator6")
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

	if indikator1 == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Indikator 1 tidak boleh kosong."},
		)
	}

	if indikator2 == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Indikator 2 tidak boleh kosong."},
		)
	}

	if indikator3 == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Indikator 3 tidak boleh kosong."},
		)
	}

	if indikator4 == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Indikator 4 tidak boleh kosong."},
		)
	}

	if indikator5 == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Indikator 5 tidak boleh kosong."},
		)
	}

	if indikator6 == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Indikator 6 tidak boleh kosong."},
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

	filepernyataan, err := c.FormFile("pernyataan")
	if err != nil {
		fmt.Println(err.Error())
		filePernyataan = 1
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

	if filePernyataan == 0 {
		if string(filepath.Ext(filepernyataan.Filename)) == ".jpg" ||
			string(filepath.Ext(filepernyataan.Filename)) == ".pdf" ||
			string(filepath.Ext(filepernyataan.Filename)) == ".png" ||
			string(filepath.Ext(filepernyataan.Filename)) == ".jpeg" {

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

	filenamepernyataan := ""
	var fileLocationpernyataan string

	if fileFotorumah == 0 || fileKk == 0 || fileKtp == 0 || fileLain == 0 || filePernyataan == 0 {
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

			newfilenamelainnya := helpers.GenerateFilename("LAINNYA-" + *result.Nik)
			filenamelainnya = fmt.Sprintf("%s%s", newfilenamelainnya, filepath.Ext(filelainnya.Filename))

			fileLocationlainnya = filepath.Join(dir, "uploads/sktm", filenamelainnya)
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

			newfilenamektp := helpers.GenerateFilename("KTP-" + *result.Nik)
			filenamektp = fmt.Sprintf("%s%s", newfilenamektp, filepath.Ext(filektp.Filename))

			fileLocationktp = filepath.Join(dir, "uploads/sktm", filenamektp)
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

			newfilenamekk := helpers.GenerateFilename("KK-" + *result.Nik)
			filenamekk = fmt.Sprintf("%s%s", newfilenamekk, filepath.Ext(filekk.Filename))

			fileLocationkk = filepath.Join(dir, "uploads/sktm", filenamekk)
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

			fileLocationfotorumah = filepath.Join(dir, "uploads/sktm", filenamefotorumah)
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

		if filePernyataan == 0 {
			srcpernyataan, err := filepernyataan.Open()
			if err != nil {
				return err
			}
			defer srcpernyataan.Close()

			newfilenamepernyataan := helpers.GenerateFilename("PERNYATAAN-" + *result.Nik)
			filenamepernyataan = fmt.Sprintf("%s%s", newfilenamepernyataan, filepath.Ext(filepernyataan.Filename))

			fileLocationpernyataan = filepath.Join(dir, "uploads/sktm", filenamepernyataan)
			targetFilepernyataan, err := os.OpenFile(fileLocationpernyataan, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err.Error())
				return c.JSON(http.StatusOK,
					models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
				)
			}
			defer targetFilepernyataan.Close()

			// Copy
			if _, err = io.Copy(targetFilepernyataan, srcpernyataan); err != nil {
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

	// indikator_1, _ := strconv.ParseFloat(indikator1, 64)
	indikator_1, _ := strconv.Atoi(indikator1)
	indikator_2, _ := strconv.Atoi(indikator2)
	indikator_3, _ := strconv.Atoi(indikator3)
	indikator_4, _ := strconv.Atoi(indikator4)
	indikator_5, _ := strconv.Atoi(indikator5)
	indikator_6, _ := strconv.Atoi(indikator6)

	// var skor float64
	skor := ((float64(indikator_1) + float64(indikator_2) + float64(indikator_3) + float64(indikator_4) + float64(indikator_5) + float64(indikator_6)) / 16) * 100

	if _, err := models.PostSktm(userId, kode_permohonan, *result.Kelurahan, "kakam", *result.Nik, *result.Fullname, keperluan_fix, "SKTM", "0", currentTime.Format("2006-01-02 15:04:05"), filenamektp, filenamekk, filenamefotorumah, filenamelainnya, filenamepernyataan, indikator_1, indikator_2, indikator_3, indikator_4, indikator_5, indikator_6, skor); err != nil {
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

			newfilenamelainnya := helpers.GenerateFilename("LAINNYA-" + *result.Nik)
			filenamelainnya = fmt.Sprintf("%s%s", newfilenamelainnya, filepath.Ext(filelainnya.Filename))

			fileLocationlainnya = filepath.Join(dir, "uploads/dtks", filenamelainnya)
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

			newfilenamektp := helpers.GenerateFilename("KTP-" + *result.Nik)
			filenamektp = fmt.Sprintf("%s%s", newfilenamektp, filepath.Ext(filektp.Filename))

			fileLocationktp = filepath.Join(dir, "uploads/dtks", filenamektp)
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

			newfilenamekk := helpers.GenerateFilename("KK-" + *result.Nik)
			filenamekk = fmt.Sprintf("%s%s", newfilenamekk, filepath.Ext(filekk.Filename))

			fileLocationkk = filepath.Join(dir, "uploads/dtks", filenamekk)
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

			fileLocationfotorumah = filepath.Join(dir, "uploads/dtks", filenamefotorumah)
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
