package middlewares

import (
	"github.com/Rahmatuldani/digitalent-project/config"
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
)

func CheckUser(ctx *gin.Context) {
	db := config.GetDB()
	id := ctx.MustGet("userId").(uint64)
	err := db.Model(&models.User{}).Where("id = ?", id).First(&models.User{}).Error
	if err != nil {
		ctx.JSON(401, response.ErrorResponse{
			Message: "Unauthorized",
			Error: "User not found",
		})
		ctx.Abort()
	}
	ctx.Next()
}