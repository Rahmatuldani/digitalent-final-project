package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title		string	`json:"title" gorm:"not null,varchar(191)" validate:"required"`
	Caption		string	`json:"caption" gorm:"varchar(191)"`
	PhotoUrl	string	`json:"photo_url" gorm:"not null,varchar(191)" validate:"required"`
	UserId		uint	`json:"user_id"`
	User		User	`json:"User" gorm:"foreignKey:ID;references:UserId"`
}