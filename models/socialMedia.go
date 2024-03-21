package models

import (
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
	Post(uint, request.PostSocialMediaReq) (SocialMedia, error)
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

func (s *SocialMediaStruct) Post(userId uint, data request.PostSocialMediaReq) (SocialMedia, error) {
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