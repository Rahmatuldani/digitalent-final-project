package controllers

import (
	"strconv"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SocialMediaControllerStruct struct {
	model    models.SocialMediaInterface
	validate *validator.Validate
}

func SocialMediaController(model models.SocialMediaInterface, v *validator.Validate) *SocialMediaControllerStruct {
	return &SocialMediaControllerStruct{
		model:    model,
		validate: v,
	}
}

// SocialMedia godoc
// @Summary Get social media
// @Description Get social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Success 200 {object} response.GetSocialMedia
// @Failure 500 {object} response.ErrorResponse
// @Router /socialmedias [get]
func (m *SocialMediaControllerStruct) GetSocialMedia(ctx *gin.Context) {
	socialMedias, err := m.model.Get()
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Get social media error",
			Error:   err.Error(),
		})
		return
	}

	array := []response.SocialMediaStruct{}
	for _, v := range socialMedias {
		array = append(array, response.SocialMediaStruct{
			Id:        v.ID,
			Name:      v.Name,
			Url:       v.Url,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: response.UserSocialMedia{
				Id:       v.User.ID,
				Username: v.User.Username,
				Email:    v.User.Email,
			},
		})
	}
	result := response.GetSocialMedia{
		SocialMedias: array,
	}
	ctx.JSON(200, result)
}

// SocialMedia godoc
// @Summary Post social media
// @Description Post social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param req body request.SocialMediaReq true "Request Body"
// @Success 201 {object} response.PostSocialMedia
// @Failure 500 {object} response.ErrorResponse
// @Router /socialmedias [post]
func (m *SocialMediaControllerStruct) PostSocialMedia(ctx *gin.Context) {
	id := ctx.MustGet("userId").(uint64)
	var req request.SocialMediaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't bind JSON",
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
	socialMedia, err := m.model.Post(uint(id), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Create social media error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(201, response.PostSocialMedia{
		Id:        socialMedia.ID,
		Name:      socialMedia.Name,
		Url:       socialMedia.Url,
		UserId:    socialMedia.UserId,
		CreatedAt: socialMedia.CreatedAt,
	})
}

// SocialMedia godoc
// @Summary Update social media
// @Description Update social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param socialMediaId path int true "Social Media ID"
// @Param req body request.SocialMediaReq true "Request Body"
// @Success 200 {object} response.UpdateSocialMedia
// @Failure 500 {object} response.ErrorResponse
// @Router /socialmedias/{socialMediaId} [put]
func (m *SocialMediaControllerStruct) UpdateSocialMedia(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't read param id",
			Error:   err.Error(),
		})
		return
	}
	var req request.SocialMediaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't bind JSON",
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
	socialMedia, err := m.model.Update(uint(userId), uint(aid), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Update social media error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(201, response.UpdateSocialMedia{
		Id:        socialMedia.ID,
		Name:      socialMedia.Name,
		Url:       socialMedia.Url,
		UserId:    socialMedia.UserId,
		UpdatedAt: socialMedia.UpdatedAt,
	})
}

// SocialMedia godoc
// @Summary Delete social media
// @Description Delete social media
// @Tags SocialMedia
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param socialMediaId path int true "Social Media ID"
// @Success 200 {object} response.WebResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /socialmedias/{socialMediaId} [delete]
func (m *SocialMediaControllerStruct) DeleteSocialMedia(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(uint64)
	id := ctx.Param("id")
	aid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(400, response.ErrorResponse{
			Message: "Can't read param id",
			Error:   err.Error(),
		})
		return
	}
	if err := m.model.Delete(uint(userId), uint(aid)); err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Delete social media error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(200, response.WebResponse{
		Message: "Your social media has been successfully deleted",
	})
}