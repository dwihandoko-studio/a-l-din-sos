package models

import (
	"api-dinsos/db"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type LksIns struct {
	Id_permohonan                  *string `json:"id_permohonan"`
	Nama_lembaga                   *string `json:"nama_lembaga"`
	Jenis_lembaga                  *string `json:"jenis_lembaga"`
	Tgl_berdiri_lembaga            *string `json:"tgl_berdiri_lembaga"`
	Nama_notaris_lembaga           *string `json:"nama_notaris_lembaga"`
	Nomor_notaris_lembaga          *string `json:"nomor_notaris_lembaga"`
	Nomor_Kemenkumham_lembaga      *string `json:"nomor_Kemenkumham_lembaga"`
	Akreditasi_lembaga             *string `json:"akreditasi_lembaga"`
	Nomor_surat_akreditasi_lembaga *string `json:"nomor_surat_akreditasi_lembaga"`
	Tgl_expired_akreditasi_lembaga *string `json:"tgl_expired_akreditasi_lembaga"`
	Npwp_lembaga                   *string `json:"npwp_lembaga"`
	Modal_usaha_lembaga            int64   `json:"modal_usaha_lembaga"`
	Status_lembaga                 *string `json:"status_lembaga"`
	Lingkup_wilayah_kerja_lembaga  *string `json:"lingkup_wilayah_kerja_lembaga"`
	Bidang_kegiatan_lembaga        *string `json:"bidang_kegiatan_lembaga"`
	No_telp_lembaga                *string `json:"no_telp_lembaga"`
	Email_lembaga                  *string `json:"email_lembaga"`
	Lat_long_lembaga               *string `json:"lat_long_lembaga"`
	Alamat_lembaga                 *string `json:"alamat_lembaga"`
	Rt_lembaga                     *string `json:"rt_lembaga"`
	Rw_lembaga                     *string `json:"rw_lembaga"`
	Kecamatan_lembaga              *string `json:"kecamatan_lembaga"`
	Kelurahan_lembaga              *string `json:"kelurahan_lembaga"`
	Nama_ketua_pengurus            *string `json:"nama_ketua_pengurus"`
	Nik_ketua_pengurus             *string `json:"nik_ketua_pengurus"`
	Nohp_ketua_pengurus            *string `json:"nohp_ketua_pengurus"`
	Nama_sekretaris_pengurus       *string `json:"nama_sekretaris_pengurus"`
	Nik_sekretaris_pengurus        *string `json:"nik_sekretaris_pengurus"`
	Nohp_sekretaris_pengurus       *string `json:"nohp_sekretaris_pengurus"`
	Nama_bendahara_pengurus        *string `json:"nama_bendahara_pengurus"`
	Nik_bendahara_pengurus         *string `json:"nik_bendahara_pengurus"`
	Nohp_bendahara_pengurus        *string `json:"nohp_bendahara_pengurus"`
	Jumlah_pengurus                int64   `json:"jumlah_pengurus"`
	Jumlah_binaan_dalam            int64   `json:"jumlah_binaan_dalam"`
	Jumlah_binaan_luar             int64   `json:"jumlah_binaan_luar"`
}

func PostLks(userId, kode_permohonan, kelurahan, ttd, nik, nama, jenis, layanan, status_permohonan, created_at, filektpketua, filektpsekretaris, filektpbendahara, fileaktanotaris, filepengesahan_kemenkumham, fileadrt, fileketerangan_domisili, fileakreditasi, filestruktur_organisasi, filenpwp, filefoto_lokasi, filefoto_usaha_ekonomi, filelogo_lembaga, filedata_binaan string, dataLks LksIns) (Response, error) {

	var res Response

	con := db.CreateCon()

	tx, err := con.Begin()

	var uuid = uuid.New()
	var filenameKtpKetua *string
	var filenameKtpSekretaris *string
	var filenameKtpBendahara *string
	var filenameAktaNotaris *string
	var filenamePengesahanKemenkumham *string
	var filenameAdrt *string
	var filenameKeteranganDomisili *string
	var filenameAkreditasi *string
	var filenameStrukturOrganisasi *string
	var filenameNpwp *string
	var filenameFotoLokasi *string
	var filenameFotoUsahaEkonomi *string
	var filenameLogoLembaga *string
	var filenameDataBinaan *string
	if filektpketua != "" {
		filenameKtpKetua = &filektpketua
	}
	if filektpsekretaris != "" {
		filenameKtpSekretaris = &filektpsekretaris
	}
	if filektpbendahara != "" {
		filenameKtpBendahara = &filektpbendahara
	}
	if fileaktanotaris != "" {
		filenameAktaNotaris = &fileaktanotaris
	}
	if filepengesahan_kemenkumham != "" {
		filenamePengesahanKemenkumham = &filepengesahan_kemenkumham
	}
	if fileadrt != "" {
		filenameAdrt = &fileadrt
	}
	if fileketerangan_domisili != "" {
		filenameKeteranganDomisili = &fileketerangan_domisili
	}
	if fileakreditasi != "" {
		filenameAkreditasi = &fileakreditasi
	}
	if filestruktur_organisasi != "" {
		filenameStrukturOrganisasi = &filestruktur_organisasi
	}
	if filenpwp != "" {
		filenameNpwp = &filenpwp
	}
	if filefoto_lokasi != "" {
		filenameFotoLokasi = &filefoto_lokasi
	}
	if filefoto_usaha_ekonomi != "" {
		filenameFotoUsahaEkonomi = &filefoto_usaha_ekonomi
	}
	if filelogo_lembaga != "" {
		filenameLogoLembaga = &filelogo_lembaga
	}
	if filedata_binaan != "" {
		filenameDataBinaan = &filedata_binaan
	}

	sqlStatementDaftar := "INSERT INTO _permohonan_temp(id,kode_permohonan,kelurahan,ttd,nik,nama,user_id,jenis,layanan,status_permohonan,created_at) VALUES (?,?,?,?,?,?,?,?,?,?,?)"

	exeInsertDaftar, err := tx.Exec(sqlStatementDaftar, uuid, kode_permohonan, kelurahan, ttd, nik, nama, userId, jenis, layanan, status_permohonan, created_at)
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
		sqlStatementLks := "INSERT INTO _permohonan_lksa(id_permohonan,nama_lembaga,jenis_lembaga,tgl_berdiri_lembaga,nama_notaris_lembaga,nomor_notaris_lembaga,nomor_kemenkumham_lembaga,akreditasi_lembaga,nomor_surat_akreditasi_lembaga,tgl_expired_akreditasi_lembaga,npwp_lembaga,modal_usaha_lembaga, status_lembaga, lingkup_wilayah_kerja_lembaga, bidang_kegiatan_lembaga, no_telp_lembaga, email_lembaga, lat_long_lembaga, alamat_lembaga, rt_lembaga, rw_lembaga, kecamatan_lembaga, kelurahan_lembaga, nama_ketua_pengurus, nik_ketua_pengurus, nohp_ketua_pengurus, nama_sekretaris_pengurus, nik_sekretaris_pengurus, nohp_sekretaris_pengurus, nama_bendahara_pengurus, nik_bendahara_pengurus, nohp_bendahara_pengurus, jumlah_pengurus, jumlah_binaan_dalam, jumlah_binaan_luar, created_at, lampiran_ktp_ketua, lampiran_ktp_sekretaris, lampiran_ktp_bendahara, lampiran_akta_notaris, lampiran_kemenkumham, lampiran_adrt, lampiran_domisili, lampiran_akreditasi, lampiran_struktur_organisasi, lampiran_npwp, lampiran_foto_lokasi, lampiran_foto_usaha, lampiran_logo, lampiran_data_binaan) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)"

		exeInsertLks, err := tx.Exec(sqlStatementLks, uuid, dataLks.Nama_lembaga, dataLks.Jenis_lembaga, dataLks.Tgl_berdiri_lembaga, dataLks.Nama_notaris_lembaga, dataLks.Nomor_notaris_lembaga, dataLks.Nomor_Kemenkumham_lembaga, dataLks.Akreditasi_lembaga, dataLks.Nomor_surat_akreditasi_lembaga, dataLks.Tgl_expired_akreditasi_lembaga, dataLks.Npwp_lembaga, dataLks.Modal_usaha_lembaga, dataLks.Status_lembaga, dataLks.Lingkup_wilayah_kerja_lembaga, dataLks.Bidang_kegiatan_lembaga, dataLks.No_telp_lembaga, dataLks.Email_lembaga, dataLks.Lat_long_lembaga, dataLks.Alamat_lembaga, dataLks.Rt_lembaga, dataLks.Rw_lembaga, dataLks.Kecamatan_lembaga, dataLks.Kelurahan_lembaga, dataLks.Nama_ketua_pengurus, dataLks.Nik_ketua_pengurus, dataLks.Nohp_ketua_pengurus, dataLks.Nama_sekretaris_pengurus, dataLks.Nik_sekretaris_pengurus, dataLks.Nohp_sekretaris_pengurus, dataLks.Nama_bendahara_pengurus, dataLks.Nik_bendahara_pengurus, dataLks.Nohp_bendahara_pengurus, dataLks.Jumlah_pengurus, dataLks.Jumlah_binaan_dalam, dataLks.Jumlah_binaan_luar, created_at, filenameKtpKetua, filenameKtpSekretaris, filenameKtpBendahara, filenameAktaNotaris, filenamePengesahanKemenkumham, filenameAdrt, filenameKeteranganDomisili, filenameAkreditasi, filenameStrukturOrganisasi, filenameNpwp, filenameFotoLokasi, filenameFotoUsahaEkonomi, filenameLogoLembaga, filenameDataBinaan)
		if err != nil {
			fmt.Println("Query error")
			fmt.Println(err.Error())
			tx.Rollback()
			return res, errors.New("Gagal menyimpan permohonan.")
		}

		rowsInsertLks, err := exeInsertLks.RowsAffected()
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			return res, errors.New("Gagal menyimpan permohonan.")
		}

		if rowsInsertLks > 0 {

			tx.Commit()
			res.Status = 200
			res.Message = "Permohonan layanan " + layanan + " Berhasil Diajukan."

			return res, nil
		}
		tx.Rollback()
		return res, errors.New("Gagal menyimpan permohonan.")
		// tx.Commit()
		// res.Status = 200
		// res.Message = "Permohonan layanan " + layanan + " Berhasil Diajukan."

		// return res, nil
	}
	tx.Rollback()
	return res, errors.New("Gagal menyimpan permohonan.")
}

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
