package models

import (
	"api-dinsos/db"
	"api-dinsos/helpers"
	"api-dinsos/services"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id       string  `json:"id"`
	Username string  `json:"email"`
	Nik      string  `json:"nik"`
	Nohp     *string `json:"no_hp"`
	Level    string  `json:"level"`
}

type Useremail struct {
	Id             string  `json:"id"`
	Username       string  `json:"email"`
	Nohp           *string `json:"no_hp"`
	Email_verified int     `json:"email_verified"`
}

func CheckLogin(username, password string) (Response, error) {
	var obj User
	var pwd *string

	// var scope *string
	var is_active *int
	var email_verified *int
	var wa_verified *int
	// var created_at *string
	// var updated_at *string
	// var update_firs_login *string
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT a.id, a.email, a.nik, a.no_hp, a.password, a.is_active, a.email_verified, a.wa_verified, b.role_user as level From _users_tb a LEFT JOIN _profil_users_tb b ON a.id = b.id WHERE a.email = '" + username + "' OR a.nik = '" + username + "' OR (a.no_hp = '" + username + "' AND a.wa_verified = 1)"

	err := con.QueryRow(sqlStatement).Scan(
		&obj.Id, &obj.Username, &obj.Nik, &obj.Nohp, &pwd, &is_active, &email_verified, &wa_verified, &obj.Level,
	)

	if err == sql.ErrNoRows {
		// fmt.Println(err.Error())
		return res, errors.New("user belum terdaftar")
	}

	if err != nil {
		// fmt.Println("Query error")
		log.Println(err.Error())
		return res, errors.New("gagal dalam memuat data")
	}

	match, err := helpers.CheckPasswordHash(password, *pwd)
	if err != nil {
		fmt.Println(err)
		return res, errors.New("username / password salah")
	}
	if !match {
		fmt.Println("Hash and password doesn't match.")
		// fmt.Println(err)
		return res, errors.New("username / password salah")
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = obj

	return res, nil
}

func SendResetUsingEmail(email string) (Response, error) {
	var obj Useremail
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, email, email_verified From _users_tb WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id,
		&obj.Username,
		&obj.Email_verified,
	)

	// erro = errors.New("email sudah terdaftar")

	if err == sql.ErrNoRows {
		return res, errors.New("email tidak terdaftar")
	}

	if err != nil {
		// fmt.Println("Query error")
		return res, err
	}

	if obj.Email_verified == 1 {
		token, errt := InsertToken(obj.Id, obj.Username)
		if errt != nil {
			return res, errt
		}
		urlWeb := os.Getenv("WEB_URL")

		content := `
				<h1>Reset Password</h1>
				<p>Anda menerima permintaan untuk mereset kata sandi untuk akun layanan dinsos kab. lampung tengah (Si-Lastri) dengan email <br/><strong>` + email + `</strong>.</p>
		<p>Silakan klik tautan di bawah ini untuk mereset kata sandi:</p>
		<p><a style="color: #fff !important;background-color: #556ee6;border-color: #556ee6;font-weight: 400;line-height: 1.5;color: #495057;text-align: center;vertical-align: middle;border: 1px solid transparent;border-top-color: transparent;border-right-color: transparent;border-bottom-color: transparent;border-left-color: transparent;padding: .47rem .75rem;font-size: .8125rem;border-radius: .25rem;transition: color .15s ease-in-out,background-color .15s ease-in-out,border-color .15s ease-in-out,box-shadow .15s ease-in-out,-webkit-box-shadow .15s ease-in-out;" href="` + urlWeb + `/auth/confirmresetpassword?token=` + *token.Token + `&email=` + email + `&code=` + *token.Code + `">Reset Password Sekarang</a></p>
		`
		errem := services.SendEmail(email, *token.Token, *token.Code, "Reset Password", content)
		if errem != nil {
			return res, errem
		}

		res.Status = http.StatusOK
		res.Message = "Tautan reset password berhasil dikirimkan ke email anda. Silahkan cek email anda, jika tidak ada di Kotak masuk silahkan cek di bagian folder spam."

		return res, nil
	}

	return res, errors.New("email belum terverifikasi")
}

func SendResetUsingNohp(email string) (Response, error) {
	var obj Useremail
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, email, email_verified From _users_tb WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id,
		&obj.Username,
		&obj.Email_verified,
	)

	// erro = errors.New("email sudah terdaftar")

	if err == sql.ErrNoRows {
		return res, errors.New("email tidak terdaftar")
	}

	if err != nil {
		// fmt.Println("Query error")
		return res, err
	}

	if obj.Email_verified == 1 {
		// err := services.SendEmail(email)
		// if err != nil {
		// 	return res, err
		// }

		res.Status = http.StatusOK
		res.Message = "Tautan reset password berhasil dikirimkan ke email anda. Silahkan cek email anda, jika tidak ada di Kotak masuk silahkan cek di bagian folder spam."

		return res, nil
	}

	return res, errors.New("email belum terverifikasi")
}

func CheckUserRegistrasiEmail(email string) (Response, error) {
	var obj User
	var res Response
	var erro error

	con := db.CreateCon()

	sqlStatement := "SELECT id From _users_tb WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id,
	)

	// erro = errors.New("email sudah terdaftar")

	if err == sql.ErrNoRows {
		res.Status = http.StatusOK
		res.Message = "Success"

		return res, nil
	}

	if err != nil {
		// fmt.Println("Query error")
		return res, err
	}

	erro = errors.New("email sudah terdaftar")

	return res, erro
}

func CheckUserRegistrasiNik(nik string) (Response, error) {
	var obj User
	var res Response
	var erro error

	con := db.CreateCon()

	sqlStatement := "SELECT id From _users_tb WHERE nik = ?"

	err := con.QueryRow(sqlStatement, nik).Scan(
		&obj.Id,
	)

	if err == sql.ErrNoRows {
		res.Status = http.StatusOK
		res.Message = "Success"

		return res, nil
	}

	if err != nil {
		// fmt.Println("Query error")
		return res, err
	}

	erro = errors.New("nik sudah terdaftar")

	return res, erro
}

func CheckUserRegistrasiNohp(nohp string) (Response, error) {
	var obj User
	var res Response
	var erro error

	con := db.CreateCon()

	sqlStatement := "SELECT id From _users_tb WHERE no_hp = ? AND wa_verified = 1"

	err := con.QueryRow(sqlStatement, nohp).Scan(
		&obj.Id,
	)

	if err == sql.ErrNoRows {
		res.Status = http.StatusOK
		res.Message = "Success"

		return res, nil
	}

	if err != nil {
		// fmt.Println("Query error")
		return res, err
	}

	erro = errors.New("no handphone sudah terdaftar")

	return res, erro
}

// func CreateUser(id, provinsi, kabupaten, kecamatan, kelurahan, dusun, alamat, latitude, longitude, date string) (models.Response, error) {
// 	var obj models.Userdetail

// 	var res models.Response

// 	con := db.CreateCon()

// 	tx, err := con.Begin()

// 	sqlStatement := "SELECT * From _users_profil_tb WHERE id = ?"

// 	// errs := tx.QueryRow(sqlStatement, id)
// 	errs := tx.QueryRow(sqlStatement, id).Scan(
// 		&obj.Id,
// 		&obj.Fullname,
// 		&obj.Email,
// 		&obj.Nip,
// 		&obj.Nohp,
// 		&obj.Jeniskelamin,
// 		&obj.Jabatan,
// 		&obj.Npsn,
// 		&obj.Provinsi,
// 		&obj.Kabupaten,
// 		&obj.Kecamatan,
// 		&obj.Kelurahan,
// 		&obj.Dusun,
// 		&obj.Alamat,
// 		&obj.Surattugas,
// 		&obj.Profilepicture,
// 		&obj.Roleuser,
// 		&obj.Sekolahasal,
// 		&obj.Npsnasal,
// 		&obj.Pesertadidikid,
// 		&obj.Nisn,
// 		&obj.Latitude,
// 		&obj.Longitude,
// 		&obj.Sekolahid,
// 		&obj.Details,
// 		&obj.Lastactive,
// 		&obj.Createdat,
// 		&obj.Updateat,
// 		&obj.Edited_map,
// 	)

// 	if errs == sql.ErrNoRows {
// 		fmt.Println(errs.Error())
// 		tx.Rollback()
// 		return res, errors.New("Gagal mengambil informasi user.")
// 	}

// 	if errs != nil {
// 		fmt.Println("Query error")
// 		fmt.Println(errs)
// 		tx.Rollback()
// 		return res, errors.New("Gagal mengambil informasi user.")
// 	}

// 	sqlStatementProfil := "UPDATE _users_profil_tb SET edited_map=?, provinsi=?, kabupaten=?, kecamatan=?, kelurahan=?, dusun=?, alamat=?, latitude=?, longitude=?, updated_at=? WHERE id=?"

// 	exeProfil, err := tx.Exec(sqlStatementProfil, 1, provinsi, kabupaten, kecamatan, kelurahan, dusun, alamat, latitude, longitude, date, id)

// 	if err != nil {
// 		fmt.Println("Query error")
// 		fmt.Println(err.Error())
// 		tx.Rollback()
// 		return res, errors.New("Gagal mengupdate alamat profil.")
// 	}

// 	rowsProfil, err := exeProfil.RowsAffected()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		tx.Rollback()
// 		return res, errors.New("Gagal mengupdate alamat profil.")
// 	}

// 	if rowsProfil > 0 {
// 		tx.Commit()
// 		res.Status = 200
// 		res.Message = "Alamat berhasil diupdate."

// 		return res, nil
// 	}
// 	tx.Rollback()
// 	return res, errors.New("Gagal mengupdate alamat profil.")
// }
