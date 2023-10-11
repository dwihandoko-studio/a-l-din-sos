package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
	"net/http"
)

type Kecamatan struct {
	Id           *int    `json:"id"`
	Id_kabupaten *string `json:"id_kabupaten"`
	Nama         *string `json:"kecamatan"`
}

func GetKecamatan(id string) (Response, error) {
	var obj Kecamatan
	var arrobj []Kecamatan
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM ref_kecamatan where id_kabupaten = ? ORDER BY kecamatan ASC"

	rows, err := con.Query(sqlStatement, id)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New("Gagal mengambil data.")
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Id_kabupaten, &obj.Nama)
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
