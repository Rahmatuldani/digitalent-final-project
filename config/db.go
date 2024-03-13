package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	host     = "127.0.0.1"
	port     = "3306"
	username = "root"
	password = "root"
	dbname   = "digitalent"
	db		*gorm.DB
	err		error
)

func DBConnect() {
	dsn := username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection Error : "+ err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}