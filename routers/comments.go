package routers

import (
	"github.com/Rahmatuldani/digitalent-project/controllers"
	"github.com/Rahmatuldani/digitalent-project/middlewares"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CommentRoutes(db *gorm.DB, v *validator.Validate, r *gin.RouterGroup) {
	model := models.CommentsModel(db)
	controller := controllers.CommentsController(model, v)

	basePath := r.Group("/comments")
	{
		basePath.Use(middlewares.Authentication)
		basePath.GET("/", controller.GetComment)
		basePath.POST("/", controller.PostComment)
	}
}