package models

import (
	"errors"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"gorm.io/gorm"
)

type Photo struct {
	gorm.Model
	Title		string	`json:"title" gorm:"not null;varchar(191)"`
	Caption		string	`json:"caption" gorm:"varchar(191)"`
	PhotoUrl	string	`json:"photo_url" gorm:"not null;varchar(191)"`
	UserId		uint	`json:"user_id"`
	User		User	`json:"User" gorm:"foreignKey:ID;references:UserId"`
}

type PhotosInterface interface {
	Get() ([]Photo, error)
	Post(uint, request.PhotoReq) (Photo, error)
	Update(uint, uint, request.PhotoReq) (Photo, error)
	Delete(uint, uint) error
}

type PhotoImpl struct {
	Db *gorm.DB
}

func PhotosModel(Db *gorm.DB) PhotosInterface {
	Db.AutoMigrate(&Photo{})
	return &PhotoImpl{Db: Db}
}

func (p *PhotoImpl) Get() ([]Photo, error) {
	var photos []Photo
	if err := p.Db.Preload("User").Find(&photos).Error; err != nil {
		return nil, err
	}
	return photos, nil
}

func (p *PhotoImpl) Post(id uint, data request.PhotoReq) (Photo, error) {
	photo := Photo{
		Title: data.Title,
		Caption: data.Caption,
		PhotoUrl: data.PhotoUrl,
		UserId: id,
	}
	if err := p.Db.Create(&photo).Error; err != nil {
		return Photo{}, err
	}
	return photo, nil
}

func (p *PhotoImpl) Update(userId, id uint, data request.PhotoReq) (Photo, error) {
	var photo Photo
	if err := p.Db.First(&photo, id).Error; err != nil {
		return Photo{}, errors.New("photo not found")
	}
	if userId != photo.UserId {
		return Photo{}, errors.New("can't update photos that aren't yours")
	}
	photo.Title = data.Title
	photo.Caption = data.Caption
	photo.PhotoUrl = data.PhotoUrl
	if err := p.Db.Save(&photo).Error; err != nil {
		return Photo{}, nil
	}
	return photo, nil
}

func (p *PhotoImpl) Delete(userId, id uint) error {
	var photo Photo
	if err := p.Db.First(&photo, id).Error; err != nil {
		return errors.New("photo not found")
	}
	if userId != photo.UserId {
		return errors.New("can't delete photos that aren't yours")
	}
	if err := p.Db.Unscoped().Delete(&Photo{}, id).Error; err != nil {
		return err
	}
	return nil
}