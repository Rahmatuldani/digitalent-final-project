package request

type CommentRequest struct {
	Message	string	`json:"message" validate:"required"`
	PhotoId	uint	`json:"photo_id" validate:"required"`
}

type CommentUpdateReq struct {
	Message	string	`json:"message" validate:"required"`
}