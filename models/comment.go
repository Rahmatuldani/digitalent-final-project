package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId	uint	`json:"user_id"`
	PhotoId	uint	`json:"photo_id"`
	Message	string	`json:"message" validate:"required"`
	User	User	`json:"User" gorm:"foreignKey:ID;references:UserId"`
	Photo	Photo	`json:"Photo" gorm:"foreignKey:ID;references:PhotoId"`
}