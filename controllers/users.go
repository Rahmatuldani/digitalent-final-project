package controllers

import (
	"errors"
	"strconv"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/helper"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsersControllerStruct struct {
	model    models.UsersInterface
	validate *validator.Validate
}

func UsersController(model models.UsersInterface, v *validator.Validate) *UsersControllerStruct {
	return &UsersControllerStruct{
		model:    model,
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
// @Param req body request.UserLogin true "Request Body"
// @Success 200 {object} response.TokenJWT
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/login [post]
func (m *UsersControllerStruct) Login(ctx *gin.Context) {
	var req request.UserLogin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't Bind JSON",
			Error:   err.Error(),
		})
		return
	}

	if err := m.validate.Struct(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "JSON does not match the request",
			Error:   err.Error(),
		})
		return
	}

	user, err := m.model.Login(req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Login failed",
			Error:   err.Error(),
		})
		return
	}

	token, err := helper.GenerateJWT([]byte(strconv.FormatUint(uint64(user.ID), 10)))
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Login failed",
			Error:   errors.New("error generate jwt").Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"token": token,
	})
}

// Users godoc
// @Summary User register
// @Schemes
// @Description User register
// @Tags users
// @Accept json
// @Produce json
// @Param req body request.UserRegReq true "Request Body"
// @Success 201 {object} response.UserRegRes
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/register [post]
func (m *UsersControllerStruct) Register(ctx *gin.Context) {
	var req request.UserRegReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't Bind JSON",
			Error:   err.Error(),
		})
		return
	}

	if err := m.validate.Struct(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "JSON does not match the request",
			Error:   err.Error(),
		})
		return
	}
	result, err := m.model.Register(req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Can't register user",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(201, response.UserRegRes{
		Age:      result.Age,
		Email:    result.Email,
		Id:       result.ID,
		Username: result.Username,
	})
}

// Users godoc
// @Summary User update
// @Schemes
// @Description User update
// @Tags users
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param id path int true "ID"
// @Param req body request.UserUpdateReq true "Request Body"
// @Success 200 {object} response.WebResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users/{id} [put]
func (m *UsersControllerStruct) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't read param id",
			Error: err.Error(),
		})
		return
	}

	var req request.UserUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't Bind JSON",
			Error:   err.Error(),
		})
		return
	}

	if err := m.validate.Struct(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "JSON does not match the request",
			Error:   err.Error(),
		})
		return
	}
	
	user, err := m.model.Update(uint8(aid), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Server Error",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(200, response.UserUpdateRes{
		Id: user.ID,
		Email: user.Email,
		Username: user.Username,
		Age: user.Age,
		UpdatedAt: user.UpdatedAt,
	})
}

// Users godoc
// @Summary User delete
// @Schemes
// @Description User delete
// @Tags users
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Success 200 {object} response.WebResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /users [delete]
func (m *UsersControllerStruct) Delete(ctx *gin.Context) {
	id := ctx.MustGet("userId").(uint64)
	if err := m.model.Delete(uint8(id)); err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Delete user error",
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(200, response.WebResponse{
		Message: "Your account has been successfully deleted",
	})
}