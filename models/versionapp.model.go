package models

import (
	"api-dinsos/db"
	"database/sql"
	"errors"
)

type Versionapp struct {
	Id               int64   `json:"id"`
	App_name         *string `json:"app_name"`
	Package_name     *string `json:"package_name"`
	Version_app      *string `json:"version_app"`
	Build_number_app *string `json:"build_number_app"`
	Url_playstore    *string `json:"url_playstore"`
	Url_appstore     *string `json:"url_appstore"`
	Url_windows      *string `json:"url_windows"`
	Is_active        int     `json:"is_active"`
	Created_at       *string `json:"created_at"`
	Updated_at       *string `json:"updated_at"`
}

func GetVersionApp() (Response, error) {
	var res Response
	var obj Versionapp
	con := db.CreateCon()

	sqlStatement := "SELECT * From _version_app_tb WHERE is_active = 1 ORDER BY updated_at DESC"

	err := con.QueryRow(sqlStatement).Scan(
		&obj.Id, &obj.App_name, &obj.Package_name, &obj.Version_app, &obj.Build_number_app, &obj.Url_playstore, &obj.Url_appstore, &obj.Url_windows, &obj.Is_active, &obj.Created_at, &obj.Updated_at,
	)

	if err == sql.ErrNoRows {
		// fmt.Println("Version not found")
		return res, errors.New("gagal mengambil version app")
	}

	if err != nil {
		// fmt.Println("Query error")
		return res, errors.New("gagal mengambil version app")
	}
	res.Status = 200
	res.Message = "Version app berhasil dimuat."
	res.Data = obj
	return res, nil
}
