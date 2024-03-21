package response

import "time"

type UserSocialMedia struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type SocialMediaStruct struct {
	Id        uint            `json:"id"`
	Name      string          `json:"name"`
	Url       string          `json:"social_media_url"`
	UserId    uint            `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	User      UserSocialMedia `json:"User"`
}

type GetSocialMedia struct {
	SocialMedias []SocialMediaStruct `json:"social_medias"`
}

type PostSocialMedia struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"social_media_url"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateSocialMedia struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"social_media_url"`
	UserId    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
