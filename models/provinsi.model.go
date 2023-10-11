package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
	"net/http"
)

type Provinsi struct {
	Id   int    `json:"id"`
	Nama string `json:"provinsi"`
}

func GetProvinsi() (Response, error) {
	var obj Provinsi
	var arrobj []Provinsi
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM ref_provinsi ORDER BY provinsi ASC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New("Gagal mengambil data.")
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama)
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
