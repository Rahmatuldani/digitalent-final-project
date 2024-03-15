package models

import (
	"github.com/Rahmatuldani/digitalent-project/data/request"
	"github.com/Rahmatuldani/digitalent-project/helper"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string		`json:"username" gorm:"not null;varchar(191);unique" validate:"unique,required"`
	Email		string		`json:"email" gorm:"not null;varchar(191);unique" validate:"email,unique,required"`
	Password	string		`json:"password" gorm:"not null;varchar(191)" validate:"required,min=6"`
	Age			uint8		`json:"age" gorm:"not null" validate:"required,min=8"`
}

type UsersInterface interface {
	Login() User
	Register(data request.UserRegReq) (User, error)
}

type UserImpl struct {
	Db *gorm.DB
}

func UsersModel(Db *gorm.DB) UsersInterface {
	Db.AutoMigrate(&User{})
	return &UserImpl{Db: Db}
}

func (u *UserImpl) Login() User {
	var user User

	return user
}

func (u *UserImpl) Register(data request.UserRegReq) (User, error) {
	password, err := helper.Encrypt(data.Password)
	if err != nil {
		return User{}, err
	}
	user := User{
		Username: data.Username,
		Email: data.Email,
		Password: password,
		Age: data.Age,
	}
	err = u.Db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}