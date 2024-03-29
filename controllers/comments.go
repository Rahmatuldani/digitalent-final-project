package controllers

import (
	"strconv"

	"github.com/Rahmatuldani/digitalent-project/data/request"
	"github.com/Rahmatuldani/digitalent-project/data/response"
	"github.com/Rahmatuldani/digitalent-project/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentsControllerStruct struct {
	model    models.CommentsInterface
	validate *validator.Validate
}

func CommentsController(model models.CommentsInterface, v *validator.Validate) *CommentsControllerStruct {
	return &CommentsControllerStruct{
		model:    model,
		validate: v,
	}
}

// Comments godoc
// @Summary Get comments
// @Description Get comments
// @Tags comments
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Success 200 {object} response.GetComments
// @Failure 500 {object} response.ErrorResponse
// @Router /comments [get]
func (m *CommentsControllerStruct) GetComment(ctx *gin.Context) {
	comments, err := m.model.Get()
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Get comments error",
			Error:   err.Error(),
		})
		return
	}
	var result []response.GetComments
	for _, v := range comments {
		result = append(result, response.GetComments{
			Id:        v.ID,
			Message:   v.Message,
			PhotoId:   v.PhotoId,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: response.UserComment{
				Id:       v.User.ID,
				Email:    v.User.Email,
				Username: v.User.Username,
			},
			Photo: response.PhotoComment{
				Id:       v.Photo.ID,
				Title:    v.Photo.Title,
				Caption:  v.Photo.Caption,
				PhotoUrl: v.Photo.PhotoUrl,
				UserId:   v.Photo.UserId,
			},
		})
	}
	ctx.JSON(200, result)
}

// Comments godoc
// @Summary Post comments
// @Description Post comments
// @Tags comments
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param req body request.CommentRequest true "Request Body"
// @Success 201 {object} response.PostComment
// @Failure 500 {object} response.ErrorResponse
// @Router /comments [post]
func (m *CommentsControllerStruct) PostComment(ctx *gin.Context) {
	id := ctx.MustGet("userId").(uint64)
	var req request.CommentRequest
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
	comment, err := m.model.Post(uint(id), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Create comment error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(201, response.PostComment{
		Id:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
	})
}

// Comments godoc
// @Summary Update comment
// @Description Update comment
// @Tags comments
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param commentId path int true "Comment ID"
// @Param req body request.CommentUpdateReq true "Request Body"
// @Success 200 {object} response.UpdateComment
// @Failure 500 {object} response.ErrorResponse
// @Router /comments/{commentId} [put]
func (m *CommentsControllerStruct) UpdateComment(ctx *gin.Context) {
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
	var req request.CommentUpdateReq
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
	comment, err := m.model.Update(uint(userId), uint(aid), req)
	if err != nil {
		ctx.JSON(500, response.ErrorResponse{
			Message: "Update comment error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(200, response.UpdateComment{
		Id:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		UpdatedAt: comment.UpdatedAt,
	})
}

// Comments godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags comments
// @Accept json
// @Produce json
// @Param Bearer header string true "Bearer Token"
// @Param commentId path int true "Comment ID"
// @Success 200 {object} response.WebResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /comments/{commentId} [delete]
func (m *CommentsControllerStruct) DeleteComment(ctx *gin.Context) {
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
			Message: "Delete comment error",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(200, response.WebResponse{
		Message: "Your comment has been successfully deleted",
	})
}
