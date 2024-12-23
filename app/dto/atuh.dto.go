package dto

type LoginInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterInput struct {
	Name        string  `json:"name" validate:"required,min=3,max=32"`
	Phone       string  `json:"phone" validate:"required,email"`
	Email       *string `json:"email" validate:"omitempty,email"`
	Password    string  `json:"password" validate:"required,min=8,max=32"`
	CountryCode string  `json:"country_code" validate:"required"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type Token struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}

type ResponseLogin struct {
	User        interface{} `json:"user"`
	AccessToken Token       `json:"access_token"`
}
type UserData struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email"`
	Phone string  `json:"phone"`
}
