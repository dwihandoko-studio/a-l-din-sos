package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func PostDtks(userId, kode_permohonan, kelurahan, ttd, nik, nama, jenis, layanan, status_permohonan, created_at, filektp, filekk, filefotorumah, filelainnya string) (Response, error) {

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()

	var uuid = uuid.New()
	var filenameLainnya *string
	var filenameKtp *string
	var filenameKk *string
	var filenameFotorumah *string
	if filelainnya != "" {
		filenameLainnya = &filelainnya
	}
	if filektp != "" {
		filenameKtp = &filektp
	}
	if filekk != "" {
		filenameKk = &filekk
	}
	if filefotorumah != "" {
		filenameFotorumah = &filefotorumah
	}

	sqlStatementDaftar := "INSERT INTO _permohonan_temp(id,kode_permohonan,kelurahan,ttd,nik,nama,user_id,jenis,layanan,status_permohonan,created_at,lampiran_ktp,lampiran_kk,lampiran_foto_rumah,lampiran_lainnya) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	exeInsertDaftar, err := tx.Exec(sqlStatementDaftar, uuid, kode_permohonan, kelurahan, ttd, nik, nama, userId, jenis, layanan, status_permohonan, created_at, filenameKtp, filenameKk, filenameFotorumah, filenameLainnya)
	if err != nil {
		fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
	}

	rowsInsertDaftar, err := exeInsertDaftar.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
	}

	if rowsInsertDaftar > 0 {
		tx.Commit()
		res.Status = 200
		res.Message = "Permohonan layanan " + layanan + " Berhasil Diajukan."

		return res, nil
	}
	tx.Rollback()
	return res, errors.New("Gagal menyimpan permohonan.")
}

func PostSktm(userId, kode_permohonan, kelurahan, ttd, nik, nama, jenis, layanan, status_permohonan, created_at, filektp, filekk, filefotorumah, filelainnya, filepernyataan string, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6 int, skor float64) (Response, error) {

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()

	var uuid = uuid.New()
	var filenameLainnya *string
	var filenameKtp *string
	var filenameKk *string
	var filenameFotorumah *string
	var filenamePernyataan *string
	if filelainnya != "" {
		filenameLainnya = &filelainnya
	}
	if filektp != "" {
		filenameKtp = &filektp
	}
	if filekk != "" {
		filenameKk = &filekk
	}
	if filefotorumah != "" {
		filenameFotorumah = &filefotorumah
	}
	if filepernyataan != "" {
		filenamePernyataan = &filepernyataan
	}

	sqlStatementDaftar := "INSERT INTO _permohonan_temp(id,kode_permohonan,kelurahan,ttd,nik,nama,user_id,jenis,layanan,status_permohonan,created_at,lampiran_ktp,lampiran_kk,lampiran_foto_rumah,lampiran_lainnya,lampiran_pernyataan,indikator1,indikator2,indikator3,indikator4,indikator5,indikator6,skor) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	exeInsertDaftar, err := tx.Exec(sqlStatementDaftar, uuid, kode_permohonan, kelurahan, ttd, nik, nama, userId, jenis, layanan, status_permohonan, created_at, filenameKtp, filenameKk, filenameFotorumah, filenameLainnya, filenamePernyataan, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor)
	if err != nil {
		fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
	}

	rowsInsertDaftar, err := exeInsertDaftar.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
	}

	if rowsInsertDaftar > 0 {
		tx.Commit()
		res.Status = 200
		res.Message = "Permohonan layanan " + layanan + " Berhasil Diajukan."

		return res, nil
	}
	tx.Rollback()
	return res, errors.New("Gagal menyimpan permohonan.")
}

func PostPbi(userId, kode_permohonan, kelurahan, ttd, nik, nama, jenis, layanan, status_permohonan, created_at, filektp, filekk, filefotorumah, filelainnya, filesktm string) (Response, error) {

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()

	var uuid = uuid.New()
	var filenameLainnya *string
	var filenameKtp *string
	var filenameKk *string
	var filenameFotorumah *string
	var filenameSktm *string
	if filelainnya != "" {
		filenameLainnya = &filelainnya
	}
	if filektp != "" {
		filenameKtp = &filektp
	}
	if filekk != "" {
		filenameKk = &filekk
	}
	if filefotorumah != "" {
		filenameFotorumah = &filefotorumah
	}
	if filesktm != "" {
		filenameSktm = &filesktm
	}

	sqlStatementDaftar := "INSERT INTO _permohonan_temp(id,kode_permohonan,kelurahan,ttd,nik,nama,user_id,jenis,layanan,status_permohonan,created_at,lampiran_ktp,lampiran_kk,lampiran_foto_rumah,lampiran_lainnya,lampiran_pernyataan) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

	exeInsertDaftar, err := tx.Exec(sqlStatementDaftar, uuid, kode_permohonan, kelurahan, ttd, nik, nama, userId, jenis, layanan, status_permohonan, created_at, filenameKtp, filenameKk, filenameFotorumah, filenameLainnya, filenameSktm)
	if err != nil {
		fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
	}

	rowsInsertDaftar, err := exeInsertDaftar.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
	}

	if rowsInsertDaftar > 0 {
		tx.Commit()
		res.Status = 200
		res.Message = "Permohonan layanan " + layanan + " Berhasil Diajukan."

		return res, nil
	}
	tx.Rollback()
	return res, errors.New("Gagal menyimpan permohonan.")
}
