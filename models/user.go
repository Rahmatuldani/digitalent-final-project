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
	Update(uint, uint, request.UserUpdateReq) (User, error)
	Delete(uint) error
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
	if err := u.Db.Where("email = ?", data.Email).First(&user).Error; err != nil {
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
	if err = u.Db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserImpl) Update(userId, id uint, data request.UserUpdateReq) (User, error) {
	var user User
	if err := u.Db.First(&user, id).Error; err != nil {
		return User{}, errors.New("user not found")
	}
	if user.ID != userId {
		return User{}, errors.New("can't update user data that aren't yours")
	}
	user.Email = data.Email
	user.Username = data.Username
	if err := u.Db.Save(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserImpl) Delete(id uint) error {
	if err := u.Db.First(&User{}, id).Error; err != nil {
		return errors.New("user not found")
	}
	if err := u.Db.Unscoped().Select(clause.Associations).Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
} 

func (u *UserImpl) CheckUser(id uint8) bool {
	var user User
	err := u.Db.Where("id = ?", id).First(&user).Error
	return err == nil
}