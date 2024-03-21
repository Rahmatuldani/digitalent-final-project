package models

import (
	"errors"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"gorm.io/gorm"
)

type SocialMedia struct {
	gorm.Model
	Name	string	`json:"name" gorm:"not null" validate:"required"`
	Url		string	`json:"social_media_url" gorm:"not null" validate:"required"`
	UserId	uint	`json:"user_id"`
	User	User	`json:"User" gorm:"foreignKey:ID;references:UserId"`
}

type SocialMediaInterface interface {
	Get() ([]SocialMedia, error)
	Post(uint, request.SocialMediaReq) (SocialMedia, error)
	Update(uint, uint, request.SocialMediaReq) (SocialMedia, error)
	Delete(uint, uint) error
}

type SocialMediaStruct struct {
	Db *gorm.DB
}

func SocialMediaModel(Db *gorm.DB) SocialMediaInterface {
	Db.AutoMigrate(&SocialMedia{})
	return &SocialMediaStruct{Db: Db}
}

func (s *SocialMediaStruct) Get() ([]SocialMedia, error) {
	var sosmed []SocialMedia
	if err := s.Db.Preload("User").Find(&sosmed).Error; err != nil {
		return nil, err
	}
	return sosmed, nil
}

func (s *SocialMediaStruct) Post(userId uint, data request.SocialMediaReq) (SocialMedia, error) {
	socialMedia := SocialMedia{
		Name: data.Name,
		Url: data.SocialMediaUrl,
		UserId: userId,
	}
	if err := s.Db.Create(&socialMedia).Error; err != nil {
		return SocialMedia{}, err
	}
	return socialMedia, nil
}

func (s *SocialMediaStruct) Update(userId, id uint, data request.SocialMediaReq) (SocialMedia, error) {
	var socialMedia SocialMedia
	if err := s.Db.First(&socialMedia, id).Error; err != nil {
		return SocialMedia{}, err
	}
	if socialMedia.UserId != userId {
		return SocialMedia{}, errors.New("can't update social media that aren't yours")
	}
	socialMedia.Name = data.Name
	socialMedia.Url = data.SocialMediaUrl
	s.Db.Save(&socialMedia)
	return socialMedia, nil
}

func (s *SocialMediaStruct) Delete(userId, id uint) error {
	var socialMedia SocialMedia
	if err := s.Db.First(&socialMedia, id).Error; err != nil {
		return err
	}
	if socialMedia.UserId != userId {
		return errors.New("can't delete social media that aren't yours")
	}
	if err := s.Db.Unscoped().Delete(&socialMedia).Error; err != nil {
		return err
	}
	return nil
}