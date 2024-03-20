package routers

import (
	"github.com/Rahmatuldani/digitalent-project/controllers"
	"github.com/Rahmatuldani/digitalent-project/middlewares"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func PhotoRoutes(db *gorm.DB, v *validator.Validate, r *gin.RouterGroup) {
	model := models.PhotosModel(db)
	controller := controllers.PhotosController(model, v)

	basePath := r.Group("/photos")
	{
		basePath.Use(middlewares.Authentication)
		basePath.GET("/", controller.GetPhotos)
		basePath.POST("/", controller.PostPhoto)
		basePath.PUT("/:id", controller.UpdatePhoto)
		basePath.DELETE("/:id", controller.DeletePhoto)
	}
}