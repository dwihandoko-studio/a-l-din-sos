package controllers

import (
	"api-dinsos/helpers"
	"api-dinsos/middleware"
	"api-dinsos/models"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

type jwtCustomClaims struct {
	Id    string `json:"id"`
	Level int    `json:"level"`
	jwt.StandardClaims
}

// var (
// 	googleOauthConfig = &oauth2.config{
// 		RedirectURL: "http://localhost:8080/callback",
// 		ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
// 		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
// 		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
// 		Endpoint: google.Endpoint,
// 	}
// )

func PostRegister(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")
	// username := c.Request().Form.Get("username")
	// fmt.Println(keyGet)
	// fmt.Println("RUNNING")
	var nama string
	var nik string
	var no_hp string
	var email string
	var password string
	var re_password string

	// c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	// c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	// fmt.Println(c.FormValue("username"))

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		// fmt.Println("PARSING REQUEST FORM")
		nama = c.FormValue("nama")
		nik = c.FormValue("nik")
		no_hp = c.FormValue("no_hp")
		email = c.FormValue("email")
		password = c.FormValue("password")
		re_password = c.FormValue("re_password")
	} else {
		//json_map has the JSON Payload decoded into a map
		// fmt.Println("PARSING REQUEST JSON")
		nama = fmt.Sprintf("%s", json_map["nama"])
		nik = fmt.Sprintf("%s", json_map["nik"])
		no_hp = fmt.Sprintf("%s", json_map["no_hp"])
		email = fmt.Sprintf("%s", json_map["email"])
		password = fmt.Sprintf("%s", json_map["password"])
		re_password = fmt.Sprintf("%s", json_map["re_password"])
	}

	if nama == "" || nik == "" || no_hp == "" || email == "" || password == "" || re_password == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Nama, NIK, No HP, Email dan Password tidak boleh kosong."},
		)
	}
	log.Println(email)

	res, err := models.CheckUserRegistrasiEmail(email)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	if res.Status != http.StatusOK {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Email sudah terdaftar."},
		)
	}

	res1, err := models.CheckUserRegistrasiNik(nik)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	if res1.Status != http.StatusOK {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "NIK sudah terdaftar."},
		)
	}

	res2, err := models.CheckUserRegistrasiNohp(no_hp)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	if res2.Status != http.StatusOK {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "No Handphone sudah terdaftar."},
		)
	}

	var uuid = uuid.New()

	hash, _ := helpers.HashPassword(password)

	loca, _ := time.LoadLocation("Asia/Jakarta")
	currentTime := time.Now().In(loca)
	var emailLower = strings.ToLower(email)

	var user = models.Userregister{
		Id:                uuid.String(),
		Email:             emailLower,
		Nik:               nik,
		No_hp:             &no_hp,
		Password:          &hash,
		Scope:             string("app"),
		Is_active:         1,
		Email_verified:    0,
		Wa_verified:       0,
		Created_at:        currentTime.Format("2006-01-02 15:04:05"),
		Updated_at:        nil,
		Update_firs_login: nil,
	}

	var profil = models.Userprofilregister{
		Id:         user.Id,
		Fullname:   &nama,
		Email:      &emailLower,
		Nik:        &nik,
		No_hp:      &no_hp,
		Role_user:  2,
		Created_at: &user.Created_at,
	}

	result, err := models.SaveRegistered(user, profil)
	// result, err := models.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	if result.Status != 200 {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: result.Message},
		)
	}

	timeExpired := time.Now().In(loca).Add(time.Hour * 24 * 1).Unix()

	claims := &jwtCustomClaims{
		user.Id,
		profil.Role_user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().In(loca).Add(time.Hour * 24 * 1).Unix(),
		},
	}

	claimsRefresh := &jwtCustomClaims{
		user.Id,
		profil.Role_user,
		jwt.StandardClaims{
			ExpiresAt: timeExpired,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	t, err := token.SignedString([]byte("secret key handokowae.my.id"))
	tRefresh, errRefresh := tokenRefresh.SignedString([]byte("secret key handokowae.my.id"))
	if err != nil {
		// log.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Pendaftaran gagal."},
		)
	}

	if errRefresh != nil {
		// fmt.Println(errRefresh.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Pendaftaran gagal."},
		)
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: res.Message,
			Data: models.Token{
				Access:     &t,
				Expired_in: timeExpired,
				Refresh:    &tRefresh,
			},
		},
	)
}

func CheckLogin(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")
	formValues := c.Request().Form
	log.Println("All Form Values:", formValues)
	// username := c.Request().Form.Get("username")
	// fmt.Println(keyGet)
	// fmt.Println("RUNNING")
	var username string
	var password string

	// c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	// c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	// fmt.Println(c.FormValue("username"))

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		log.Println("PARSING REQUEST FORM")
		username = c.FormValue("username")
		password = c.FormValue("password")
	} else {
		//json_map has the JSON Payload decoded into a map
		log.Println("PARSING REQUEST JSON")
		username = fmt.Sprintf("%s", json_map["username"])
		password = fmt.Sprintf("%s", json_map["password"])
	}

	if username == "" || password == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Username dan Password tidak boleh kosong."},
		)
	}
	log.Println(username)

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()},
		)
	}

	if res.Status != http.StatusOK {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Login gagal."},
		)
	}
	loca, _ := time.LoadLocation("Asia/Jakarta")

	user := res.Data.(models.User)

	timeExpired := time.Now().In(loca).Add(time.Hour * 24 * 1).Unix()

	userLevel, _ := strconv.Atoi(user.Level)

	claims := &jwtCustomClaims{
		user.Id,
		userLevel,
		jwt.StandardClaims{
			ExpiresAt: time.Now().In(loca).Add(time.Hour * 24 * 1).Unix(),
		},
	}

	claimsRefresh := &jwtCustomClaims{
		user.Id,
		userLevel,
		jwt.StandardClaims{
			ExpiresAt: timeExpired,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	t, err := token.SignedString([]byte("secret key handokowae.my.id"))
	tRefresh, errRefresh := tokenRefresh.SignedString([]byte("secret key handokowae.my.id"))
	if err != nil {
		// log.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Login gagal."},
		)
	}

	if errRefresh != nil {
		// fmt.Println(errRefresh.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Login gagal."},
		)
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: res.Message,
			Data: models.Token{
				Access:     &t,
				Expired_in: timeExpired,
				Refresh:    &tRefresh,
			},
		},
	)
}

func PostLupapassword(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")
	aksi := c.Param("aksi")
	var email string

	// c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	// c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	switch aksi {
	case "email":
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)

		if err != nil {
			// fmt.Println("PARSING REQUEST FORM")
			email = c.FormValue("email")
		} else {
			//json_map has the JSON Payload decoded into a map
			// fmt.Println("PARSING REQUEST JSON")
			email = fmt.Sprintf("%s", json_map["email"])
		}

		if email == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Email tidak valid."},
			)
		}
		// log.Println(email)

		res, err := models.SendResetUsingEmail(email)
		if err != nil {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		if res.Status != http.StatusOK {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Reset password gagal dikirim."},
			)
		}

		log.Println("SUKSES MENGIRIM EMAIL")

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: res.Message,
			},
		)
	case "nohp":
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)

		if err != nil {
			// fmt.Println("PARSING REQUEST FORM")
			email = c.FormValue("nohp")
		} else {
			//json_map has the JSON Payload decoded into a map
			// fmt.Println("PARSING REQUEST JSON")
			email = fmt.Sprintf("%s", json_map["nohp"])
		}

		if email == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "No handphone tidak valid."},
			)
		}
		log.Println(email)

		res, err := models.SendResetUsingNohp(email)
		if err != nil {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		if res.Status != http.StatusOK {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Login gagal."},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: res.Message,
			},
		)
	default:
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)

		if err != nil {
			// fmt.Println("PARSING REQUEST FORM")
			email = c.FormValue("email")
		} else {
			//json_map has the JSON Payload decoded into a map
			// fmt.Println("PARSING REQUEST JSON")
			email = fmt.Sprintf("%s", json_map["email"])
		}

		if email == "" {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Email tidak valid."},
			)
		}
		log.Println(email)

		res, err := models.SendResetUsingEmail(email)
		if err != nil {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: err.Error()},
			)
		}

		if res.Status != http.StatusOK {
			return c.JSON(http.StatusOK,
				models.Response{Status: 404, Message: "Login gagal."},
			)
		}

		return c.JSON(http.StatusOK,
			models.Response{
				Status:  200,
				Message: res.Message,
			},
		)
	}
}

func RefreshToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middleware.JwtCustomClaims)
	id := claims.Id

	loca, _ := time.LoadLocation("Asia/Jakarta")
	timeExpired := time.Now().In(loca).Add(time.Hour * 24 * 7).Unix()

	claimsNew := &jwtCustomClaims{
		id,
		1,
		jwt.StandardClaims{
			ExpiresAt: time.Now().In(loca).Add(time.Hour * 24 * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsNew)

	t, err := token.SignedString([]byte("secret key handokowae.my.id"))

	if err != nil {
		// fmt.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Login gagal."},
		)
	}

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "Berhasil",
			Data: models.Token{
				Access:     &t,
				Expired_in: timeExpired,
			},
		},
	)
}

func GenerateFromPassword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckGoogleToken(c echo.Context) error {
	var tokenId string

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		// fmt.Println("ERROR GUYS")
		tokenId = c.FormValue("gtoken_id")
	} else {
		fmt.Println(c.FormValue("gtoken_id"))
		fmt.Println(c.Request().Body)
		//json_map has the JSON Payload decoded into a map
		tokenId = fmt.Sprintf("%s", json_map["gtoken_id"])
	}

	if tokenId == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: "Token google invalid."},
		)
	}

	// fmt.Println(fmt.Sprintf("Tokennya : %s", tokenId))
	payload, err := idtoken.Validate(context.Background(), tokenId, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: err.Error()},
		)
	}

	fmt.Println(payload.Claims)

	return c.JSON(http.StatusOK,
		models.Response{
			Status:  200,
			Message: "Token berhasil diverifikasi",
			Data:    payload.Claims,
		},
	)
}

func TakeDapoLocal(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")
	var npsn string

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	// json_map := make(map[string]interface{})
	// err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	// if err != nil {
	fmt.Println("PARSING REQUEST FORM")
	// data = c.FormValue("data")
	npsn = c.FormValue("npsn")
	// } else {
	// 	data = fmt.Sprintf("%s", json_map["data"])
	// 	npsn = fmt.Sprintf("%s", json_map["npsn"])
	// }

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "File upload tidak boleh kosong"},
		)
	}

	if string(filepath.Ext(file.Filename)) != ".json" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "Type file tidak dizinkan."},
		)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
		)
	}

	srckk, err := file.Open()
	if err != nil {
		return err
	}
	defer srckk.Close()

	filenamekk := fmt.Sprintf("%s%s", npsn, filepath.Ext(file.Filename))

	fileLocationkk := filepath.Join(dir, "uploads", filenamekk)
	targetFilekk, err := os.OpenFile(fileLocationkk, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
		)
	}
	defer targetFilekk.Close()

	// Copy
	if _, err = io.Copy(targetFilekk, srckk); err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusOK,
			models.Response{Status: 204, Message: "Kesalahan dalam mengupload file."},
		)
	}

	log.Println(npsn)

	return c.JSON(http.StatusOK,
		models.Response{Status: 200, Message: "BERHASIL"},
	)
}
