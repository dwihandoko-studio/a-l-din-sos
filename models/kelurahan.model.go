package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
	"net/http"
)

type Kelurahan struct {
	Id           *int    `json:"id"`
	Id_kecamatan *string `json:"id_kecamatan"`
	Nama         *string `json:"kelurahan"`
}

func GetKelurahan(id string) (Response, error) {
	var obj Kelurahan
	var arrobj []Kelurahan
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM ref_kelurahan where id_kecamatan = ? ORDER BY kelurahan ASC"

	rows, err := con.Query(sqlStatement, id)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New("Gagal mengambil data.")
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Id_kecamatan, &obj.Nama)
		if err != nil {
			fmt.Println(err.Error())
			return res, errors.New("Gagal mengambil data.")
		}

		arrobj = append(arrobj, obj)
	}
	if len(arrobj) > 0 {
		res.Status = http.StatusOK
	} else {
		res.Status = http.StatusNoContent
	}
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
