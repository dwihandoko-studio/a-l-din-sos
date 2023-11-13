package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
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
