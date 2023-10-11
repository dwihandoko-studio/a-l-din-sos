package main

import (
	// "api-login/db"
	"api-dinsos/db"
	"api-dinsos/routes"
	"os"
	// "github.com/joho/godotenv"
)

func main() {
	portApp := os.Getenv("APP_PORT")

	if portApp == "" {
		portApp = "1990"
		// err := godotenv.Load("local.env")
		// if err != nil {
		// 	panic("connectionStringGorm error..." + err.Error())
		// }
	}

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + portApp))
}
