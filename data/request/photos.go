package request

type PhotoReq struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}