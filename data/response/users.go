package response

import "time"

type UserRegRes struct {
	Age      uint8  `json:"age"`
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type UserUpdateRes struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint8     `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}
