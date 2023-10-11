package models

import (
	"api-dinsos/db"
	"api-dinsos/services"
	"errors"
	"fmt"
	"time"
)

type Tokenactivation struct {
	Token      *string `json:"token"`
	Email      *string `json:"email"`
	Code       *string `json:"code"`
	Created_at *string `json:"created_at"`
}

func InsertToken(token, email string) (Tokenactivation, error) {
	con := db.CreateCon()

	tx, _ := con.Begin()

	loca, _ := time.LoadLocation("Asia/Jakarta")

	currentTime := time.Now().In(loca)
	code := services.GenerateRandomNumber(4)

	tok := Tokenactivation{
		Token: &token,
		Email: &email,
		Code:  &code,
	}
	sqlStatementInsert := `INSERT INTO _token_reset(token, email, code, created_at) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE email = VALUES(email), code = VALUES(code), created_at = VALUES(created_at)`

	exeInsertDoc, err := tx.Exec(sqlStatementInsert, &tok.Token, &tok.Email, &tok.Code, currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		// fmt.Println("Query error")
		fmt.Println(err.Error())
		tx.Rollback()
		return tok, errors.New("gagal insert token reset")
	}

	rowsInsertDoc, err := exeInsertDoc.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return tok, errors.New("gagal insert token reset")
	}

	if rowsInsertDoc > 0 {
		tx.Commit()
		return tok, nil

	}
	tx.Rollback()
	return tok, errors.New("gagal menyimpan token reset")
}
