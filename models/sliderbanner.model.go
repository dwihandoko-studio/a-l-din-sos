package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"
	"net/http"
)

type Sliderbanner struct {
	Id                    *string `json:"id"`
	Slider_title          *string `json:"slider_title"`
	Slider_description    *string `json:"slider_description"`
	Slider_url            *string `json:"slider_url"`
	Bg_color              *string `json:"bg_color"`
	Slider_featured_image *string `json:"slider_featured_image"`
	Slider_is_active      *string `json:"slider_is_active"`
	Slider_user_id        int     `json:"slider_user_id"`
	Slider_created_at     *string `json:"slider_created_at"`
	Slider_updated_at     *string `json:"slider_updated_at"`
}

func GetSliderBanner() (Response, error) {
	var res Response
	var obj Sliderbanner
	var arrobj []Sliderbanner
	con := db.CreateCon()

	sqlStatement := "SELECT * From _slider_tb_b WHERE slider_is_active = 1 ORDER BY slider_updated_at DESC"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		fmt.Println(err.Error())
		return res, errors.New("Gagal mengambil data.")
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Slider_title, &obj.Slider_description, &obj.Slider_url, &obj.Bg_color, &obj.Slider_featured_image, &obj.Slider_is_active, &obj.Slider_user_id, &obj.Slider_created_at, &obj.Slider_updated_at)
		if err != nil {
			fmt.Println(err.Error())
			return res, errors.New("Gagal mengambil data.")
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}
