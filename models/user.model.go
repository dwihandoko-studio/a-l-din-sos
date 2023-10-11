package models

import (
	"api-dinsos/db"
	"api-dinsos/helpers"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type Userdetail struct {
	Id             string  `json:"id"`
	Email          *string `json:"email"`
	Nik            *string `json:"nik"`
	Nohp           *string `json:"no_hp"`
	Is_active      int     `json:"is_active"`
	Email_verified int     `json:"email_verified"`
	Email_tertaut  int     `json:"email_tertaut"`
	Tautan_email   *string `json:"tautan_email"`
}

type Userprofil struct {
	Id                string  `json:"id"`
	Fullname          *string `json:"fullname"`
	Email             *string `json:"email"`
	No_hp             *string `json:"no_hp"`
	Nik               *string `json:"nik"`
	Kk                *string `json:"kk"`
	Tempat_lahir      *string `json:"tempat_lahir"`
	Tgl_lahir         *string `json:"tgl_lahir"`
	Jenis_kelamin     *string `json:"jenis_kelamin"`
	Provinsi          *string `json:"provinsi"`
	Kabupaten         *string `json:"kabupaten"`
	Kecamatan         *string `json:"kecamatan"`
	Kelurahan         *string `json:"kelurahan"`
	Alamat            *string `json:"alamat"`
	Pekerjaan         *string `json:"pekerjaan"`
	Image             *string `json:"image"`
	Role_user         int     `json:"role_user"`
	Last_active       *string `json:"last_active"`
	Created_at        *string `json:"created_at"`
	Updated_at        *string `json:"updated_at"`
	Is_active         int     `json:"is_active"`
	Wa_verified       int     `json:"wa_verified"`
	Email_verified    int     `json:"email_verified"`
	Email_tertaut     int     `json:"email_tertaut"`
	Update_firs_login *string `json:"update_firs_login"`
}

type Userpassword struct {
	Id             string  `json:"id"`
	Email          *string `json:"email"`
	Pass           *string `json:"password"`
	Is_active      int     `json:"is_active"`
	Email_verified int     `json:"email_verified"`
}

func UpdatePassword(id, old_password, new_password, date string) (Response, error) {
	var obj Userpassword

	var res Response

	con := db.CreateCon()

	tx, _ := con.Begin()

	sqlStatement := "SELECT id, email, password, is_active, email_verified From _users_tb WHERE id = ?"

	// errs := tx.QueryRow(sqlStatement, id)
	errs := tx.QueryRow(sqlStatement, id).Scan(
		&obj.Id, &obj.Email, &obj.Pass, &obj.Is_active, &obj.Email_verified,
	)

	if errs == sql.ErrNoRows {
		// fmt.Println(errs.Error())
		tx.Rollback()
		return res, errors.New("gagal mengambil informasi user")
	}

	if errs != nil {
		// fmt.Println("Query error")
		// fmt.Println(errs)
		tx.Rollback()
		return res, errors.New("gagal mengambil informasi user")
	}

	match, err := helpers.CheckPasswordHash(old_password, *obj.Pass)
	if err != nil {
		// fmt.Println(err.Error())
		return res, errors.New("kata sandi lama tidak sama")
	}
	if !match {
		// fmt.Println("Hash and password doesn't match.")
		return res, errors.New("kata sandi lama tidak sama")
	}

	hash, _ := helpers.HashPassword(new_password)

	sqlStatementProfil := `UPDATE _users_tb SET password=?, updated_at=? WHERE id=?`

	exeProfil, err := tx.Exec(sqlStatementProfil, hash, date, id)

	if err != nil {
		// fmt.Println("Query error")
		// fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("gagal mengupdate Kata sandi")
	}

	rowsProfil, err := exeProfil.RowsAffected()
	if err != nil {
		// fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("gagal mengupdate kata sandi")
	}

	if rowsProfil > 0 {
		tx.Commit()
		res.Status = 200
		res.Message = "Kata sandi baru berhasil disimpan."

		return res, nil
	}
	tx.Rollback()
	return res, errors.New("gagal mengupdate kata sandi")
}

func UpdateUser(id, provinsi, kabupaten, kecamatan, kelurahan, kk, alamat, tempat_lahir, tgl_lahir, nama, jenis_kelamin, date string) (Response, error) {
	var obj Userdetail

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()

	sqlStatement := "SELECT id, email, nik, no_hp, is_active, email_verified, email_tertaut, tautan_email From _users_tb WHERE id = ?"

	// errs := tx.QueryRow(sqlStatement, id)
	errs := tx.QueryRow(sqlStatement, id).Scan(
		&obj.Id,
		&obj.Email,
		&obj.Nik,
		&obj.Nohp,
		&obj.Is_active,
		&obj.Email_verified,
		&obj.Email_tertaut,
		&obj.Tautan_email,
	)

	if errs == sql.ErrNoRows {
		fmt.Println(errs.Error())
		tx.Rollback()
		return res, errors.New("Gagal mengambil informasi user.")
	}

	if errs != nil {
		fmt.Println("Query error")
		fmt.Println(errs)
		tx.Rollback()
		return res, errors.New("Gagal mengambil informasi user.")
	}

	sqlStatementProfil := "UPDATE _profil_users_tb SET fullname=?, provinsi=?, kabupaten=?, kecamatan=?, kelurahan=?, tempat_lahir=?, alamat=?, tgl_lahir=?, kk=?, jenis_kelamin=?, updated_at=? WHERE id=?"

	exeProfil, err := tx.Exec(sqlStatementProfil, nama, provinsi, kabupaten, kecamatan, kelurahan, tempat_lahir, alamat, tgl_lahir, kk, jenis_kelamin, date, id)

	if err != nil {
		fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal mengupdate profil.")
	}

	rowsProfil, err := exeProfil.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal mengupdate profil.")
	}

	if rowsProfil > 0 {
		tx.Commit()
		res.Status = 200
		res.Message = "Profil berhasil diupdate."

		return res, nil
	}
	tx.Rollback()
	return res, errors.New("Gagal mengupdate profil.")
}

func UpdateFotoProfil(id, foto, date string) (Userprofil, error) {
	var obj Userprofil

	con := db.CreateCon()

	tx, _ := con.Begin()

	sqlStatement := "SELECT * From v_user WHERE id = ?"

	// errs := tx.QueryRow(sqlStatement, id)
	errs := tx.QueryRow(sqlStatement, id).Scan(
		&obj.Id,
		&obj.Fullname,
		&obj.Email,
		&obj.No_hp,
		&obj.Nik,
		&obj.Kk,
		&obj.Tempat_lahir,
		&obj.Tgl_lahir,
		&obj.Jenis_kelamin,
		&obj.Provinsi,
		&obj.Kabupaten,
		&obj.Kecamatan,
		&obj.Kelurahan,
		&obj.Alamat,
		&obj.Pekerjaan,
		&obj.Image,
		&obj.Role_user,
		&obj.Last_active,
		&obj.Created_at,
		&obj.Updated_at,
		&obj.Is_active,
		&obj.Wa_verified,
		&obj.Email_verified,
		&obj.Email_tertaut,
		&obj.Update_firs_login,
	)

	if errs == sql.ErrNoRows {
		// fmt.Println(errs.Error())
		tx.Rollback()
		return obj, errors.New("gagal mengambil informasi user")
	}

	if errs != nil {
		// fmt.Println("Query error")
		// fmt.Println(errs)
		tx.Rollback()
		return obj, errors.New("gagal mengambil informasi user")
	}

	sqlStatementProfil := "UPDATE _profil_users_tb SET image=?, updated_at=? WHERE id=?"

	exeProfil, err := tx.Exec(sqlStatementProfil, foto, date, id)

	if err != nil {
		// fmt.Println("Query error")
		// fmt.Println(err.Error())
		tx.Rollback()
		return obj, errors.New("gagal mengupdate foto profil")
	}

	rowsProfil, err := exeProfil.RowsAffected()
	if err != nil {
		// fmt.Println(err.Error())
		tx.Rollback()
		return obj, errors.New("gagal mengupdate foto profil")
	}

	if rowsProfil > 0 {
		tx.Commit()

		return obj, nil
	}
	tx.Rollback()
	return obj, errors.New("gagal mengupdate foto profil")
}

func GetUser(id string) (Userdetail, error) {
	var obj Userdetail

	con := db.CreateCon()

	sqlStatement := "SELECT id, email, nik, no_hp, is_active, email_verified, email_tertaut, tautan_email From _users_tb WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(
		&obj.Id,
		&obj.Email,
		&obj.Nik,
		&obj.Nohp,
		&obj.Is_active,
		&obj.Email_verified,
		&obj.Email_tertaut,
		&obj.Tautan_email,
	)

	if err == sql.ErrNoRows {
		// fmt.Println("User not found")
		return obj, errors.New("user tidak ditemukan")
	}

	if err != nil {
		// fmt.Println("Query error")
		// fmt.Println(err.Error())
		return obj, errors.New("gagal mengambil data")
	}

	return obj, nil
}

func GetUserDetail(id string) (Userprofil, error) {
	var obj Userprofil

	con := db.CreateCon()

	sqlStatement := "SELECT * From v_user WHERE id = ?"

	err := con.QueryRow(sqlStatement, id).Scan(
		&obj.Id,
		&obj.Fullname,
		&obj.Email,
		&obj.No_hp,
		&obj.Nik,
		&obj.Kk,
		&obj.Tempat_lahir,
		&obj.Tgl_lahir,
		&obj.Jenis_kelamin,
		&obj.Provinsi,
		&obj.Kabupaten,
		&obj.Kecamatan,
		&obj.Kelurahan,
		&obj.Alamat,
		&obj.Pekerjaan,
		&obj.Image,
		&obj.Role_user,
		&obj.Last_active,
		&obj.Created_at,
		&obj.Updated_at,
		&obj.Is_active,
		&obj.Wa_verified,
		&obj.Email_verified,
		&obj.Email_tertaut,
		&obj.Update_firs_login,
	)

	if err == sql.ErrNoRows {
		// fmt.Println("User not found")
		return obj, errors.New("user tidak ditemukan")
	}

	if err != nil {
		// fmt.Println("Query error")
		// fmt.Println(err.Error())
		return obj, errors.New("gagal mengambil data")
	}

	return obj, nil
}

type Users struct {
	Data      []Userprofil `json:"data"`
	Total     int          `json:"total"`
	Page      int          `json:"page"`
	Last_page int          `json:"last_page"`
}

func GetUsers(queryString string, startFrom string, lengthShow string) (Users, error) {
	var obj Userprofil
	var arryobj []Userprofil
	var result Users
	// var query string
	var total int
	start, _ := strconv.Atoi(startFrom)
	length, _ := strconv.Atoi(lengthShow)

	con := db.CreateCon()
	jumDat, err := con.Query("SELECT COUNT(*) as total FROM v_user " + queryString)
	if err != nil {
		// fmt.Println("Error line jumlah: " + err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer jumDat.Close()
	for jumDat.Next() {
		jumDat.Scan(&total)
	}

	queryString = "SELECT * From v_user " + queryString + " ORDER BY role_user DESC"
	queryString = fmt.Sprintf("%s LIMIT %d OFFSET %d", queryString, length, start)

	respo, err := con.Query(queryString)
	if err != nil {
		// fmt.Println(err.Error())
		return result, errors.New("gagal memuat data")
	}
	defer respo.Close()

	for respo.Next() {
		err = respo.Scan(
			&obj.Id,
			&obj.Fullname,
			&obj.Email,
			&obj.No_hp,
			&obj.Nik,
			&obj.Kk,
			&obj.Tempat_lahir,
			&obj.Tgl_lahir,
			&obj.Jenis_kelamin,
			&obj.Provinsi,
			&obj.Kabupaten,
			&obj.Kecamatan,
			&obj.Kelurahan,
			&obj.Kecamatan,
			&obj.Alamat,
			&obj.Image,
			&obj.Role_user,
			&obj.Last_active,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Is_active,
			&obj.Wa_verified,
			&obj.Email_verified,
			&obj.Email_tertaut,
			&obj.Update_firs_login,
		)
		if err != nil {
			// fmt.Println(err.Error())
			return result, errors.New("gagal mengambil data")
		}

		arryobj = append(arryobj, obj)
	}

	result.Data = arryobj
	result.Total = total

	return result, nil
}
