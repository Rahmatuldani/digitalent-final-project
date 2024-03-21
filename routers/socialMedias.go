package routers

import (
	"github.com/Rahmatuldani/digitalent-project/controllers"
	"github.com/Rahmatuldani/digitalent-project/middlewares"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SocialMediaRoutes(db *gorm.DB, v *validator.Validate, r *gin.RouterGroup) {
	model := models.SocialMediaModel(db)
	controller := controllers.SocialMediaController(model, v)

	basePath := r.Group("/socialmedias")
	{
		basePath.Use(middlewares.Authentication)
		basePath.GET("/", controller.GetSocialMedia)
		basePath.POST("/", controller.PostSocialMedia)
		basePath.PUT("/:id", controller.UpdateSocialMedia)
		basePath.DELETE("/:id", controller.DeleteSocialMedia)
	}
}