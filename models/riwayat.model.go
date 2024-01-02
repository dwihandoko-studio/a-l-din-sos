package models

import (
	"api-dinsos/db"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Riwayat struct {
	Id                *string `json:"id"`
	Layanan           *string `json:"layanan"`
	Kode_permohonan   *string `json:"kode_permohonan"`
	Nik               *string `json:"nik"`
	Nama              *string `json:"nama"`
	User_id           *string `json:"user_id"`
	Jenis             *string `json:"jenis"`
	Status_permohonan int64   `json:"status_permohonan"`
	Created_at        *string `json:"created_at"`
	// Nisn       *string `json:"nisn"`
	// Email      *string `json:"email"`
	// No_hp      *string `json:"no_hp"`
}

type DafRiwayat struct {
	Data      []Riwayat `json:"data"`
	Total     int       `json:"total"`
	Page      int       `json:"page"`
	Last_page int       `json:"last_page"`
}

func GetRiwayatLayanan(queryString string, pageInitial string) (DafRiwayat, error) {
	var obj Riwayat
	var arryobj []Riwayat
	var result DafRiwayat
	// var query string
	perPage := 10
	var total int
	page, _ := strconv.Atoi(pageInitial)

	con := db.CreateCon()
	jumDat, err := con.Query("SELECT COUNT(*) as total FROM v_riwayat_permohonan " + queryString)
	if err != nil {
		// fmt.Println("Error line jumlah: " + err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer jumDat.Close()
	for jumDat.Next() {
		jumDat.Scan(&total)
	}

	if total <= 0 {
		result.Data = arryobj
		result.Total = total
		result.Page = page
		// result.Last_page = int(float64(total / perPage))
		return result, nil
	}

	queryString = "SELECT * FROM v_riwayat_permohonan " + queryString + " ORDER BY created_at DESC"
	queryString = fmt.Sprintf("%s LIMIT %d OFFSET %d", queryString, perPage, (page-1)*perPage)

	respo, err := con.Query(queryString)
	if err != nil {
		// fmt.Println(err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer respo.Close()

	for respo.Next() {
		err = respo.Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
		)
		if err != nil {
			// fmt.Println(err.Error())
			return result, errors.New("gagal mengambil data")
		}

		arryobj = append(arryobj, obj)
	}

	result.Data = arryobj
	result.Total = total
	result.Page = page
	result.Last_page = int(float64(total / perPage))
	// result.Last_page = int(math.Ceil(float64(total / perPage)))

	return result, nil
}

type Riwayatpengaduan struct {
	Id           *string `json:"id"`
	User_id      *string `json:"user_id"`
	Kode_aduan   *string `json:"kode_aduan"`
	Kategori     *string `json:"kategori"`
	Status_aduan int64   `json:"status_aduan"`
	Created_at   *string `json:"created_at"`
}

type DafRiwayatpengaduan struct {
	Data      []Riwayatpengaduan `json:"data"`
	Total     int                `json:"total"`
	Page      int                `json:"page"`
	Last_page int                `json:"last_page"`
}

func GetRiwayatPengaduan(queryString string, pageInitial string) (DafRiwayatpengaduan, error) {
	var obj Riwayatpengaduan
	var arryobj []Riwayatpengaduan
	var result DafRiwayatpengaduan
	// var query string
	perPage := 10
	var total int
	page, _ := strconv.Atoi(pageInitial)

	con := db.CreateCon()
	jumDat, err := con.Query("SELECT COUNT(*) as total FROM v_riwayat_pengaduan " + queryString)
	if err != nil {
		// fmt.Println("Error line jumlah: " + err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer jumDat.Close()
	for jumDat.Next() {
		jumDat.Scan(&total)
	}

	if total <= 0 {
		result.Data = arryobj
		result.Total = total
		result.Page = page
		// result.Last_page = int(float64(total / perPage))
		return result, nil
	}

	queryString = "SELECT * FROM v_riwayat_pengaduan " + queryString + " ORDER BY created_at DESC"
	queryString = fmt.Sprintf("%s LIMIT %d OFFSET %d", queryString, perPage, (page-1)*perPage)

	respo, err := con.Query(queryString)
	if err != nil {
		// fmt.Println(err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer respo.Close()

	for respo.Next() {
		err = respo.Scan(
			&obj.Id,
			&obj.User_id,
			&obj.Kode_aduan,
			&obj.Kategori,
			&obj.Status_aduan,
			&obj.Created_at,
		)
		if err != nil {
			// fmt.Println(err.Error())
			return result, errors.New("gagal mengambil data")
		}

		arryobj = append(arryobj, obj)
	}

	result.Data = arryobj
	result.Total = total
	result.Page = page
	result.Last_page = int(float64(total / perPage))
	// result.Last_page = int(math.Ceil(float64(total / perPage)))

	return result, nil
}

type DetailAduan struct {
	Id                  *string `json:"id"`
	User_id             *string `json:"user_id"`
	Kode_aduan          *string `json:"kode_aduan"`
	Nama                *string `json:"nama"`
	Nik                 *string `json:"nik"`
	Nohp                *string `json:"nohp"`
	Alamat              *string `json:"alamat"`
	Kecamatan           *string `json:"kecamatan"`
	Kelurahan           *string `json:"kelurahan"`
	Nama_aduan          *string `json:"nama_aduan"`
	Nik_aduan           *string `json:"nik_aduan"`
	Nohp_aduan          *string `json:"nohp_aduan"`
	Alamat_aduan        *string `json:"alamat_aduan"`
	Kecamatan_aduan     *string `json:"kecamatan_aduan"`
	Kelurahan_aduan     *string `json:"kelurahan_aduan"`
	Kategori            *string `json:"kategori"`
	Identitas_aduan     *string `json:"identitas_aduan"`
	Uraian_aduan        *string `json:"uraian_aduan"`
	Media_pengaduan     *string `json:"media_pengaduan"`
	Lampiran_aduan_1    *string `json:"lampiran_aduan_1"`
	Lampiran_aduan_2    *string `json:"lampiran_aduan_2"`
	Lampiran_aduan_3    *string `json:"lampiran_aduan_3"`
	Lampiran_aduan_4    *string `json:"lampiran_aduan_4"`
	Lampiran_aduan_5    *string `json:"lampiran_aduan_5"`
	Permasalahan        *string `json:"permasalahan"`
	Status_aduan        int64   `json:"status_aduan"`
	Diteruskan_ke       *string `json:"diteruskan_ke"`
	Penyelesaian        *string `json:"penyelesaian"`
	Ke_instansi_terkait *string `json:"ke_instansi_terkait"`
	Admin_proses        *string `json:"admin_proses"`
	Date_proses         *string `json:"date_proses"`
	Admin_approve       *string `json:"admin_approve"`
	Date_approve        *string `json:"date_approve"`
	Admin_assesment     *string `json:"admin_assesment"`
	Date_assesment      *string `json:"date_assesment"`
	Created_at          *string `json:"created_at"`
}

func GetDetailPengaduan(id, status string) (DetailAduan, error) {
	var obj DetailAduan

	con := db.CreateCon()

	// if status == "0" {
	sqlStatement := "SELECT id, user_id, kode_aduan, nama, nik, nohp, alamat, kecamatan, kelurahan, nama_aduan, nik_aduan, nohp_aduan, alamat_aduan, kecamatan_aduan, kelurahan_aduan, kategori, identitas_aduan, uraian_aduan, media_pengaduan, lampiran_aduan_1, lampiran_aduan_2, lampiran_aduan_3, lampiran_aduan_4, lampiran_aduan_5, permasalahan, status_aduan, diteruskan_ke, penyelesaian, ke_instansi_terkait, admin_proses, date_proses, admin_approve, date_approve, admin_assesment, date_assesment, created_at From _pengaduan WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(
		&obj.Id,
		&obj.User_id,
		&obj.Kode_aduan,
		&obj.Nama,
		&obj.Nik,
		&obj.Nohp,
		&obj.Alamat,
		&obj.Kecamatan,
		&obj.Kelurahan,
		&obj.Nama_aduan,
		&obj.Nik_aduan,
		&obj.Nohp_aduan,
		&obj.Alamat_aduan,
		&obj.Kecamatan_aduan,
		&obj.Kelurahan_aduan,
		&obj.Kategori,
		&obj.Identitas_aduan,
		&obj.Uraian_aduan,
		&obj.Media_pengaduan,
		&obj.Lampiran_aduan_1,
		&obj.Lampiran_aduan_2,
		&obj.Lampiran_aduan_3,
		&obj.Lampiran_aduan_4,
		&obj.Lampiran_aduan_5,
		&obj.Permasalahan,
		&obj.Status_aduan,
		&obj.Diteruskan_ke,
		&obj.Penyelesaian,
		&obj.Ke_instansi_terkait,
		&obj.Admin_proses,
		&obj.Date_proses,
		&obj.Admin_approve,
		&obj.Date_approve,
		&obj.Admin_assesment,
		&obj.Date_assesment,
		&obj.Created_at,
	)

	if err == sql.ErrNoRows {
		// fmt.Println("User not found")
		return obj, errors.New("data tidak ditemukan")
	}

	if err != nil {
		// fmt.Println("Query error")
		log.Println(err.Error())
		return obj, errors.New("gagal mengambil data")
	}
	// } else if status == "1" {
	// 	sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

	// 	err := con.QueryRow(sqlStatement, id).Scan(
	// 		&obj.Id,
	// 		&obj.Layanan,
	// 		&obj.Kode_permohonan,
	// 		&obj.Kelurahan,
	// 		&obj.Ttd,
	// 		&obj.Nik,
	// 		&obj.Nama,
	// 		&obj.User_id,
	// 		&obj.Jenis,
	// 		&obj.Indikator1,
	// 		&obj.Indikator2,
	// 		&obj.Indikator3,
	// 		&obj.Indikator4,
	// 		&obj.Indikator5,
	// 		&obj.Indikator6,
	// 		&obj.Skor,
	// 		&obj.Lampiran_ktp,
	// 		&obj.Lampiran_kk,
	// 		&obj.Lampiran_pernyataan,
	// 		&obj.Lampiran_foto_rumah,
	// 		&obj.Lampiran_lainnya,
	// 		&obj.Admin_approve,
	// 		&obj.Date_approve,
	// 		&obj.Status_permohonan,
	// 		&obj.Created_at,
	// 		&obj.Updated_at,
	// 	)

	// 	if err == sql.ErrNoRows {
	// 		// fmt.Println("User not found")
	// 		return obj, errors.New("data tidak ditemukan")
	// 	}

	// 	if err != nil {
	// 		// fmt.Println("Query error")
	// 		log.Println(err.Error())
	// 		return obj, errors.New("gagal mengambil data")
	// 	}
	// } else if status == "2" {
	// 	sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

	// 	err := con.QueryRow(sqlStatement, id).Scan(
	// 		&obj.Id,
	// 		&obj.Layanan,
	// 		&obj.Kode_permohonan,
	// 		&obj.Kelurahan,
	// 		&obj.Ttd,
	// 		&obj.Nik,
	// 		&obj.Nama,
	// 		&obj.User_id,
	// 		&obj.Jenis,
	// 		&obj.Indikator1,
	// 		&obj.Indikator2,
	// 		&obj.Indikator3,
	// 		&obj.Indikator4,
	// 		&obj.Indikator5,
	// 		&obj.Indikator6,
	// 		&obj.Skor,
	// 		&obj.Lampiran_ktp,
	// 		&obj.Lampiran_kk,
	// 		&obj.Lampiran_pernyataan,
	// 		&obj.Lampiran_foto_rumah,
	// 		&obj.Lampiran_lainnya,
	// 		&obj.Admin_approve,
	// 		&obj.Date_approve,
	// 		&obj.Status_permohonan,
	// 		&obj.Created_at,
	// 		&obj.Updated_at,
	// 	)

	// 	if err == sql.ErrNoRows {
	// 		// fmt.Println("User not found")
	// 		return obj, errors.New("data tidak ditemukan")
	// 	}

	// 	if err != nil {
	// 		// fmt.Println("Query error")
	// 		log.Println(err.Error())
	// 		return obj, errors.New("gagal mengambil data")
	// 	}
	// } else if status == "3" {
	// 	sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan_tolak WHERE id = ?"

	// 	err := con.QueryRow(sqlStatement, id).Scan(
	// 		&obj.Id,
	// 		&obj.Layanan,
	// 		&obj.Kode_permohonan,
	// 		&obj.Kelurahan,
	// 		&obj.Ttd,
	// 		&obj.Nik,
	// 		&obj.Nama,
	// 		&obj.User_id,
	// 		&obj.Jenis,
	// 		&obj.Indikator1,
	// 		&obj.Indikator2,
	// 		&obj.Indikator3,
	// 		&obj.Indikator4,
	// 		&obj.Indikator5,
	// 		&obj.Indikator6,
	// 		&obj.Skor,
	// 		&obj.Lampiran_ktp,
	// 		&obj.Lampiran_kk,
	// 		&obj.Lampiran_pernyataan,
	// 		&obj.Lampiran_foto_rumah,
	// 		&obj.Lampiran_lainnya,
	// 		&obj.Admin_reject,
	// 		&obj.Date_reject,
	// 		&obj.Keterangan_reject,
	// 		&obj.Status_permohonan,
	// 		&obj.Created_at,
	// 		&obj.Updated_at,
	// 	)

	// 	if err == sql.ErrNoRows {
	// 		// fmt.Println("User not found")
	// 		return obj, errors.New("data tidak ditemukan")
	// 	}

	// 	if err != nil {
	// 		// fmt.Println("Query error")
	// 		log.Println(err.Error())
	// 		return obj, errors.New("gagal mengambil data")
	// 	}
	// } else if status == "4" {
	// 	sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

	// 	err := con.QueryRow(sqlStatement, id).Scan(
	// 		&obj.Id,
	// 		&obj.Layanan,
	// 		&obj.Kode_permohonan,
	// 		&obj.Kelurahan,
	// 		&obj.Ttd,
	// 		&obj.Nik,
	// 		&obj.Nama,
	// 		&obj.User_id,
	// 		&obj.Jenis,
	// 		&obj.Indikator1,
	// 		&obj.Indikator2,
	// 		&obj.Indikator3,
	// 		&obj.Indikator4,
	// 		&obj.Indikator5,
	// 		&obj.Indikator6,
	// 		&obj.Skor,
	// 		&obj.Lampiran_ktp,
	// 		&obj.Lampiran_kk,
	// 		&obj.Lampiran_pernyataan,
	// 		&obj.Lampiran_foto_rumah,
	// 		&obj.Lampiran_lainnya,
	// 		&obj.Admin_reject,
	// 		&obj.Date_reject,
	// 		&obj.Keterangan_reject,
	// 		&obj.Status_permohonan,
	// 		&obj.Created_at,
	// 		&obj.Updated_at,
	// 	)

	// 	if err == sql.ErrNoRows {
	// 		// fmt.Println("User not found")
	// 		return obj, errors.New("data tidak ditemukan")
	// 	}

	// 	if err != nil {
	// 		// fmt.Println("Query error")
	// 		log.Println(err.Error())
	// 		return obj, errors.New("gagal mengambil data")
	// 	}
	// } else if status == "5" {
	// 	sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

	// 	err := con.QueryRow(sqlStatement, id).Scan(
	// 		&obj.Id,
	// 		&obj.Layanan,
	// 		&obj.Kode_permohonan,
	// 		&obj.Kelurahan,
	// 		&obj.Ttd,
	// 		&obj.Nik,
	// 		&obj.Nama,
	// 		&obj.User_id,
	// 		&obj.Jenis,
	// 		&obj.Indikator1,
	// 		&obj.Indikator2,
	// 		&obj.Indikator3,
	// 		&obj.Indikator4,
	// 		&obj.Indikator5,
	// 		&obj.Indikator6,
	// 		&obj.Skor,
	// 		&obj.Lampiran_ktp,
	// 		&obj.Lampiran_kk,
	// 		&obj.Lampiran_pernyataan,
	// 		&obj.Lampiran_foto_rumah,
	// 		&obj.Lampiran_lainnya,
	// 		&obj.Admin_reject,
	// 		&obj.Date_reject,
	// 		&obj.Keterangan_reject,
	// 		&obj.Status_permohonan,
	// 		&obj.Created_at,
	// 		&obj.Updated_at,
	// 	)

	// 	if err == sql.ErrNoRows {
	// 		// fmt.Println("User not found")
	// 		return obj, errors.New("data tidak ditemukan")
	// 	}

	// 	if err != nil {
	// 		// fmt.Println("Query error")
	// 		log.Println(err.Error())
	// 		return obj, errors.New("gagal mengambil data")
	// 	}
	// }

	return obj, nil
}
