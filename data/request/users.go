package request

type UserRegReq struct {
	Age			uint	`json:"age" validate:"required"`
	Email		string	`json:"email" validate:"required"`
	Password 	string	`json:"password"`
	Username	string	`json:"username"`
}