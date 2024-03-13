package models

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name	string	`json:"name" gorm:"not null" validate:"required"`
	Url		string	`json:"social_media_url" gorm:"not null" validate:"required"`
	UserId	string	`json:"user_id"`
	User	User	`json:"User" gorm:"foreignKey:ID;references:UserId"`
}