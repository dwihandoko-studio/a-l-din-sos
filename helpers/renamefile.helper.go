package helpers

import (
	"strconv"
	"time"
)

func GenerateFilename(filename string) string {
	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)

	newfilename := filename + "-" + currentTime.Format("2006-01-02") + "-" + strconv.Itoa(int(currentTime.UnixMilli()))

	return newfilename
}

func GenerateKodePermohonan(nik string) string {
	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)

	newfilename := nik + "-" + currentTime.Format("2006-01-02") + "-" + strconv.Itoa(int(currentTime.UnixMilli()))

	return newfilename
}
