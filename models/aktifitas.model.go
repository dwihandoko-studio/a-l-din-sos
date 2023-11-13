package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Aktifitas struct {
	Id         *string `json:"id"`
	User_id    *string `json:"user_id"`
	Keterangan *string `json:"keterangan"`
	Aksi       *string `json:"aksi"`
	Icon       *string `json:"icon"`
	Created_at *string `json:"created_at"`
	// Fullname   *string `json:"fullname"`
	// Nisn       *string `json:"nisn"`
	// Email      *string `json:"email"`
	// No_hp      *string `json:"no_hp"`
}

type DafAktifitas struct {
	Data      []Aktifitas `json:"data"`
	Total     int         `json:"total"`
	Page      int         `json:"page"`
	Last_page int         `json:"last_page"`
}

func GetAktifitas(queryString string, pageInitial string) (DafAktifitas, error) {
	var obj Aktifitas
	var arryobj []Aktifitas
	var result DafAktifitas
	// var query string
	perPage := 10
	var total int
	page, _ := strconv.Atoi(pageInitial)

	con := db.CreateCon()
	jumDat, err := con.Query("SELECT COUNT(*) as total FROM v_tb_riwayat_system_peserta " + queryString)
	if err != nil {
		// fmt.Println("Error line jumlah: " + err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer jumDat.Close()
	for jumDat.Next() {
		jumDat.Scan(&total)
	}

	queryString = "SELECT * FROM v_tb_riwayat_system_peserta " + queryString + " ORDER BY created_at DESC"
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
			&obj.Keterangan,
			&obj.Aksi,
			&obj.Icon,
			&obj.Created_at,
			// &obj.Fullname,
			// &obj.Nisn,
			// &obj.Email,
			// &obj.No_hp,
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

func InsertAktifitas(keterangan, aksi, jenis, user_id string) error {
	con := db.CreateCon()

	tx, _ := con.Begin()

	// var uuid = uuid.New()
	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	sqlStatementInsert := `INSERT INTO riwayat_system(user_id,keterangan,aksi,icon,created_at) VALUES (?,?,?,?,?)`

	exeInsertDoc, err := tx.Exec(sqlStatementInsert, user_id, keterangan, aksi, jenis, currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		// fmt.Println("Query error")
		// fmt.Println(err.Error())
		tx.Rollback()
		return errors.New("gagal insert riwayat system")
	}

	rowsInsertDoc, err := exeInsertDoc.RowsAffected()
	if err != nil {
		// fmt.Println(err.Error())
		tx.Rollback()
		return errors.New("gagal insert riwayat system")
	}

	if rowsInsertDoc > 0 {
		tx.Commit()
		return nil

	}
	tx.Rollback()
	return errors.New("agal menyimpan riwayat")
}
