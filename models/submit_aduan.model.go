package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func PostAduan(userId, kode_permohonan, nama, nik, nohp, alamat, kelurahan, kecamatan, nama_aduan, nik_aduan, nohp_aduan, alamat_aduan, kelurahan_aduan, kecamatan_aduan, kategori, identitas, uraian string, status int, created_at, pengaduan1, pengaduan2, pengaduan3, pengaduan4, pengaduan5 string) (Response, error) {

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()

	var uuid = uuid.New()
	var filenamePengaduan1 *string
	var filenamePengaduan2 *string
	var filenamePengaduan3 *string
	var filenamePengaduan4 *string
	var filenamePengaduan5 *string
	if pengaduan1 != "" {
		filenamePengaduan1 = &pengaduan1
	}
	if pengaduan2 != "" {
		filenamePengaduan2 = &pengaduan2
	}
	if pengaduan3 != "" {
		filenamePengaduan3 = &pengaduan3
	}
	if pengaduan4 != "" {
		filenamePengaduan4 = &pengaduan4
	}
	if pengaduan5 != "" {
		filenamePengaduan5 = &pengaduan5
	}

	sqlStatementDaftar := "INSERT INTO _pengaduan(id,kode_aduan,nama,nik,nohp,alamat,kelurahan,kecamatan,nama_aduan,nik_aduan,nohp_aduan,alamat_aduan,kelurahan_aduan,kecamatan_aduan,user_id,kategori,identitas_aduan,uraian_aduan,status_aduan,created_at,lampiran_aduan_1,lampiran_aduan_2,lampiran_aduan_3,lampiran_aduan_4,lampiran_aduan_5,media_pengaduan) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	exeInsertDaftar, err := tx.Exec(sqlStatementDaftar, uuid, kode_permohonan, nama, nik, nohp, alamat, kelurahan, kecamatan, nama_aduan, nik_aduan, nohp_aduan, alamat_aduan, kelurahan_aduan, kecamatan_aduan, userId, kategori, identitas, uraian, status, created_at, filenamePengaduan1, filenamePengaduan2, filenamePengaduan3, filenamePengaduan4, filenamePengaduan5, "Aplikasi Layanan")
	if err != nil {
		fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan pengaduan.")
	}

	rowsInsertDaftar, err := exeInsertDaftar.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan pengaduan.")
	}

	if rowsInsertDaftar > 0 {
		tx.Commit()
		res.Status = 200
		res.Message = "Pengaduan Berhasil Diajukan."

		return res, nil
	}
	tx.Rollback()
	return res, errors.New("Gagal menyimpan pengaduan.")
}
