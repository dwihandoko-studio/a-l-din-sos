package routes

import (
	"api-dinsos/controllers"
	"api-dinsos/middleware"
	"net/http"

	middlewareEcho "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middlewareEcho.Logger())
	e.Use(middlewareEcho.Recover())

	e.Use(middlewareEcho.CORSWithConfig(middlewareEcho.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		MaxAge:           300,
		// AllowOrigins: []string{"http://localhost:3000", "http://10.20.30.99:3000"},
		AllowHeaders: []string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-API-TOKEN, X-CSRF-Token, Authorization, Access-Control-Allow-Origin"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
	}))

	e.Static("/files", "uploads")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Bagaimana Kabar Anda!")
	})

	e.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "Akses tidak diizinkan.")
	})
	e.GET("/register", func(c echo.Context) error {
		return c.String(http.StatusNotFound, "Akses tidak diizinkan.")
	})

	e.GET("/slider", controllers.GetSliderBanner)
	e.GET("/referensi/provinsi", controllers.GetProvinsi)
	e.GET("/referensi/kabupaten/:id", controllers.GetKabupaten)
	e.GET("/referensi/kecamatan/:id", controllers.GetKecamatan)
	e.GET("/referensi/kelurahan/:id", controllers.GetKelurahan)

	e.POST("/login", controllers.CheckLogin)
	e.POST("/register", controllers.PostRegister)
	e.POST("/lupapassword/:aksi", controllers.PostLupapassword)

	e.POST("/verifigtoken", controllers.CheckGoogleToken)
	e.GET("/refreshtoken", controllers.RefreshToken, middleware.IsAuthenticated)
	e.GET("/user", controllers.GetUser, middleware.IsAuthenticated)
	e.GET("/user/:id", controllers.GetUserDetail, middleware.IsAuthenticated)
	e.POST("/users", controllers.GetUsers, middleware.IsAuthenticated)
	e.GET("/versionapp", controllers.GetVersionApp)
	e.POST("/changepassword", controllers.PostChangePassword, middleware.IsAuthenticated)
	e.POST("/changefoto", controllers.EditFotoUser, middleware.IsAuthenticated)

	e.POST("/layanan/dtks", controllers.PostLayananDtks, middleware.IsAuthenticated)
	e.POST("/layanan/sktm", controllers.PostLayananSktm, middleware.IsAuthenticated)
	e.POST("/layanan/pbi", controllers.PostLayananPbi, middleware.IsAuthenticated)
	e.POST("/layanan/lks", controllers.PostLayananLks, middleware.IsAuthenticated)

	e.POST("/pengaduan", controllers.PostPengaduan, middleware.IsAuthenticated)

	e.POST("/user", controllers.PostUser, middleware.IsAuthenticated)

	e.GET("/riwayatpermohonan", controllers.GetRiwayatLayanan, middleware.IsAuthenticated)
	e.GET("/riwayatpermohonan/:id/:layanan/:status", controllers.GetDetailRiwayatLayanan, middleware.IsAuthenticated)

	e.GET("/riwayataktifitas", controllers.GetAktifitas, middleware.IsAuthenticated)

	e.GET("/riwayatpengaduan", controllers.GetRiwayatPengaduan, middleware.IsAuthenticated)
	e.GET("/riwayatpengaduan/:id/:status", controllers.GetDetailRiwayatPengaduan, middleware.IsAuthenticated)

	// e.POST("/daftar", controllers.CheckDataSiswa)

	// r := e.Group("/provinsi")

	// // Configure middleware with the custom claims type
	// config := middlewareEcho.JWTConfig{
	// 	Claims:     &jwtCustomClaims{},
	// 	SigningKey: []byte("secret key handokowae.my.id"),
	// }
	// r.Use(middlewareEcho.JWTWithConfig(config))
	// r.GET("", restricted)

	return e
}
