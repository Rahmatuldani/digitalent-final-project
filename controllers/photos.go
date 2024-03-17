package controllers

import (
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
// @Schemes
// @Description Get all photos
// @Tags photos
// @Accept json
// @Produce json
// @Success 200 {object} response.PhotosGetRes
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /photos [get]
func (m *PhotosControllerStruct) GetAllPhoto(ctx *gin.Context) {
	photos, err := m.model.GetPhotos()
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
// @Schemes
// @Description Post photo
// @Tags photos
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param req body request.PhotoPostReq true "Request Body"
// @Success 201 {object} response.PhotoPostRes
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /photos [post]
func (m *PhotosControllerStruct) PostPhoto(ctx *gin.Context) {
	id := ctx.MustGet("userId").(uint64)
	var req request.PhotoPostReq
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
	result, err := m.model.PostPhoto(uint(id), req)
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