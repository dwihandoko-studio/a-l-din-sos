package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
	"net/http"
)

type Kabupaten struct {
	Id          *int    `json:"id"`
	Id_provinsi *string `json:"id_provinsi"`
	Nama        *string `json:"kabupaten"`
}

func GetKabupaten(id string) (Response, error) {
	var obj Kabupaten
	var arrobj []Kabupaten
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM ref_kabupaten where id_provinsi = ? ORDER BY kabupaten ASC"

	rows, err := con.Query(sqlStatement, id)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New("Gagal mengambil data.")
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Id_provinsi, &obj.Nama)
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
