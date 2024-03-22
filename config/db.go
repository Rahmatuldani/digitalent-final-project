package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

var (
	db *gorm.DB
	err error
)

func DBConnect() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=true"
	fmt.Println(dsn)
	// dsn := "root:podyQsHisuzhRaTyILQUQPQvzrwMKICA@tcp(roundhouse.proxy.rlwy.net:52387)/railway?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection Error : " + err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
