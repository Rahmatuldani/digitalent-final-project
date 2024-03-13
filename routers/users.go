package routers

import (
	"github.com/Rahmatuldani/digitalent-project/controllers"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, v *validator.Validate, r *gin.RouterGroup) {
	model := models.UsersModel(db)
	controller := controllers.UsersController(model, v)

	basePath := r.Group("/users")
	{
		basePath.POST("/login", controller.Login)
		basePath.POST("/register", controller.Register)
	}
}