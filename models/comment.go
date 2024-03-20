package models

import (
	"errors"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId	uint	`json:"user_id"`
	PhotoId	uint	`json:"photo_id"`
	Message	string	`json:"message" validate:"required"`
	User	User	`json:"User" gorm:"foreignKey:ID;references:UserId"`
	Photo	Photo	`json:"Photo" gorm:"foreignKey:ID;references:PhotoId"`
}

type CommentsInterface interface {
	Get() ([]Comment, error)
	Post(uint, request.CommentRequest) (Comment, error)
}
type CommentImpl struct {
	Db *gorm.DB
}

func CommentsModel(Db *gorm.DB) CommentsInterface {
	Db.AutoMigrate(&Comment{})
	return &CommentImpl{Db: Db}
}

func (c *CommentImpl) Get() ([]Comment, error) {
	var comments []Comment
	if err := c.Db.Preload("User").Preload("Photo").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (c *CommentImpl) Post(userId uint, data request.CommentRequest) (Comment, error) {
	if err := c.Db.First(&Photo{}, data.PhotoId).Error; err != nil {
		return Comment{}, errors.New("photo not found")
	}
	comment := Comment{
		UserId: userId,
		PhotoId: data.PhotoId,
		Message: data.Message,
	}
	if err := c.Db.Create(&comment).Error; err != nil {
		return Comment{}, nil
	}
	return comment, nil
}