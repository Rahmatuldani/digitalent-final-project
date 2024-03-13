package response

type UserRegRes struct {
	Age			uint	`json:"age"`
	Email		string	`json:"email"`
	Id			uint	`json:"id"`
	Username	string	`json:"username"`
}