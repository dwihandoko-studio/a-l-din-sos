package db

import (
	"database/sql"
	"os"

	// _ "github.com/lib/pq"
	// "gorm.io/driver/postgres"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB
var err error

func Init() {
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DATABASE")
	// host := os.Getenv("DB_HOST")
	// port := os.Getenv("DB_PORT")
	// user := os.Getenv("DB_USERNAME")
	// password := os.Getenv("DB_PASSWORD")
	// dbname := os.Getenv("DB_NAME")

	connectionString := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error...")
	}

	// psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err = sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	panic("connectionString error...")
	// }

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}

func CreateConGorm() *gorm.DB {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})
	// gormDB, err := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: db,
	// }), &gorm.Config{})
	if err != nil {
		panic("connectionStringGorm error...")
	}
	return gormDB
}
