package controllers

import (
	"strconv"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotosControllerStruct struct {
	model    models.PhotosInterface
	validate *validator.Validate
}

func PhotosController(model models.PhotosInterface, v *validator.Validate) *PhotosControllerStruct {
	return &PhotosControllerStruct{
		model:    model,
		validate: v,
	}
}

// Photos godoc
// @Summary Get all photos
// @Description Get all photos
// @Tags photos
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Success 200 {object} response.PhotosGetRes
// @Failure 500 {object} response.ErrorResponse
// @Router /photos [get]
func (m *PhotosControllerStruct) GetPhotos(ctx *gin.Context) {
	photos, err := m.model.Get()
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Server Error",
			Error: err.Error(),
		})
		return
	}
	var result []response.PhotosGetRes
	for _, v := range photos {
		result = append(result, response.PhotosGetRes{
			Id: v.ID,
			Title: v.Title,
			Caption: v.Caption,
			PhotoUrl: v.PhotoUrl,
			UserId: v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: response.UserPhotos{
				Email: v.User.Email,
				Username: v.User.Username,
			},
		})
	}
	ctx.JSON(200, result)

}

// Photos godoc
// @Summary Post photo
// @Description Post photo
// @Tags photos
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param req body request.PhotoReq true "Request Body"
// @Success 201 {object} response.PhotoPostRes
// @Failure 500 {object} response.ErrorResponse
// @Router /photos [post]
func (m *PhotosControllerStruct) PostPhoto(ctx *gin.Context) {
	id := ctx.MustGet("userId").(uint64)
	var req request.PhotoReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't bind JSON",
			Error: err.Error(),
		})
		return
	}
	if err := m.validate.Struct(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "JSON does not match the request",
			Error: err.Error(),
		})
		return
	}
	result, err := m.model.Post(uint(id), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Server Error",
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(201, response.PhotoPostRes{
		Id: result.ID,
		Title: result.Title,
		Caption: result.Caption,
		PhotoUrl: result.PhotoUrl,
		UserId: result.UserId,
		CreatedAt: result.CreatedAt,
	})
}

// Photos godoc
// @Summary Update photo
// @Description Update photo
// @Tags photos
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param photoId path int true "Photo ID"
// @Param req body request.PhotoReq true "Request Body"
// @Success 200 {object} response.PhotoUpdateRes
// @Failure 500 {object} response.ErrorResponse
// @Router /photos/{photoId} [put]
func (m *PhotosControllerStruct) UpdatePhoto(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't read param id",
			Error: err.Error(),
		})
		return
	}
	var req request.PhotoReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't bind JSON",
			Error: err.Error(),
		})
		return
	}
	if err := m.validate.Struct(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "JSON does not match the request",
			Error: err.Error(),
		})
		return
	}
	result, err := m.model.Update(uint(userId), uint(aid), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Update photo error",
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(200, response.PhotoUpdateRes{
		Id: result.ID,
		Title: result.Title,
		Caption: result.Caption,
		PhotoUrl: result.PhotoUrl,
		UserId: result.UserId,
		UpdatedAt: result.UpdatedAt,
	})
}

// Photos godoc
// @Summary Delete photo
// @Description Delete photo
// @Tags photos
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param photoId path int true "User ID"
// @Success 200 {object} response.WebResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /photos/{photoId} [delete]
func (m *PhotosControllerStruct) DeletePhoto(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't read param id",
			Error: err.Error(),
		})
		return
	}
	if err := m.model.Delete(uint(userId), uint(aid)); err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Delete photo error",
			Error: err.Error(),
		})
		return
	}
	ctx.JSON(200, response.WebResponse{
		Message: "Your photo has been successfully deleted",
	})
}