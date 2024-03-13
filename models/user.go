package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string		`json:"username" gorm:"not null,varchar(191)" validate:"unique,required"`
	Email		string		`json:"email" gorm:"not null,varchar(191)" validate:"email,unique,required"`
	Password	string		`json:"password" gorm:"not null,varchar(191)" validate:"required,min=6"`
	Age			uint		`json:"age" gorm:"not null" validate:"required,min=8"`
}

type UsersInterface interface {
	Login() User
	Register() User
}

type UserImpl struct {
	Db *gorm.DB
}

func UsersModel(Db *gorm.DB) UsersInterface {
	return &UserImpl{Db: Db}
}

func (u *UserImpl) Login() User {
	var user User

	return user
}

func (u *UserImpl) Register() User {
	var user User
	return user
}