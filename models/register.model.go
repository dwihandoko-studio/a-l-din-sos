package models

import (
	"api-dinsos/db"
	// "api-dinsos/helpers"
	// "database/sql"
	"errors"
	"log"
	// "net/http"
)

type Userregister struct {
	Id                string  `json:"id"`
	Email             string  `json:"email"`
	Nik               string  `json:"nik"`
	No_hp             *string `json:"no_hp"`
	Password          *string `json:"password"`
	Scope             string  `json:"scope"`
	Is_active         int     `json:"is_active"`
	Email_verified    int     `json:"email_verified"`
	Wa_verified       int     `json:"wa_verified"`
	Created_at        string  `json:"created_at"`
	Updated_at        *string `json:"updated_at"`
	Update_firs_login *string `json:"update_firs_login"`
}

type Userprofilregister struct {
	Id            string  `json:"id"`
	Fullname      *string `json:"fullname"`
	Email         *string `json:"email"`
	No_hp         *string `json:"no_hp"`
	Nik           *string `json:"nik"`
	Kk            *string `json:"kk"`
	Tempat_lahir  *string `json:"tempat_lahir"`
	Tgl_lahir     *string `json:"tgl_lahir"`
	Jenis_kelamin *string `json:"jenis_kelamin"`
	Kecamatan     *string `json:"kecamatan"`
	Kelurahan     *string `json:"kelurahan"`
	Alamat        *string `json:"alamat"`
	Pekerjaan     *string `json:"pekerjaan"`
	Image         *string `json:"image"`
	Role_user     int     `json:"role_user"`
	Last_active   *string `json:"last_active"`
	Created_at    *string `json:"created_at"`
	Updated_at    *string `json:"updated_at"`
}

func SaveRegistered(user Userregister, profil Userprofilregister) (Response, error) {
	var res Response

	con := db.CreateCon()

	tx, _ := con.Begin()

	sqlStatement := "INSERT INTO _users_tb(id,email,nik,no_hp,password,scope,is_active,email_verified,created_at) VALUES (?,?,?,?,?,?,?,?,?)"

	exe, err := tx.Exec(sqlStatement, user.Id, user.Email, user.Nik, user.No_hp, user.Password, user.Scope, user.Is_active, user.Email_verified, user.Created_at)

	if err != nil {
		log.Println("Query error")
		log.Println(err.Error())
		tx.Rollback()
		return res, errors.New("gagal membuat akun")
	}

	rows, err := exe.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		tx.Rollback()
		return res, errors.New("gagal membuat akun")
	}

	if rows > 0 {
		sqlStatementProfil := "INSERT INTO _profil_users_tb(id,fullname,email,no_hp,nik,kk,tempat_lahir,tgl_lahir,jenis_kelamin,kecamatan,kelurahan,alamat,pekerjaan,image,role_user,created_at) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

		exeProfil, err := tx.Exec(sqlStatementProfil,
			profil.Id,
			profil.Fullname,
			profil.Email,
			profil.No_hp,
			profil.Nik,
			profil.Kk,
			profil.Tempat_lahir,
			profil.Tgl_lahir,
			profil.Jenis_kelamin,
			profil.Kecamatan,
			profil.Kelurahan,
			profil.Alamat,
			profil.Pekerjaan,
			profil.Image,
			profil.Role_user,
			profil.Created_at,
		)

		if err != nil {
			log.Println("Query error")
			log.Println(err.Error())
			tx.Rollback()
			return res, errors.New("gagal membuat akun")
		}

		rowsProfil, err := exeProfil.RowsAffected()
		if err != nil {
			log.Println(err.Error())
			tx.Rollback()
			return res, errors.New("gagal membuat akun")
		}

		if rowsProfil > 0 {
			tx.Commit()
			res.Status = 200
			res.Message = "Akun Anda berhasil dibuat."
			res.Data = profil

			return res, nil
		}
		tx.Rollback()
		return res, errors.New("gagal membuat akun")
	}
	tx.Rollback()
	return res, errors.New("gagal membuat akun")
}
