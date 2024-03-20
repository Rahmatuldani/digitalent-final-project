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
	Update(uint, uint, request.CommentUpdateReq) (Comment, error)
	Delete(uint, uint) error
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
		return Comment{}, err
	}
	return comment, nil
}

func (c *CommentImpl) Update(userId, id uint, data request.CommentUpdateReq) (Comment, error) {
	var comment Comment
	if err := c.Db.First(&comment, id).Error; err != nil {
		return Comment{}, errors.New("comment not found")
	}
	if comment.UserId != userId {
		return Comment{}, errors.New("can't update comment that aren't yours")
	}
	comment.Message = data.Message
	if err := c.Db.Save(&comment).Error; err != nil {
		return Comment{}, err
	}
	return comment, nil
}

func (c *CommentImpl) Delete(userId, id uint) error {
	var comment Comment
	if err := c.Db.First(&comment, id).Error; err != nil {
		return errors.New("comment not found")
	}
	if comment.UserId != userId {
		return errors.New("can't delete comment that aren't yours")
	}
	if err := c.Db.Unscoped().Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}