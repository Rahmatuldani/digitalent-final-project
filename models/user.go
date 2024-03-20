package models

import (
	"errors"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"github.com/Rahmatuldani/digitalent-project/helper"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	gorm.Model
	Username	string		`json:"username" gorm:"not null;type:varchar(191);unique"`
	Email		string		`json:"email" gorm:"not null;type:varchar(191);unique"`
	Password	string		`json:"password" gorm:"not null;type:varchar(191)"`
	Age			uint8		`json:"age" gorm:"not null"`
}

type UsersInterface interface {
	Login(request.UserLogin) (User, error)
	Register(request.UserRegReq) (User, error)
	Update(uint8, request.UserUpdateReq) (User, error)
	Delete(uint8) error
	CheckUser(uint8) bool
}

type UserImpl struct {
	Db *gorm.DB
}

func UsersModel(Db *gorm.DB) UsersInterface {
	Db.AutoMigrate(&User{})
	return &UserImpl{Db: Db}
}

func (u *UserImpl) Login(data request.UserLogin) (User, error) {
	var user User
	err := u.Db.Where("email = ?", data.Email).First(&user).Error
	if err != nil {
		return User{}, errors.New("user not found")
	}

	if !helper.ComparePassword(user.Password, data.Password) {
		return User{}, errors.New("password did not match")
	}

	return user, nil
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

func (u *UserImpl) Update(id uint8, data request.UserUpdateReq) (User, error) {
	var user User
	err := u.Db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	user.Email = data.Email
	user.Username = data.Username
	u.Db.Save(&user)
	return user, nil
}

func (u *UserImpl) Delete(id uint8) error {
	err := u.Db.First(&User{}, id).Error
	if err != nil {
		return err
	}
	// u.Db.Model(&Photo{}).Unscoped().Where("user_id = ?", id).Delete(&Photo{})
	err = u.Db.Unscoped().Select(clause.Associations).Delete(&User{}, id).Error
	if err != nil {
		return err
	}
	return nil
} 

func (u *UserImpl) CheckUser(id uint8) bool {
	var user User
	err := u.Db.Where("id = ?", id).First(&user).Error
	return err == nil
}