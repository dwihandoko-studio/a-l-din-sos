package models

import (
	"api-dinsos/db"
	"database/sql"
	"errors"
	"log"
)

type Sktm struct {
	Id                  string   `json:"id"`
	Layanan             *string  `json:"layanan"`
	Kode_permohonan     string   `json:"kode_permohonan"`
	Kelurahan           *string  `json:"kelurahan"`
	Ttd                 *string  `json:"ttd"`
	Nik                 *string  `json:"nik"`
	Nama                *string  `json:"nama"`
	User_id             *string  `json:"user_id"`
	Jenis               *string  `json:"jenis"`
	Indikator1          *int64   `json:"indikator1"`
	Indikator2          *int64   `json:"indikator2"`
	Indikator3          *int64   `json:"indikator3"`
	Indikator4          *int64   `json:"indikator4"`
	Indikator5          *int64   `json:"indikator5"`
	Indikator6          *int64   `json:"indikator6"`
	Skor                *float64 `json:"skor"`
	Lampiran_ktp        *string  `json:"lampiran_ktp"`
	Lampiran_kk         *string  `json:"lampiran_kk"`
	Lampiran_pernyataan *string  `json:"lampiran_pernyataan"`
	Lampiran_foto_rumah *string  `json:"lampiran_foto_rumah"`
	Lampiran_lainnya    *string  `json:"lampiran_lainnya"`
	Admin_approve       *string  `json:"admin_approve"`
	Date_approve        *string  `json:"date_approve"`
	Admin_reject        *string  `json:"admin_reject"`
	Date_reject         *string  `json:"date_reject"`
	Keterangan_reject   *string  `json:"keterangan_reject"`
	Lampiran_selesai    *string  `json:"lampiran_selesai"`
	Keterangan_selesai  *string  `json:"keterangan_selesai"`
	Status_permohonan   int64    `json:"status_permohonan"`
	Created_at          *string  `json:"created_at"`
	Updated_at          *string  `json:"updated_at"`
}

type Pbi struct {
	Id                  string   `json:"id"`
	Layanan             *string  `json:"layanan"`
	Kode_permohonan     string   `json:"kode_permohonan"`
	Kelurahan           *string  `json:"kelurahan"`
	Ttd                 *string  `json:"ttd"`
	Nik                 *string  `json:"nik"`
	Nama                *string  `json:"nama"`
	User_id             *string  `json:"user_id"`
	Jenis               *string  `json:"jenis"`
	Indikator1          *int64   `json:"indikator1"`
	Indikator2          *int64   `json:"indikator2"`
	Indikator3          *int64   `json:"indikator3"`
	Indikator4          *int64   `json:"indikator4"`
	Indikator5          *int64   `json:"indikator5"`
	Indikator6          *int64   `json:"indikator6"`
	Skor                *float64 `json:"skor"`
	Lampiran_ktp        *string  `json:"lampiran_ktp"`
	Lampiran_kk         *string  `json:"lampiran_kk"`
	Lampiran_pernyataan *string  `json:"lampiran_pernyataan"`
	Lampiran_foto_rumah *string  `json:"lampiran_foto_rumah"`
	Lampiran_lainnya    *string  `json:"lampiran_lainnya"`
	Admin_approve       *string  `json:"admin_approve"`
	Date_approve        *string  `json:"date_approve"`
	Admin_reject        *string  `json:"admin_reject"`
	Date_reject         *string  `json:"date_reject"`
	Keterangan_reject   *string  `json:"keterangan_reject"`
	Lampiran_selesai    *string  `json:"lampiran_selesai"`
	Keterangan_selesai  *string  `json:"keterangan_selesai"`
	Status_permohonan   int64    `json:"status_permohonan"`
	Created_at          *string  `json:"created_at"`
	Updated_at          *string  `json:"updated_at"`
}

type Lks struct {
	Id                             string  `json:"id"`
	Layanan                        *string `json:"layanan"`
	Kode_permohonan                string  `json:"kode_permohonan"`
	Kelurahan                      *string `json:"kelurahan"`
	Ttd                            *string `json:"ttd"`
	Nik                            *string `json:"nik"`
	Nama                           *string `json:"nama"`
	User_id                        *string `json:"user_id"`
	Jenis                          *string `json:"jenis"`
	Admin_approve                  *string `json:"admin_approve"`
	Date_approve                   *string `json:"date_approve"`
	Admin_reject                   *string `json:"admin_reject"`
	Date_reject                    *string `json:"date_reject"`
	Keterangan_reject              *string `json:"keterangan_reject"`
	Lampiran_selesai               *string `json:"lampiran_selesai"`
	Keterangan_selesai             *string `json:"keterangan_selesai"`
	Status_permohonan              int64   `json:"status_permohonan"`
	Created_at                     *string `json:"created_at"`
	Updated_at                     *string `json:"updated_at"`
	Nama_lembaga                   *string `json:"nama_lembaga"`
	Jenis_lembaga                  *string `json:"jenis_lembaga"`
	Tgl_berdiri_lembaga            *string `json:"tgl_berdiri_lembaga"`
	Nama_notaris_lembaga           *string `json:"nama_notaris_lembaga"`
	Nomor_notaris_lembaga          *string `json:"nomor_notaris_lembaga"`
	Nomor_kemenkumham_lembaga      *string `json:"nomor_kemenkumham_lembaga"`
	Akreditasi_lembaga             *string `json:"akreditasi_lembaga"`
	Nomor_surat_akreditasi_lembaga *string `json:"nomor_surat_akreditasi_lembaga"`
	Tgl_expired_akreditasi_lembaga *string `json:"tgl_expired_akreditasi_lembaga"`
	Npwp_lembaga                   *string `json:"npwp_lembaga"`
	Modal_usaha_lembaga            *int64  `json:"modal_usaha_lembaga"`
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
	Jumlah_pengurus                *int64  `json:"jumlah_pengurus"`
	Jumlah_binaan_dalam            *int64  `json:"jumlah_binaan_dalam"`
	Jumlah_binaan_luar             *int64  `json:"jumlah_binaan_luar"`
	Lampiran_ktp_ketua             *string `json:"lampiran_ktp_ketua"`
	Lampiran_ktp_sekretaris        *string `json:"lampiran_ktp_sekretaris"`
	Lampiran_ktp_bendahara         *string `json:"lampiran_ktp_bendahara"`
	Lampiran_akta_notaris          *string `json:"lampiran_akta_notaris"`
	Lampiran_kemenkumham           *string `json:"lampiran_kemenkumham"`
	Lampiran_adrt                  *string `json:"lampiran_adrt"`
	Lampiran_domisili              *string `json:"lampiran_domisili"`
	Lampiran_akreditasi            *string `json:"lampiran_akreditasi"`
	Lampiran_struktur_organisasi   *string `json:"lampiran_struktur_organisasi"`
	Lampiran_npwp                  *string `json:"lampiran_npwp"`
	Lampiran_foto_lokasi           *string `json:"lampiran_foto_lokasi"`
	Lampiran_foto_usaha            *string `json:"lampiran_foto_usaha"`
	Lampiran_logo                  *string `json:"lampiran_logo"`
	Lampiran_data_binaan           *string `json:"lampiran_data_binaan"`
}

type Skdtks struct {
	Id                  string   `json:"id"`
	Layanan             *string  `json:"layanan"`
	Kode_permohonan     string   `json:"kode_permohonan"`
	Kelurahan           *string  `json:"kelurahan"`
	Ttd                 *string  `json:"ttd"`
	Nik                 *string  `json:"nik"`
	Nama                *string  `json:"nama"`
	User_id             *string  `json:"user_id"`
	Jenis               *string  `json:"jenis"`
	Indikator1          *int64   `json:"indikator1"`
	Indikator2          *int64   `json:"indikator2"`
	Indikator3          *int64   `json:"indikator3"`
	Indikator4          *int64   `json:"indikator4"`
	Indikator5          *int64   `json:"indikator5"`
	Indikator6          *int64   `json:"indikator6"`
	Skor                *float64 `json:"skor"`
	Lampiran_ktp        *string  `json:"lampiran_ktp"`
	Lampiran_kk         *string  `json:"lampiran_kk"`
	Lampiran_pernyataan *string  `json:"lampiran_pernyataan"`
	Lampiran_foto_rumah *string  `json:"lampiran_foto_rumah"`
	Lampiran_lainnya    *string  `json:"lampiran_lainnya"`
	Admin_approve       *string  `json:"admin_approve"`
	Date_approve        *string  `json:"date_approve"`
	Admin_reject        *string  `json:"admin_reject"`
	Date_reject         *string  `json:"date_reject"`
	Keterangan_reject   *string  `json:"keterangan_reject"`
	Lampiran_selesai    *string  `json:"lampiran_selesai"`
	Keterangan_selesai  *string  `json:"keterangan_selesai"`
	Status_permohonan   int64    `json:"status_permohonan"`
	Created_at          *string  `json:"created_at"`
	Updated_at          *string  `json:"updated_at"`
}

func GetDetailLks(id, status string) (Lks, error) {
	var obj Lks

	con := db.CreateCon()

	if status == "0" {
		sqlStatement := "SELECT a.id, a.layanan, a.kode_permohonan, a.kelurahan, a.ttd, a.nik, a.nama, a.user_id, a.jenis, a.status_permohonan, a.created_at, a.updated_at, b.nama_lembaga, b.jenis_lembaga, b.tgl_berdiri_lembaga, b.nama_notaris_lembaga, b.nomor_notaris_lembaga, b.nomor_kemenkumham_lembaga, b.akreditasi_lembaga, b.nomor_surat_akreditasi_lembaga, b.tgl_expired_akreditasi_lembaga, b.npwp_lembaga, b.modal_usaha_lembaga, b.status_lembaga, b.lingkup_wilayah_kerja_lembaga, b.bidang_kegiatan_lembaga, b.no_telp_lembaga, b.email_lembaga, b.lat_long_lembaga, b.alamat_lembaga, b.rt_lembaga, b.rw_lembaga, (SELECT kecamatan FROM ref_kecamatan Where id = b.kecamatan_lembaga) as kecamatan_lembaga, (SELECT kelurahan FROM ref_kelurahan Where id = b.kelurahan_lembaga) as kelurahan_lembaga, b.nama_ketua_pengurus, b.nik_ketua_pengurus, b.nohp_ketua_pengurus, b.nama_sekretaris_pengurus, b.nik_sekretaris_pengurus, b.nohp_sekretaris_pengurus, b.nama_bendahara_pengurus, b.nik_bendahara_pengurus, b.nohp_bendahara_pengurus, b.jumlah_pengurus, b.jumlah_binaan_dalam, b.jumlah_binaan_luar, b.lampiran_ktp_ketua, b.lampiran_ktp_sekretaris, b.lampiran_ktp_bendahara, b.lampiran_akta_notaris, b.lampiran_kemenkumham, b.lampiran_adrt, b.lampiran_domisili, b.lampiran_akreditasi, b.lampiran_struktur_organisasi, b.lampiran_npwp, b.lampiran_foto_lokasi, b.lampiran_foto_usaha, b.lampiran_logo, b.lampiran_data_binaan From _permohonan_temp a INNER JOIN _permohonan_lksa b ON b.id_permohonan = a.id WHERE a.id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Nama_lembaga,
			&obj.Jenis_lembaga,
			&obj.Tgl_berdiri_lembaga,
			&obj.Nama_notaris_lembaga,
			&obj.Nomor_notaris_lembaga,
			&obj.Nomor_kemenkumham_lembaga,
			&obj.Akreditasi_lembaga,
			&obj.Nomor_surat_akreditasi_lembaga,
			&obj.Tgl_expired_akreditasi_lembaga,
			&obj.Npwp_lembaga,
			&obj.Modal_usaha_lembaga,
			&obj.Status_lembaga,
			&obj.Lingkup_wilayah_kerja_lembaga,
			&obj.Bidang_kegiatan_lembaga,
			&obj.No_telp_lembaga,
			&obj.Email_lembaga,
			&obj.Lat_long_lembaga,
			&obj.Alamat_lembaga,
			&obj.Rt_lembaga,
			&obj.Rw_lembaga,
			&obj.Kecamatan_lembaga,
			&obj.Kelurahan_lembaga,
			&obj.Nama_ketua_pengurus,
			&obj.Nik_ketua_pengurus,
			&obj.Nohp_ketua_pengurus,
			&obj.Nama_sekretaris_pengurus,
			&obj.Nik_sekretaris_pengurus,
			&obj.Nohp_sekretaris_pengurus,
			&obj.Nama_bendahara_pengurus,
			&obj.Nik_bendahara_pengurus,
			&obj.Nohp_bendahara_pengurus,
			&obj.Jumlah_pengurus,
			&obj.Jumlah_binaan_dalam,
			&obj.Jumlah_binaan_luar,
			&obj.Lampiran_ktp_ketua,
			&obj.Lampiran_ktp_sekretaris,
			&obj.Lampiran_ktp_bendahara,
			&obj.Lampiran_akta_notaris,
			&obj.Lampiran_kemenkumham,
			&obj.Lampiran_adrt,
			&obj.Lampiran_domisili,
			&obj.Lampiran_akreditasi,
			&obj.Lampiran_struktur_organisasi,
			&obj.Lampiran_npwp,
			&obj.Lampiran_foto_lokasi,
			&obj.Lampiran_foto_usaha,
			&obj.Lampiran_logo,
			&obj.Lampiran_data_binaan,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "1" {
		sqlStatement := "SELECT a.id, a.layanan, a.kode_permohonan, a.kelurahan, a.ttd, a.nik, a.nama, a.user_id, a.jenis, a.status_permohonan, a.created_at, a.updated_at, b.nama_lembaga, b.jenis_lembaga, b.tgl_berdiri_lembaga, b.nama_notaris_lembaga, b.nomor_notaris_lembaga, b.nomor_kemenkumham_lembaga, b.akreditasi_lembaga, b.nomor_surat_akreditasi_lembaga, b.tgl_expired_akreditasi_lembaga, b.npwp_lembaga, b.modal_usaha_lembaga, b.status_lembaga, b.lingkup_wilayah_kerja_lembaga, b.bidang_kegiatan_lembaga, b.no_telp_lembaga, b.email_lembaga, b.lat_long_lembaga, b.alamat_lembaga, b.rt_lembaga, b.rw_lembaga, (SELECT kecamatan FROM ref_kecamatan Where id = b.kecamatan_lembaga) as kecamatan_lembaga, (SELECT kelurahan FROM ref_kelurahan Where id = b.kelurahan_lembaga) as kelurahan_lembaga, b.nama_ketua_pengurus, b.nik_ketua_pengurus, b.nohp_ketua_pengurus, b.nama_sekretaris_pengurus, b.nik_sekretaris_pengurus, b.nohp_sekretaris_pengurus, b.nama_bendahara_pengurus, b.nik_bendahara_pengurus, b.nohp_bendahara_pengurus, b.jumlah_pengurus, b.jumlah_binaan_dalam, b.jumlah_binaan_luar, b.lampiran_ktp_ketua, b.lampiran_ktp_sekretaris, b.lampiran_ktp_bendahara, b.lampiran_akta_notaris, b.lampiran_kemenkumham, b.lampiran_adrt, b.lampiran_domisili, b.lampiran_akreditasi, b.lampiran_struktur_organisasi, b.lampiran_npwp, b.lampiran_foto_lokasi, b.lampiran_foto_usaha, b.lampiran_logo, b.lampiran_data_binaan From _permohonan_temp a INNER JOIN _permohonan_lksa b ON b.id_permohonan = a.id WHERE a.id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Nama_lembaga,
			&obj.Jenis_lembaga,
			&obj.Tgl_berdiri_lembaga,
			&obj.Nama_notaris_lembaga,
			&obj.Nomor_notaris_lembaga,
			&obj.Nomor_kemenkumham_lembaga,
			&obj.Akreditasi_lembaga,
			&obj.Nomor_surat_akreditasi_lembaga,
			&obj.Tgl_expired_akreditasi_lembaga,
			&obj.Npwp_lembaga,
			&obj.Modal_usaha_lembaga,
			&obj.Status_lembaga,
			&obj.Lingkup_wilayah_kerja_lembaga,
			&obj.Bidang_kegiatan_lembaga,
			&obj.No_telp_lembaga,
			&obj.Email_lembaga,
			&obj.Lat_long_lembaga,
			&obj.Alamat_lembaga,
			&obj.Rt_lembaga,
			&obj.Rw_lembaga,
			&obj.Kecamatan_lembaga,
			&obj.Kelurahan_lembaga,
			&obj.Nama_ketua_pengurus,
			&obj.Nik_ketua_pengurus,
			&obj.Nohp_ketua_pengurus,
			&obj.Nama_sekretaris_pengurus,
			&obj.Nik_sekretaris_pengurus,
			&obj.Nohp_sekretaris_pengurus,
			&obj.Nama_bendahara_pengurus,
			&obj.Nik_bendahara_pengurus,
			&obj.Nohp_bendahara_pengurus,
			&obj.Jumlah_pengurus,
			&obj.Jumlah_binaan_dalam,
			&obj.Jumlah_binaan_luar,
			&obj.Lampiran_ktp_ketua,
			&obj.Lampiran_ktp_sekretaris,
			&obj.Lampiran_ktp_bendahara,
			&obj.Lampiran_akta_notaris,
			&obj.Lampiran_kemenkumham,
			&obj.Lampiran_adrt,
			&obj.Lampiran_domisili,
			&obj.Lampiran_akreditasi,
			&obj.Lampiran_struktur_organisasi,
			&obj.Lampiran_npwp,
			&obj.Lampiran_foto_lokasi,
			&obj.Lampiran_foto_usaha,
			&obj.Lampiran_logo,
			&obj.Lampiran_data_binaan,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "2" {
		sqlStatement := "SELECT a.id, a.layanan, a.kode_permohonan, a.kelurahan, a.ttd, a.nik, a.nama, a.user_id, a.jenis, a.status_permohonan, a.created_at, a.updated_at, b.nama_lembaga, b.jenis_lembaga, b.tgl_berdiri_lembaga, b.nama_notaris_lembaga, b.nomor_notaris_lembaga, b.nomor_kemenkumham_lembaga, b.akreditasi_lembaga, b.nomor_surat_akreditasi_lembaga, b.tgl_expired_akreditasi_lembaga, b.npwp_lembaga, b.modal_usaha_lembaga, b.status_lembaga, b.lingkup_wilayah_kerja_lembaga, b.bidang_kegiatan_lembaga, b.no_telp_lembaga, b.email_lembaga, b.lat_long_lembaga, b.alamat_lembaga, b.rt_lembaga, b.rw_lembaga, (SELECT kecamatan FROM ref_kecamatan Where id = b.kecamatan_lembaga) as kecamatan_lembaga, (SELECT kelurahan FROM ref_kelurahan Where id = b.kelurahan_lembaga) as kelurahan_lembaga, b.nama_ketua_pengurus, b.nik_ketua_pengurus, b.nohp_ketua_pengurus, b.nama_sekretaris_pengurus, b.nik_sekretaris_pengurus, b.nohp_sekretaris_pengurus, b.nama_bendahara_pengurus, b.nik_bendahara_pengurus, b.nohp_bendahara_pengurus, b.jumlah_pengurus, b.jumlah_binaan_dalam, b.jumlah_binaan_luar, b.lampiran_ktp_ketua, b.lampiran_ktp_sekretaris, b.lampiran_ktp_bendahara, b.lampiran_akta_notaris, b.lampiran_kemenkumham, b.lampiran_adrt, b.lampiran_domisili, b.lampiran_akreditasi, b.lampiran_struktur_organisasi, b.lampiran_npwp, b.lampiran_foto_lokasi, b.lampiran_foto_usaha, b.lampiran_logo, b.lampiran_data_binaan From _permohonan_temp a INNER JOIN _permohonan_lksa b ON b.id_permohonan = a.id WHERE a.id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Nama_lembaga,
			&obj.Jenis_lembaga,
			&obj.Tgl_berdiri_lembaga,
			&obj.Nama_notaris_lembaga,
			&obj.Nomor_notaris_lembaga,
			&obj.Nomor_kemenkumham_lembaga,
			&obj.Akreditasi_lembaga,
			&obj.Nomor_surat_akreditasi_lembaga,
			&obj.Tgl_expired_akreditasi_lembaga,
			&obj.Npwp_lembaga,
			&obj.Modal_usaha_lembaga,
			&obj.Status_lembaga,
			&obj.Lingkup_wilayah_kerja_lembaga,
			&obj.Bidang_kegiatan_lembaga,
			&obj.No_telp_lembaga,
			&obj.Email_lembaga,
			&obj.Lat_long_lembaga,
			&obj.Alamat_lembaga,
			&obj.Rt_lembaga,
			&obj.Rw_lembaga,
			&obj.Kecamatan_lembaga,
			&obj.Kelurahan_lembaga,
			&obj.Nama_ketua_pengurus,
			&obj.Nik_ketua_pengurus,
			&obj.Nohp_ketua_pengurus,
			&obj.Nama_sekretaris_pengurus,
			&obj.Nik_sekretaris_pengurus,
			&obj.Nohp_sekretaris_pengurus,
			&obj.Nama_bendahara_pengurus,
			&obj.Nik_bendahara_pengurus,
			&obj.Nohp_bendahara_pengurus,
			&obj.Jumlah_pengurus,
			&obj.Jumlah_binaan_dalam,
			&obj.Jumlah_binaan_luar,
			&obj.Lampiran_ktp_ketua,
			&obj.Lampiran_ktp_sekretaris,
			&obj.Lampiran_ktp_bendahara,
			&obj.Lampiran_akta_notaris,
			&obj.Lampiran_kemenkumham,
			&obj.Lampiran_adrt,
			&obj.Lampiran_domisili,
			&obj.Lampiran_akreditasi,
			&obj.Lampiran_struktur_organisasi,
			&obj.Lampiran_npwp,
			&obj.Lampiran_foto_lokasi,
			&obj.Lampiran_foto_usaha,
			&obj.Lampiran_logo,
			&obj.Lampiran_data_binaan,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "3" {
		sqlStatement := "SELECT a.id, a.layanan, a.kode_permohonan, a.kelurahan, a.ttd, a.nik, a.nama, a.user_id, a.jenis, a.status_permohonan, a.created_at, a.updated_at, b.nama_lembaga, b.jenis_lembaga, b.tgl_berdiri_lembaga, b.nama_notaris_lembaga, b.nomor_notaris_lembaga, b.nomor_kemenkumham_lembaga, b.akreditasi_lembaga, b.nomor_surat_akreditasi_lembaga, b.tgl_expired_akreditasi_lembaga, b.npwp_lembaga, b.modal_usaha_lembaga, b.status_lembaga, b.lingkup_wilayah_kerja_lembaga, b.bidang_kegiatan_lembaga, b.no_telp_lembaga, b.email_lembaga, b.lat_long_lembaga, b.alamat_lembaga, b.rt_lembaga, b.rw_lembaga, (SELECT kecamatan FROM ref_kecamatan Where id = b.kecamatan_lembaga) as kecamatan_lembaga, (SELECT kelurahan FROM ref_kelurahan Where id = b.kelurahan_lembaga) as kelurahan_lembaga, b.nama_ketua_pengurus, b.nik_ketua_pengurus, b.nohp_ketua_pengurus, b.nama_sekretaris_pengurus, b.nik_sekretaris_pengurus, b.nohp_sekretaris_pengurus, b.nama_bendahara_pengurus, b.nik_bendahara_pengurus, b.nohp_bendahara_pengurus, b.jumlah_pengurus, b.jumlah_binaan_dalam, b.jumlah_binaan_luar, b.lampiran_ktp_ketua, b.lampiran_ktp_sekretaris, b.lampiran_ktp_bendahara, b.lampiran_akta_notaris, b.lampiran_kemenkumham, b.lampiran_adrt, b.lampiran_domisili, b.lampiran_akreditasi, b.lampiran_struktur_organisasi, b.lampiran_npwp, b.lampiran_foto_lokasi, b.lampiran_foto_usaha, b.lampiran_logo, b.lampiran_data_binaan From _permohonan_temp a INNER JOIN _permohonan_lksa b ON b.id_permohonan = a.id WHERE a.id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Nama_lembaga,
			&obj.Jenis_lembaga,
			&obj.Tgl_berdiri_lembaga,
			&obj.Nama_notaris_lembaga,
			&obj.Nomor_notaris_lembaga,
			&obj.Nomor_kemenkumham_lembaga,
			&obj.Akreditasi_lembaga,
			&obj.Nomor_surat_akreditasi_lembaga,
			&obj.Tgl_expired_akreditasi_lembaga,
			&obj.Npwp_lembaga,
			&obj.Modal_usaha_lembaga,
			&obj.Status_lembaga,
			&obj.Lingkup_wilayah_kerja_lembaga,
			&obj.Bidang_kegiatan_lembaga,
			&obj.No_telp_lembaga,
			&obj.Email_lembaga,
			&obj.Lat_long_lembaga,
			&obj.Alamat_lembaga,
			&obj.Rt_lembaga,
			&obj.Rw_lembaga,
			&obj.Kecamatan_lembaga,
			&obj.Kelurahan_lembaga,
			&obj.Nama_ketua_pengurus,
			&obj.Nik_ketua_pengurus,
			&obj.Nohp_ketua_pengurus,
			&obj.Nama_sekretaris_pengurus,
			&obj.Nik_sekretaris_pengurus,
			&obj.Nohp_sekretaris_pengurus,
			&obj.Nama_bendahara_pengurus,
			&obj.Nik_bendahara_pengurus,
			&obj.Nohp_bendahara_pengurus,
			&obj.Jumlah_pengurus,
			&obj.Jumlah_binaan_dalam,
			&obj.Jumlah_binaan_luar,
			&obj.Lampiran_ktp_ketua,
			&obj.Lampiran_ktp_sekretaris,
			&obj.Lampiran_ktp_bendahara,
			&obj.Lampiran_akta_notaris,
			&obj.Lampiran_kemenkumham,
			&obj.Lampiran_adrt,
			&obj.Lampiran_domisili,
			&obj.Lampiran_akreditasi,
			&obj.Lampiran_struktur_organisasi,
			&obj.Lampiran_npwp,
			&obj.Lampiran_foto_lokasi,
			&obj.Lampiran_foto_usaha,
			&obj.Lampiran_logo,
			&obj.Lampiran_data_binaan,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "4" {
		sqlStatement := "SELECT a.id, a.layanan, a.kode_permohonan, a.kelurahan, a.ttd, a.nik, a.nama, a.user_id, a.jenis, a.status_permohonan, a.created_at, a.updated_at, b.nama_lembaga, b.jenis_lembaga, b.tgl_berdiri_lembaga, b.nama_notaris_lembaga, b.nomor_notaris_lembaga, b.nomor_kemenkumham_lembaga, b.akreditasi_lembaga, b.nomor_surat_akreditasi_lembaga, b.tgl_expired_akreditasi_lembaga, b.npwp_lembaga, b.modal_usaha_lembaga, b.status_lembaga, b.lingkup_wilayah_kerja_lembaga, b.bidang_kegiatan_lembaga, b.no_telp_lembaga, b.email_lembaga, b.lat_long_lembaga, b.alamat_lembaga, b.rt_lembaga, b.rw_lembaga, (SELECT kecamatan FROM ref_kecamatan Where id = b.kecamatan_lembaga) as kecamatan_lembaga, (SELECT kelurahan FROM ref_kelurahan Where id = b.kelurahan_lembaga) as kelurahan_lembaga, b.nama_ketua_pengurus, b.nik_ketua_pengurus, b.nohp_ketua_pengurus, b.nama_sekretaris_pengurus, b.nik_sekretaris_pengurus, b.nohp_sekretaris_pengurus, b.nama_bendahara_pengurus, b.nik_bendahara_pengurus, b.nohp_bendahara_pengurus, b.jumlah_pengurus, b.jumlah_binaan_dalam, b.jumlah_binaan_luar, b.lampiran_ktp_ketua, b.lampiran_ktp_sekretaris, b.lampiran_ktp_bendahara, b.lampiran_akta_notaris, b.lampiran_kemenkumham, b.lampiran_adrt, b.lampiran_domisili, b.lampiran_akreditasi, b.lampiran_struktur_organisasi, b.lampiran_npwp, b.lampiran_foto_lokasi, b.lampiran_foto_usaha, b.lampiran_logo, b.lampiran_data_binaan From _permohonan_temp a INNER JOIN _permohonan_lksa b ON b.id_permohonan = a.id WHERE a.id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Nama_lembaga,
			&obj.Jenis_lembaga,
			&obj.Tgl_berdiri_lembaga,
			&obj.Nama_notaris_lembaga,
			&obj.Nomor_notaris_lembaga,
			&obj.Nomor_kemenkumham_lembaga,
			&obj.Akreditasi_lembaga,
			&obj.Nomor_surat_akreditasi_lembaga,
			&obj.Tgl_expired_akreditasi_lembaga,
			&obj.Npwp_lembaga,
			&obj.Modal_usaha_lembaga,
			&obj.Status_lembaga,
			&obj.Lingkup_wilayah_kerja_lembaga,
			&obj.Bidang_kegiatan_lembaga,
			&obj.No_telp_lembaga,
			&obj.Email_lembaga,
			&obj.Lat_long_lembaga,
			&obj.Alamat_lembaga,
			&obj.Rt_lembaga,
			&obj.Rw_lembaga,
			&obj.Kecamatan_lembaga,
			&obj.Kelurahan_lembaga,
			&obj.Nama_ketua_pengurus,
			&obj.Nik_ketua_pengurus,
			&obj.Nohp_ketua_pengurus,
			&obj.Nama_sekretaris_pengurus,
			&obj.Nik_sekretaris_pengurus,
			&obj.Nohp_sekretaris_pengurus,
			&obj.Nama_bendahara_pengurus,
			&obj.Nik_bendahara_pengurus,
			&obj.Nohp_bendahara_pengurus,
			&obj.Jumlah_pengurus,
			&obj.Jumlah_binaan_dalam,
			&obj.Jumlah_binaan_luar,
			&obj.Lampiran_ktp_ketua,
			&obj.Lampiran_ktp_sekretaris,
			&obj.Lampiran_ktp_bendahara,
			&obj.Lampiran_akta_notaris,
			&obj.Lampiran_kemenkumham,
			&obj.Lampiran_adrt,
			&obj.Lampiran_domisili,
			&obj.Lampiran_akreditasi,
			&obj.Lampiran_struktur_organisasi,
			&obj.Lampiran_npwp,
			&obj.Lampiran_foto_lokasi,
			&obj.Lampiran_foto_usaha,
			&obj.Lampiran_logo,
			&obj.Lampiran_data_binaan,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "5" {
		sqlStatement := "SELECT a.id, a.layanan, a.kode_permohonan, a.kelurahan, a.ttd, a.nik, a.nama, a.user_id, a.jenis, a.status_permohonan, a.created_at, a.updated_at, b.nama_lembaga, b.jenis_lembaga, b.tgl_berdiri_lembaga, b.nama_notaris_lembaga, b.nomor_notaris_lembaga, b.nomor_kemenkumham_lembaga, b.akreditasi_lembaga, b.nomor_surat_akreditasi_lembaga, b.tgl_expired_akreditasi_lembaga, b.npwp_lembaga, b.modal_usaha_lembaga, b.status_lembaga, b.lingkup_wilayah_kerja_lembaga, b.bidang_kegiatan_lembaga, b.no_telp_lembaga, b.email_lembaga, b.lat_long_lembaga, b.alamat_lembaga, b.rt_lembaga, b.rw_lembaga, (SELECT kecamatan FROM ref_kecamatan Where id = b.kecamatan_lembaga) as kecamatan_lembaga, (SELECT kelurahan FROM ref_kelurahan Where id = b.kelurahan_lembaga) as kelurahan_lembaga, b.nama_ketua_pengurus, b.nik_ketua_pengurus, b.nohp_ketua_pengurus, b.nama_sekretaris_pengurus, b.nik_sekretaris_pengurus, b.nohp_sekretaris_pengurus, b.nama_bendahara_pengurus, b.nik_bendahara_pengurus, b.nohp_bendahara_pengurus, b.jumlah_pengurus, b.jumlah_binaan_dalam, b.jumlah_binaan_luar, b.lampiran_ktp_ketua, b.lampiran_ktp_sekretaris, b.lampiran_ktp_bendahara, b.lampiran_akta_notaris, b.lampiran_kemenkumham, b.lampiran_adrt, b.lampiran_domisili, b.lampiran_akreditasi, b.lampiran_struktur_organisasi, b.lampiran_npwp, b.lampiran_foto_lokasi, b.lampiran_foto_usaha, b.lampiran_logo, b.lampiran_data_binaan From _permohonan_temp a INNER JOIN _permohonan_lksa b ON b.id_permohonan = a.id WHERE a.id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
			&obj.Nama_lembaga,
			&obj.Jenis_lembaga,
			&obj.Tgl_berdiri_lembaga,
			&obj.Nama_notaris_lembaga,
			&obj.Nomor_notaris_lembaga,
			&obj.Nomor_kemenkumham_lembaga,
			&obj.Akreditasi_lembaga,
			&obj.Nomor_surat_akreditasi_lembaga,
			&obj.Tgl_expired_akreditasi_lembaga,
			&obj.Npwp_lembaga,
			&obj.Modal_usaha_lembaga,
			&obj.Status_lembaga,
			&obj.Lingkup_wilayah_kerja_lembaga,
			&obj.Bidang_kegiatan_lembaga,
			&obj.No_telp_lembaga,
			&obj.Email_lembaga,
			&obj.Lat_long_lembaga,
			&obj.Alamat_lembaga,
			&obj.Rt_lembaga,
			&obj.Rw_lembaga,
			&obj.Kecamatan_lembaga,
			&obj.Kelurahan_lembaga,
			&obj.Nama_ketua_pengurus,
			&obj.Nik_ketua_pengurus,
			&obj.Nohp_ketua_pengurus,
			&obj.Nama_sekretaris_pengurus,
			&obj.Nik_sekretaris_pengurus,
			&obj.Nohp_sekretaris_pengurus,
			&obj.Nama_bendahara_pengurus,
			&obj.Nik_bendahara_pengurus,
			&obj.Nohp_bendahara_pengurus,
			&obj.Jumlah_pengurus,
			&obj.Jumlah_binaan_dalam,
			&obj.Jumlah_binaan_luar,
			&obj.Lampiran_ktp_ketua,
			&obj.Lampiran_ktp_sekretaris,
			&obj.Lampiran_ktp_bendahara,
			&obj.Lampiran_akta_notaris,
			&obj.Lampiran_kemenkumham,
			&obj.Lampiran_adrt,
			&obj.Lampiran_domisili,
			&obj.Lampiran_akreditasi,
			&obj.Lampiran_struktur_organisasi,
			&obj.Lampiran_npwp,
			&obj.Lampiran_foto_lokasi,
			&obj.Lampiran_foto_usaha,
			&obj.Lampiran_logo,
			&obj.Lampiran_data_binaan,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	}

	return obj, nil
}

func GetDetailPbi(id, status string) (Pbi, error) {
	var obj Pbi

	con := db.CreateCon()

	if status == "0" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, status_permohonan, created_at, updated_at From _permohonan_temp WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "1" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_approve,
			&obj.Date_approve,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "2" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_approve,
			&obj.Date_approve,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "3" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan_tolak WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "4" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "5" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	}

	return obj, nil
}

func GetDetailSkdtks(id, status string) (Skdtks, error) {
	var obj Skdtks

	con := db.CreateCon()

	if status == "0" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, status_permohonan, created_at, updated_at From _permohonan_temp WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "1" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_approve,
			&obj.Date_approve,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "2" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_approve,
			&obj.Date_approve,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "3" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan_tolak WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "4" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "5" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	}

	return obj, nil
}

func GetDetailSktm(id, status string) (Sktm, error) {
	var obj Sktm

	con := db.CreateCon()

	if status == "0" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, status_permohonan, created_at, updated_at From _permohonan_temp WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "1" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_approve,
			&obj.Date_approve,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "2" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_approve, date_approve, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_approve,
			&obj.Date_approve,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "3" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan_tolak WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "4" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	} else if status == "5" {
		sqlStatement := "SELECT id, layanan, kode_permohonan, kelurahan, ttd, nik, nama, user_id, jenis, indikator1, indikator2, indikator3, indikator4, indikator5, indikator6, skor, lampiran_ktp, lampiran_kk, lampiran_pernyataan, lampiran_foto_rumah, lampiran_lainnya, admin_reject, date_reject, keterangan_reject, status_permohonan, created_at, updated_at From _permohonan WHERE id = ?"

		err := con.QueryRow(sqlStatement, id).Scan(
			&obj.Id,
			&obj.Layanan,
			&obj.Kode_permohonan,
			&obj.Kelurahan,
			&obj.Ttd,
			&obj.Nik,
			&obj.Nama,
			&obj.User_id,
			&obj.Jenis,
			&obj.Indikator1,
			&obj.Indikator2,
			&obj.Indikator3,
			&obj.Indikator4,
			&obj.Indikator5,
			&obj.Indikator6,
			&obj.Skor,
			&obj.Lampiran_ktp,
			&obj.Lampiran_kk,
			&obj.Lampiran_pernyataan,
			&obj.Lampiran_foto_rumah,
			&obj.Lampiran_lainnya,
			&obj.Admin_reject,
			&obj.Date_reject,
			&obj.Keterangan_reject,
			&obj.Status_permohonan,
			&obj.Created_at,
			&obj.Updated_at,
		)

		if err == sql.ErrNoRows {
			// fmt.Println("User not found")
			return obj, errors.New("data tidak ditemukan")
		}

		if err != nil {
			// fmt.Println("Query error")
			log.Println(err.Error())
			return obj, errors.New("gagal mengambil data")
		}
	}

	return obj, nil
}
