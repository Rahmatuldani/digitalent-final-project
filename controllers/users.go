package controllers

import (
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersControllerStruct struct {
	model models.UsersInterface
	validate *validator.Validate
}

func UsersController(model models.UsersInterface, v *validator.Validate) *UsersControllerStruct {
	return &UsersControllerStruct{
		model: model,
		validate: v,
	}
}

// Users godoc
// @Summary User login
// @Schemes
// @Description User login
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.TokenJWT
// @Router /users/login [post]
func (m *UsersControllerStruct) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"token": "jwt string",
	})
}

// Users godoc
// @Summary User register
// @Schemes
// @Description User register
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.UserRegRes
// @Router /users/register [post]
func (m *UsersControllerStruct) Register(ctx *gin.Context) {
	ctx.JSON(200, response.UserRegRes{
		Age: 10,
		Email: "email",
		Id: 1,
		Username: "username",
	})
}